# kube-prompt

An interactive kubernetes client featuring auto-complete written in Go.

![demo](./_resources/kube-prompt.gif)

kube-prompt's command is same with kubectl (because basically this is just wrapper).
So it doesn't require the additional cost to learn the usage of kube-prompt.

Binaries are available from github releases:

* macOS (darwin)
* Linux

## Goal

Hopefully support following commands and resource types enough to operate kubernetes as kubectl.

* [x] `get`            Display one or many resources
* [x] `describe`       Show details of a specific resource or group of resources
* [x] `create`         Create a resource by filename or stdin
* [x] `replace`        Replace a resource by filename or stdin.
* [x] `patch`          Update field(s) of a resource using strategic merge patch.
* [x] `delete`         Delete resources by filenames, stdin, resources and names, or by resources and label selector.
* [x] `edit`           Edit a resource on the server
* [x] `apply`          Apply a configuration to a resource by filename or stdin
* [x] `namespace`      SUPERSEDED: Set and view the current Kubernetes namespace
* [x] `logs`           Print the logs for a container in a pod.
* [x] `rolling-update` Perform a rolling update of the given ReplicationController.
* [x] `scale`          Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job.
* [x] `cordon`         Mark node as unschedulable
* [x] `drain`          Drain node in preparation for maintenance
* [x] `uncordon`       Mark node as schedulable
* [x] `attach`         Attach to a running container.
* [x] `exec`           Execute a command in a container.
* [x] `port-forward`   Forward one or more local ports to a pod.
* [x] `proxy`          Run a proxy to the Kubernetes API server
* [x] `run`            Run a particular image on the cluster.
* [x] `expose`         Take a replication controller, service, or pod and expose it as a new Kubernetes Service
* [x] `autoscale`      Auto-scale a Deployment, ReplicaSet, or ReplicationController
* [x] `rollout`        rollout manages a deployment
* [x] `label`          Update the labels on a resource
* [x] `annotate`       Update the annotations on a resource
* [x] `config`         config modifies kubeconfig files
* [x] `cluster-info`   Display cluster info
* [x] `api-versions`   Print the supported API versions on the server, in the form of "group/version".
* [x] `version`        Print the client and server version information.
* [x] `explain`        Documentation of resources.
* [x] `convert`        Convert config files between different API versions

## LICENSE

This software is licensed under the MIT License (See [LICENSE](./LICENSE)).
