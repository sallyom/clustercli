#!/bin/bash
set -euxo pipefail

oc apply -f vendor/github.com/openshift/api/config/v1/0000_03_config-operator_01_clustercli.crd.yaml
pushd manifests
oc apply -f namespace.yaml
oc apply -f serviceaccount.yaml
oc apply -f serviceca.yaml
oc apply -f trusted_ca.yaml
oc apply -f roles.yaml
oc apply -f index-cm.yaml
oc apply -f cr.yaml
oc apply -f clustercli-service.yaml
oc apply -f configmap.yaml
oc apply -f nginx-cm.yaml
oc apply -f deployment.yaml
oc apply -f route.yaml
popd
