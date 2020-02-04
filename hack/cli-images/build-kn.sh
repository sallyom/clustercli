#!/usr/bin/env bash

set -o errexit

# Create a container
container=$(buildah from alpine)

buildah config --label maintainer="Sally O'Malley <somalley@redhat.com>" $container

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/serverless/latest/kn-linux-amd64-0.11.0.tar.gz -o kn.tar.gz
buildah copy $container kn.tar.gz /tmp/amd64/kn.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/serverless/latest/kn-macos-amd64-0.11.0.tar.gz -o kn.tar.gz
buildah copy $container kn.tar.gz /tmp/darwin/kn.tar.gz

curl -sSL https://mirror.openshift.com/pub/openshift-v4/clients/serverless/latest/kn-windows-amd64-0.11.0.zip -o kn.zip
buildah copy $container kn.tar.gz /tmp/windows/kn.zip

# Finally saves the running container to an image
buildah commit --format docker $container quay.io/sallyom/kn:latest
podman push quay.io/sallyom/kn:latest
