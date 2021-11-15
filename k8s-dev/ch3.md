1) K8s.io/client-go
2) Clientcmd.BuildConfigFromFlags
3) JSON
4) Kind, APIVersion
5) True
6) All API Groups and resources defined in k8s.io/api, but not for custom resource definitions
7) List, Update, Create, Logs
8) 60 sec
9) Better error handling in informers, less resource heavy as compared to watch, it has in-memory cache, so it's fast, can react to changes in k8s objects nearly in real-time.
10) RESTMapping
11) Formatted in CamelCase and usually singular
12) Formatted in lowercase, and usually plural
13) Go types, REST Endpoints
14) services