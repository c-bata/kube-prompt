# kube-prompt

**kube-prompt** is powerful interactive command lines for Kubernetes.

## Usage

![example](./_resources/kube-prompt.gif)

Basically, kube-prompt wraps kubectl. So if you know the usage of kubectl, you can also use kube-prompt.
Binaries are available from:

* macOS (darwin)
* Linux

Windows support is still not because [go-prompt-toolkit](https://github.com/c-bata/go-prompt-toolkit) is only supports VT100 console emulator now.

## Goal

Hopefully support following commands and resource types enough to operate kubernetes as kubectl.

#### Commands:

* [x] `get`            Display one or many resources
* [x] `describe`       Show details of a specific resource or group of resources
* [x] `create`         Create a resource by filename or stdin
* [ ] `replace`        Replace a resource by filename or stdin.
* [ ] `patch`          Update field(s) of a resource using strategic merge patch.
* [ ] `delete`         Delete resources by filenames, stdin, resources and names, or by resources and label selector.
* [ ] `edit`           Edit a resource on the server
* [ ] `apply`          Apply a configuration to a resource by filename or stdin
* [ ] `namespace`      SUPERSEDED: Set and view the current Kubernetes namespace
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
* [ ] `componentstatuses` aka `cs`
* [ ] `configmaps` aka `cm`
* [ ] `daemonsets` aka `ds`
* [x] `deployments` aka `deploy`
* [ ] `endpoints` aka `ep`
* [ ] `events` aka `ev`
* [ ] `horizontalpodautoscalers` aka `hpa`
* [ ] `ingresses` aka `ing`
* [ ] `jobs`
* [ ] `limitranges` aka `limits`
* [ ] `namespaces` aka `ns`
* [ ] `networkpolicies`
* [x] `nodes` aka `no`
* [ ] `persistentvolumeclaims` aka `pvc`
* [ ] `persistentvolumes` aka `pv`
* [x] `pods`
* [ ] `podsecuritypolicies` aka `psp`
* [ ] `podtemplates`
* [ ] `replicasets` aka `rs`
* [ ] `replicationcontrollers` aka `rc`
* [ ] `resourcequotas` aka `quota`
* [ ] `secrets`
* [ ] `serviceaccounts` aka `sa`
* [ ] `services` aka `svc`
* [ ] `statefulsets`
* [ ] `storageclasses`
* [ ] `thirdpartyresources`

## LICENSE

This software is licensed under the MIT License (See [LICENSE](./LICENSE)).
