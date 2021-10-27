# Kubernetes module tasks

Setup a Kubernetes cluster using `kubeadm` in local machine.

## Task 1
    - Checked the manifest YAML files located at `/etc/kubernetes/manifests`. Went through the flags for specific component mentioned under `.containers.commands` and checked the significance of each flag from k8s docs (https://kubernetes.io/docs/reference/command-line-tools-reference/).
    - In kube-controller-manager's manifest file, checked the flag named `controllers` which has value of `*,bootstrapsigner,tokencleaner`, that means by default all the controllers are enabled along with `bootstrapsigner` and `tokencleaner` controllers.

## Task 2
    - Created a YAML config file for Pod haing image `praqma/network-multitool` and network mode as `host` (so that commands can be executed against host). Then it allows to execute supported network commands from the container. (using `kubectl exec <pod-name> -it -- bin/bash`)
    - Created a Go based static file server and corresponding docker image. Pushed that image into a local docker registry. Then created a Pod with mounted directory from host where the file would be hosted. Configured liveness and readiness probes to check whether the file exists in mounted path. Configured a busybox init container to create file on pod startup in mounted path. Also created a kubernates service to access the file server from host.
    - 

    https://github.com/Kaushal28/K8s-handsons.git