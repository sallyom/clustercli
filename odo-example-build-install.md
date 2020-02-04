### Example ODO image build/install

```console
$ podman run -v $HOME/installer/cli/:/home:Z --rm -it quay.io/sallyom/odo:latest sh
/ # ls -al /tmp
total 32
drwxrwxrwt    8 root     root          4096 Feb  7 18:54 .
drwxr-xr-x    2 root     root          4096 Feb  7 19:06 ..
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 390x
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 aarch64
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 amd64
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 darwin
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 ppc64le
drwxr-xr-x    2 root     root          4096 Feb  7 18:54 windows
/ # cp /tmp/amd64/odo.tar.gz /home/.
/ # exit
localhost:cli$ tar xvf odo.tar.gz
odo
localhost:cli$ ./odo version
odo v1.1.0 (86c7e2c9)

Server: https://api.ci.openshift.org:443
OpenShift: v3.11.0+e5dbec2-186
Kubernetes: v1.11.0+d4cacc0
```

