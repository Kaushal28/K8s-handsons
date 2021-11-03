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

## Task 3

- When changing the label of one of the pods of replicaset (managed by deployment) will make replicaset believe that the current state of cluster is not same as desired state because the desired state is determined by pod's label using label selector. So it'll spin up a new pod with the old label to match the current state with desired state.

## Task 4

- Created an nginx deployment and a cluster IP type service to expose the nginx pod then created a busybox pod and tried accessing the nginx using cluster IP as well as service name and it returned correct response.
- Created nginx deployment and a corresponding service to expose it on port 80. Created two namespaces and tried accessing nginx from both the namespaces and it nginx was accessible via cluster IP from both the namespaces. Then created a network policy on nginx pod to restrict traffic from all other namespaces and then tried accessing it from both the namespaces and was not able to access it from other namespace as expected.

## Task 5

- Created a config map with key-value pair and an nginx pod with env variable's value that is referenced from config map. From nginx pod, tried printing the environment variable using `echo $envname` and it printed the correct value.
- Created an opaque secret and provided base64 encoded secret values. In the pod definition, using `envFrom`, referenced the secret and then from inside the container, echoed both the env variables and got them as base64 decoded values.
- Created secret using `--from-literal` and via files. Values are required to be base64 encoded while creating using file but when using `--from-literal`, values can be plain text and k8s will encode it automatically.

## Task 6

- Went through official K8s docs as well as kode kloud video for dynamic volume provisioning.
- Created two containers in a single pod and mounted emptydir type volume in both the containers. One container wrote date in mounted volume and then from inside the other container, checked the value of date written by the first container.
- Created the same configuration in task2b. Drawbacks of hostPath as volume is that in multinode cluster, pods access the same path on all the nodes and will expect to have the same data on all the nodes. It's not possible as they are on different machines. Also, there are security risks as well where a container can access secure details and can attack other parts of the cluster.