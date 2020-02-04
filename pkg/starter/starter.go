package starter

import (
	"context"
	"time"

	configclient "github.com/openshift/client-go/config/clientset/versioned"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	imagev1client "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	"github.com/openshift/clustercli/pkg/controller/clustercli"
	cc "github.com/openshift/library-go/pkg/controller/controllercmd"
	"k8s.io/klog/v2"
)

func RunCLIController(ctx context.Context, cc *cc.ControllerContext) error {
	const resync = 10 * time.Minute

	configClient, err := configclient.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}
	imageClient, err := imagev1client.NewForConfig(cc.KubeConfig)
	if err != nil {
		return err
	}

	configInformers := configinformers.NewSharedInformerFactoryWithOptions(configClient, resync)
	recorder := cc.EventRecorder

	clusterCLIController, err := clustercli.NewClusterCLISyncController(
		ctx,
		configClient.ConfigV1().ClusterCLIs(),
		configInformers.Config().V1().ClusterCLIs(),
		// TODO: don't hardcode ns, get from api
		imageClient.ImageStreams("openshift"),
		recorder,
	)
	if err != nil {
		return err
	}

	configInformers.Start(ctx.Done())
	go clusterCLIController.Run(ctx, 1)

	<-ctx.Done()
	return nil
}
