#!/usr/bin/env bash

set -o errexit

# THIS IS ACTUALLY ALREADY DONE WITH openshift/oc cli-artifacts image
# Create a container
container=$(buildah from alpine)

buildah config --label maintainer="Sally O'Malley <somalley@redhat.com>" $container

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/linux/oc.tar.gz -o oc.tar.gz
buildah copy $container oc.tar.gz /tmp/amd64/oc.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/macosx/oc.tar.gz -o oc.tar.gz
buildah copy $container oc.tar.gz /tmp/darwin/oc.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/windows/oc.zip -o oc.zip
buildah copy $container oc.zip /tmp/windows/oc.zip

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/linux-aarch64/oc.tar.gz -o oc.tar.gz
buildah copy $container oc.tar.gz /tmp/arm64/oc.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/linux-ppc64le/oc.tar.gz -o oc.tar.gz
buildah copy $container oc.tar.gz /tmp/ppc64le/oc.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/oc/4.6/linux-s390x/oc.tar.gz -o oc.tar.gz
buildah copy $container oc.tar.gz /tmp/s390x/oc.tar.gz

# Finally saves the running container to an image
buildah commit --format docker $container quay.io/sallyom/oc:4.6
podman push quay.io/sallyom/oc:4.6
