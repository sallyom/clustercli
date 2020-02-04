package clustercli

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	configv1 "github.com/openshift/api/config/v1"
	clustercliclientv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	clustercliinformersv1 "github.com/openshift/client-go/config/informers/externalversions/config/v1"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourcemerge"
)

const (
	controllerWorkQueueKey = "cli-manager-sync-work-queue-key"
	controllerName         = "cli-manager"

	ocCustomResourceName  = "oc-clustercli"
	odoCustomResourceName = "odo-clustercli"
	knCustomResourceName  = "kn-clustercli"
	ClusterCLINamespace   = "default"
)

type ClusterCLISyncController struct {
	clusterCLIClient clustercliclientv1.ClusterCLIInterface
	ocImage          string
	odoImage         string
	knImage          string
}

func NewClusterCLISyncController(
	ctx context.Context,
	clusterCLIInterface clustercliclientv1.ClusterCLIInterface,
	clusterCLIInformers clustercliinformersv1.ClusterCLIInformer,
	isClient imagev1client.ImageStreamInterface,
	eventRecorder events.Recorder,
) (factory.Controller, error) {
	ocStream, err := isClient.Get(ctx, "cli-artifacts", metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, fmt.Errorf("no cli-artifacts image stream found")
		}
		return nil, err
	}
	c := &ClusterCLISyncController{
		clusterCLIClient: clusterCLIInterface,
		// TODO: get from imagestream in cluster
		odoImage: "quay.io/sallyom/odo:latest",
		knImage:  "quay.io/sallyom/kn:latest",
		ocImage:  ocStream.Spec.Tags[0].From.Name,
	}
	return factory.New().WithInformers(
		clusterCLIInformers.Informer(),
	).WithSync(c.sync).ToController("ClusterCLIController", eventRecorder.WithComponentSuffix("clustercli-controller")), nil
}

func (c *ClusterCLISyncController) sync(ctx context.Context, syncContext factory.SyncContext) error {
	klog.Infof("syncing ClusterCLI custom resources")
	var ocMappings []*configv1.ClusterCLIMapping
	var odoMappings []*configv1.ClusterCLIMapping
	var knMappings []*configv1.ClusterCLIMapping
	oss := []string{"darwin", "windows", "linux"}
	arches := []string{"amd64", "arm64", "ppc64le", "s390x"}
	for _, os := range oss {
		if os == "linux" {
			for _, arch := range arches {
				ocMappings = append(ocMappings, PlatformBasedMapping("oc", "linux", arch))
				odoMappings = append(odoMappings, PlatformBasedMapping("odo", "linux", arch))
			}
			knMappings = append(knMappings, PlatformBasedMapping("kn", "linux", "amd64"))
		}
		ocMappings = append(ocMappings, PlatformBasedMapping("oc", os, "amd64"))
		odoMappings = append(odoMappings, PlatformBasedMapping("odo", os, "amd64"))
		knMappings = append(knMappings, PlatformBasedMapping("kn", os, "amd64"))
	}
	klog.Infof("OC-Mappings: %v", ocMappings)
	klog.Infof("ODO-Mappings: %v", odoMappings)
	klog.Infof("KN-Mappings: %v", knMappings)
	ocClusterCLI, err := c.PlatformBasedClusterCLI("oc", ocCustomResourceName, ocMappings)
	if err != nil {
		return err
	}
	_, reason, err := ApplyClusterCLI(ctx, c.clusterCLIClient, ocClusterCLI)
	if err != nil {
		return fmt.Errorf("%s: %v", reason, err)
	}

	odoClusterCLI, err := c.PlatformBasedClusterCLI("odo", odoCustomResourceName, odoMappings)
	if err != nil {
		return err
	}
	_, reason, err = ApplyClusterCLI(ctx, c.clusterCLIClient, odoClusterCLI)
	if err != nil {
		return fmt.Errorf("%s: %v", reason, err)
	}

	knClusterCLI, err := c.PlatformBasedClusterCLI("kn", knCustomResourceName, knMappings)
	if err != nil {
		return err
	}
	_, reason, err = ApplyClusterCLI(ctx, c.clusterCLIClient, knClusterCLI)
	if err != nil {
		return fmt.Errorf("%s: %v", reason, err)
	}

	return nil
}

func PlatformBasedMapping(cli, os, arch string) *configv1.ClusterCLIMapping {
	// OS: "darwin", "windows", "linux"
	// Arch: "amd64", "arm64", "ppc64le", "s390x"
	from := ""
	switch cli {
	case "oc":
		switch os {
		case "windows":
			from = "/usr/share/openshift/windows/windows/oc.zip"
		case "darwin":
			from = "/usr/share/openshift/mac/oc.tar.gz"
		case "linux":
			// amd64, arm64, ppc64le, s390x
			from = fmt.Sprintf("/usr/share/openshift/linux_%s/oc.tar.gz", arch)
		default:
			return &configv1.ClusterCLIMapping{}
		}
	case "odo":
		switch os {
		case "windows":
			from = "/tmp/windows/odo.exe.tar.gz"
		case "darwin":
			from = "/tmp/darwin/odo.tar.gz"
		case "linux":
			// amd64, arm64, ppc64le, 390x
			from = fmt.Sprintf("/tmp/%s/odo.tar.gz", arch)
		default:
			return &configv1.ClusterCLIMapping{}
		}
	case "kn":
		switch os {
		case "windows":
			from = "/tmp/windows/kn.zip"
		case "darwin":
			from = "/tmp/darwin/kn.tar.gz"
		case "linux":
			// amd64
			from = fmt.Sprintf("/tmp/%s/kn.tar.gz", arch)
		default:
			return &configv1.ClusterCLIMapping{}
		}
	default:
		return &configv1.ClusterCLIMapping{}
	}
	//var mappings []configv1.ClusterCLIMapping
	return &configv1.ClusterCLIMapping{
		OS:   os,
		Arch: arch,
		From: from,
	}
}

func (c *ClusterCLISyncController) PlatformBasedClusterCLI(cli, crName string, mappings []*configv1.ClusterCLIMapping) (*configv1.ClusterCLI, error) {
	// OS: "darwin", "windows", "linux"
	// Arch: "amd64", "arm64", "ppc64le", "s390x"
	var image, description, displayName string
	// get home directory for mapping
	//currentUser, err := user.Current()
	//if err != nil {
	//	return nil, err
	//}
	//installDir := filepath.Join(currentUser.HomeDir, "openshift-clis")
	//os := runtime.GOOS
	//arch := runtime.GOARCH
	switch cli {
	case "oc":
		image = c.ocImage
		description = `With the OpenShift command line interface, you can create applications and manage OpenShift projects from a terminal.  The oc binary offers the same capabilities as the kubectl binary, but it is further extended to natively support OpenShift Container Platform features.`
		displayName = "oc - OpenShift Command Line Interface (CLI)"
	case "odo":
		image = c.odoImage
		description = `OpenShift Do (odo) is a fast, iterative, and straightforward CLI tool for developers who write, build, and deploy applications on OpenShift.  odo abstracts away complex Kubernetes and OpenShift concepts, thus allowing developers to focus on what is most important to them: code.`
		displayName = "odo - Developer-focused CLI for OpenShift"
	case "kn":
		image = c.knImage
		description = `Knative (pronounced “Kay - Native”) extends Kubernetes to provide a set of components for deploying, running and managing modern applications running serverless.  Knative Client (kn) is the Knative command line interface (CLI). The CLI exposes commands for managing your applications, as well as lower level tools to interact with components of OpenShift Container Platform. With kn, you can create applications and manage OpenShift Container Platform projects from the terminal.`
		displayName = "kn - Knative CLI for running and managing serverless applications"
	default:
		return nil, fmt.Errorf("unknown cli: %s", cli)
	}
	var maps []configv1.ClusterCLIMapping
	for _, m := range mappings {
		maps = append(maps, *m)
	}
	return &configv1.ClusterCLI{
		ObjectMeta: metav1.ObjectMeta{
			Name: crName,
		},
		Spec: configv1.ClusterCLISpec{
			Description: description,
			DisplayName: displayName,
			Image:       image,
			Mapping:     maps,
		},
	}, nil
}

func ApplyClusterCLI(ctx context.Context, cliClient clustercliclientv1.ClusterCLIInterface, requiredClusterCLI *configv1.ClusterCLI) (*configv1.ClusterCLI, string, error) {
	cliName := requiredClusterCLI.ObjectMeta.Name
	existingClusterCLI, err := cliClient.Get(ctx, cliName, metav1.GetOptions{})
	existingClusterCLICopy := existingClusterCLI.DeepCopy()
	if errors.IsNotFound(err) {
		actualClusterCLI, err := cliClient.Create(ctx, requiredClusterCLI, metav1.CreateOptions{})
		if err != nil {
			klog.Infof("error creating %s clusterCLI custom resource: %s", cliName, err)
			return nil, "FailedCreate", err
		}
		klog.Infof("%s clusterCLI custom resource created", cliName)
		return actualClusterCLI, "", nil
	}
	if err != nil {
		klog.Infof("error getting %s custom resource: %v", cliName, err)
		return nil, "", err
	}
	specSame := equality.Semantic.DeepEqual(existingClusterCLICopy.Spec, requiredClusterCLI.Spec)
	modified := resourcemerge.BoolPtr(false)
	resourcemerge.EnsureObjectMeta(modified, &existingClusterCLICopy.ObjectMeta, requiredClusterCLI.ObjectMeta)
	if specSame && !*modified {
		klog.Infof("%s clusterCLI custom resource exists and is in the correct state", cliName)
		return existingClusterCLICopy, "", nil
	}

	existingClusterCLICopy.Spec = requiredClusterCLI.Spec
	actualClusterCLI, err := cliClient.Update(ctx, existingClusterCLICopy, metav1.UpdateOptions{})
	if err != nil {
		klog.Infof("error updating %s clusterCLI custom resource: %v", cliName, err)
		return nil, "FailedUpdate", err
	}
	return actualClusterCLI, "", nil
}
