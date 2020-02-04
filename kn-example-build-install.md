### Example kn image build/install

```console
$ podman run -v $HOME/installer/cli/:/home:Z --rm -it quay.io/sallyom/kn:latest sh
/ # ls -al /tmp
total 20
drwxrwxrwt    5 root     root          4096 Feb  7 18:49 .
drwxr-xr-x    2 root     root          4096 Feb  7 19:02 ..
drwxr-xr-x    2 root     root          4096 Feb  7 18:49 amd64
drwxr-xr-x    2 root     root          4096 Feb  7 18:49 darwin
drwxr-xr-x    2 root     root          4096 Feb  7 18:49 windows
/ # cp /tmp/amd64/kn.tar.gz /home/.
/ # exit
localhost:cli$ ls -al
total 11784
drwxr-xr-x.  2 somalley somalley     4096 Feb  7 14:02 .
drwxrwxr-x. 11 somalley somalley     4096 Feb  7 12:06 ..
-rw-r--r--.  1 somalley somalley 12055909 Feb  7 14:02 kn.tar.gz
localhost:cli$ tar xvf kn.tar.gz 
./
./LICENSE
./kn
localhost:cli$ ./kn version
Version:      v0.11.0
Build Date:   2020-01-22 13:02:24
Git Revision: c7c93b7
Supported APIs:
* Serving
  - serving.knative.dev/v1alpha1 (knative-serving v0.11.0)
* Eventing
  - sources.eventing.knative.dev/v1alpha1 (knative-eventing v0.11.0)
  - eventing.knative.dev/v1alpha1 (knative-eventing v0.11.0)
```
