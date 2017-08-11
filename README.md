# kube-prompt

An interactive kubernetes client featuring auto-complete written in Go.

<a href="https://asciinema.org/a/DQNCOFpUeQayrYlhq2OD1jbqZ" target="_blank">
  <img src="https://asciinema.org/a/DQNCOFpUeQayrYlhq2OD1jbqZ.png" alt="asciicast">
</a>

Recently, I found kube-shell which same as kube-prompt but written in python and python-prompt-toolkit.
If you don't need your effort to set up python environment, maybe kube-shell is great.

kube-prompt written in go and go-prompt, So Binaries are available from:

* macOS (darwin)
* Linux

Windows support is still not because [go-prompt](https://github.com/c-bata/go-prompt) is only supports VT100 console emulator now.

## Goal

Hopefully support following commands and resource types enough to operate kubernetes as kubectl.

#### Commands:

* [x] `get`            Display one or many resources
* [x] `describe`       Show details of a specific resource or group of resources
* [x] `create`         Create a resource by filename or stdin
* [x] `replace`        Replace a resource by filename or stdin.
* [x] `patch`          Update field(s) of a resource using strategic merge patch.
* [x] `delete`         Delete resources by filenames, stdin, resources and names, or by resources and label selector.
* [x] `edit`           Edit a resource on the server
* [x] `apply`          Apply a configuration to a resource by filename or stdin
* [x] `namespace`      SUPERSEDED: Set and view the current Kubernetes namespace
* [ ] `logs`           Print the logs for a container in a pod.
* [ ] `rolling-update` Perform a rolling update of the given ReplicationController.
* [ ] `scale`          Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job.
* [x] `cordon`         Mark node as unschedulable
* [x] `drain`          Drain node in preparation for maintenance
* [x] `uncordon`       Mark node as schedulable
* [ ] `attach`         Attach to a running container.
* [ ] `exec`           Execute a command in a container.
* [ ] `port-forward`   Forward one or more local ports to a pod.
* [ ] `proxy`          Run a proxy to the Kubernetes API server
* [ ] `run`            Run a particular image on the cluster.
* [ ] `expose`         Take a replication controller, service, or pod and expose it as a new Kubernetes Service
* [ ] `autoscale`      Auto-scale a Deployment, ReplicaSet, or ReplicationController
* [ ] `rollout`        rollout manages a deployment
* [ ] `label`          Update the labels on a resource
* [ ] `annotate`       Update the annotations on a resource
* [ ] `config`         config modifies kubeconfig files
* [x] `cluster-info`   Display cluster info
* [x] `api-versions`   Print the supported API versions on the server, in the form of "group/version".
* [x] `version`        Print the client and server version information.
* [x] `explain`        Documentation of resources.
* [ ] `convert`        Convert config files between different API versions

#### Resource Types:

* [ ] `clusters`
* [x] `componentstatuses` aka `cs`
* [x] `configmaps` aka `cm`
* [x] `daemonsets` aka `ds`
* [x] `deployments` aka `deploy`
* [x] `endpoints` aka `ep`
* [x] `events` aka `ev`
* [ ] `horizontalpodautoscalers` aka `hpa`
* [x] `ingresses` aka `ing`
* [ ] `jobs`
* [x] `limitranges` aka `limits`
* [x] `namespaces` aka `ns`
* [ ] `networkpolicies`
* [x] `nodes` aka `no`
* [x] `persistentvolumeclaims` aka `pvc`
* [x] `persistentvolumes` aka `pv`
* [x] `pods`
* [x] `podsecuritypolicies` aka `psp`
* [x] `podtemplates`
* [x] `replicasets` aka `rs`
* [x] `replicationcontrollers` aka `rc`
* [x] `resourcequotas` aka `quota`
* [x] `secrets`
* [x] `serviceaccounts` aka `sa`
* [x] `services` aka `svc`
* [ ] `statefulsets`
* [ ] `storageclasses`
* [ ] `thirdpartyresources`

## LICENSE

This software is licensed under the MIT License (See [LICENSE](./LICENSE)).
