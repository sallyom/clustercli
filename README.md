### ClusterCLI controller for ClusterCLI CRs in OpenShift

To install this controller, fileserver, crd, and custom resources for oc, kn, and odo:

* `./install.sh` 
    * This installs ClusterCLI CRD + all resources in ./manifests in namespace `openshift-cli-manager`

* [CRD](https://github.com/openshift/api/blob/285ad97bbf708658eb4f28a5c58819bb330d7813/config/v1/0000_03_config-operator_01_clustercli.crd.yaml) 
