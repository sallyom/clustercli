#!/usr/bin/env bash

set -o errexit

# Create a container
container=$(buildah from alpine)

buildah config --label maintainer="Sally O'Malley <somalley@redhat.com>" $container

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-linux-amd64.tar.gz -o odo.tar.gz
buildah copy $container odo.tar.gz /tmp/amd64/odo.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-darwin-amd64.tar.gz -o odo.tar.gz
buildah copy $container odo.tar.gz /tmp/darwin/odo.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-windows-amd64.exe.tar.gz -o odo.exe.tar.gz
buildah copy $container odo.exe.tar.gz /tmp/windows/odo.exe.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-linux-arm64.tar.gz -o odo.tar.gz
buildah copy $container odo.tar.gz /tmp/arm64/odo.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-linux-ppc64le.tar.gz -o odo.tar.gz
buildah copy $container odo.tar.gz /tmp/ppc64le/odo.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/odo/latest/odo-linux-s390x.tar.gz -o odo.tar.gz
buildah copy $container odo.tar.gz /tmp/s390x/odo.tar.gz

# Finally saves the running container to an image
buildah commit --format docker $container quay.io/sallyom/odo:latest
podman push quay.io/sallyom/odo:latest
