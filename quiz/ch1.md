1) kube-scheduler
2) deployment -> replicaset -> kube-scheduler -> kubelet
3) Accept first write, rejects second. Controller must retry / handle conflicts