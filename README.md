# kube-prompt

An interactive kubernetes client featuring auto-complete using [go-prompt](https://github.com/c-bata/go-prompt).

![demo](./_resources/kube-prompt.gif)

kube-prompt's command is same with kubectl (because basically this is just wrapper).
So it doesn't require the additional cost to learn the usage of kube-prompt.

## Installation

#### Binary installation

Binaries are available from [github release](https://github.com/c-bata/kube-prompt/releases).

```
# macOS (darwin)
curl -L https://github.com/c-bata/kube-prompt/releases/download/v1.0.0/kube-prompt_v0.1.0_darwin_amd64.zip
unzip kube-prompt_v1.0.0_darwin_amd64.zip

# Linux
curl -L https://github.com/c-bata/kube-prompt/releases/download/v1.0.0/kube-prompt_v0.1.0_linux_amd64.zip
unzip kube-prompt_v1.0.0_linux_amd64.zip

# After that please put executable to your PATH:
chmod +x kube-prompt
sudo mv ./kube-prompt /usr/local/bin/kube-prompt
```

#### Build from source

```console
$ go get -u github.com/golang/dep/cmd/dep
$ dep ensure # download dependency package
$ go build .
```

To create a multi-platform binary, use the cross command via make:

```console
$ make cross
```

## Goal

Hopefully support following commands enough to operate kubernetes.

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
