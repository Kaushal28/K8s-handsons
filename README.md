# Kubernetes module tasks

Setup a Kubernetes cluster using `kubeadm` in local machine.

## Task 1

- Checked the manifest YAML files located at `/etc/kubernetes/manifests`. Went through the flags for specific component mentioned under `.containers.commands` and checked the significance of each flag from k8s docs (https://kubernetes.io/docs/reference/command-line-tools-reference/).
- In kube-controller-manager's manifest file, checked the flag named `controllers` which has value of `*,bootstrapsigner,tokencleaner`, that means by default all the controllers are enabled along with `bootstrapsigner` and `tokencleaner` controllers.

## Task 2

- Created a YAML config file for Pod haing image `praqma/network-multitool` and network mode as `host` (so that commands can be executed against host). Then it allows to execute supported network commands from the container. (using `kubectl exec <pod-name> -it -- bin/bash`)
- Created a Go based static file server and corresponding docker image. Pushed that image into a local docker registry. Then created a Pod with mounted directory from host where the file would be hosted. Configured liveness and readiness probes to check whether the file exists in mounted path. Configured a busybox init container to create file on pod startup in mounted path. Also created a kubernates service to access the file server from host.
- Created a daemon set with nginx container.
- Created an nginx deployment and then updated the image to an invalid tag using `kubectl set image deployment.v1.apps/nginxdeployment nginx=nginx:latest-not-available`. Then checked the replicaset and pods. As image wasn't available, the deployment got stuck. So to rollback to previous version, used the following command `kubectl rollout undo deployment/nginxdeployment`. Again checked replicaset and pods. Both were showing correct results.
- Created a persistent volume with 1Gi capacity for storing postgres data. In statefulset, created a volume claim to bind the postgres with that volume. Used postgres docker container image for database. Also, created a headless service to access the postgres. Checked the connection with database with PgAdmin4 and it's working.
- When scheduler is not available (because of image pull error after changing image tag to an invalid one), newly created pod would not get scheduled on any node. Therefore the status of newly created pod is always pending instead of running. However the scheduler is also a pod but it's a static one, so that kubelet itself manages it because there is no scheduler available at the time of creating/scheduling the kube scheduler itself. When the kube scheduler becomes available again, the newly created pod moves to running state.