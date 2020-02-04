### CLI-manager for supplying cluster clis

```console
$ podman run -v $HOME/installer/cli/:/home:Z --rm -it quay.io/sallyom/oc:4.3 sh
/ # ls -al /tmp
total 32
drwxrwxrwt    8 root     root          4096 Feb  7 18:53 .
drwxr-xr-x    2 root     root          4096 Feb  7 19:05 ..
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 390x
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 aarch64
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 amd64
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 darwin
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 ppc64le
drwxr-xr-x    2 root     root          4096 Feb  7 18:53 windows
/ # cp /tmp/amd64/oc.tar.gz /home/.
/ # exit
localhost:cli$ tar xvf oc.tar.gz 
oc
localhost:cli$ ./oc version
Client Version: v4.3.2
Kubernetes Version: v1.11.0+d4cacc0
```

