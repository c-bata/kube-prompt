package kube

import (
	"context"

	"github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "replace", Description: "Replace a resource by filename or stdin."},
	{Text: "patch", Description: "Update field(s) of a resource using strategic merge patch."},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "namespace", Description: "SUPERSEDED: Set and view the current Kubernetes namespace"},
	{Text: "logs", Description: "Print the logs for a container in a pod."},
	{Text: "rolling-update", Description: "Perform a rolling update of the given ReplicationController."},
	{Text: "scale", Description: "Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job."},
	{Text: "cordon", Description: "Mark node as unschedulable"},
	{Text: "drain", Description: "Drain node in preparation for maintenance"},
	{Text: "uncordon", Description: "Mark node as schedulable"},
	{Text: "attach", Description: "Attach to a running container."},
	{Text: "exec", Description: "Execute a command in a container."},
	{Text: "port-forward", Description: "Forward one or more local ports to a pod."},
	{Text: "proxy", Description: "Run a proxy to the Kubernetes API server"},
	{Text: "run", Description: "Run a particular image on the cluster."},
	{Text: "expose", Description: "Take a replication controller, service, or pod and expose it as a new Kubernetes Service"},
	{Text: "autoscale", Description: "Auto-scale a Deployment, ReplicaSet, or ReplicationController"},
	{Text: "rollout", Description: "rollout manages a deployment"},
	{Text: "label", Description: "Update the labels on a resource"},
	{Text: "annotate", Description: "Update the annotations on a resource"},
	{Text: "config", Description: "config modifies kubeconfig files"},
	{Text: "cluster-info", Description: "Display cluster info"},
	{Text: "api-versions", Description: "Print the supported API versions on the server, in the form of 'group/version'."},
	{Text: "version", Description: "Print the client and server version information."},
	{Text: "explain", Description: "Documentation of resources."},
	{Text: "convert", Description: "Convert config files between different API versions"},
	{Text: "top", Description: "Display Resource (CPU/Memory/Storage) usage"},

	// Custom command.
	{Text: "exit", Description: "Exit this program"},
}

var resourceTypes = []prompt.Suggest{
	{Text: "clusters"}, // valid only for federation apiservers
	{Text: "componentstatuses"},
	{Text: "configmaps"},
	{Text: "daemonsets"},
	{Text: "deployments"},
	{Text: "endpoints"},
	{Text: "events"},
	{Text: "horizontalpodautoscalers"},
	{Text: "ingresses"},
	{Text: "jobs"},
	{Text: "cronjobs"},
	{Text: "limitranges"},
	{Text: "namespaces"},
	{Text: "networkpolicies"},
	{Text: "nodes"},
	{Text: "persistentvolumeclaims"},
	{Text: "persistentvolumes"},
	{Text: "pod"},
	{Text: "podsecuritypolicies"},
	{Text: "podtemplates"},
	{Text: "replicasets"},
	{Text: "replicationcontrollers"},
	{Text: "resourcequotas"},
	{Text: "secrets"},
	{Text: "serviceaccounts"},
	{Text: "services"},
	{Text: "statefulsets"},
	{Text: "storageclasses"},
	{Text: "thirdpartyresources"},

	// aliases
	{Text: "cs"},
	{Text: "cm"},
	{Text: "ds"},
	{Text: "deploy"},
	{Text: "ep"},
	{Text: "hpa"},
	{Text: "ing"},
	{Text: "limits"},
	{Text: "ns"},
	{Text: "no"},
	{Text: "pvc"},
	{Text: "pv"},
	{Text: "po"},
	{Text: "psp"},
	{Text: "rs"},
	{Text: "rc"},
	{Text: "quota"},
	{Text: "sa"},
	{Text: "svc"},
}

func (c *Completer) argumentsCompleter(ctx context.Context, namespace string, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "get":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "componentstatuses"},
				{Text: "configmaps"},
				{Text: "daemonsets"},
				{Text: "deployments"},
				{Text: "endpoints"},
				{Text: "events"},
				{Text: "horizontalpodautoscalers"},
				{Text: "ingresses"},
				{Text: "jobs"},
				{Text: "cronjobs"},
				{Text: "limitranges"},
				{Text: "namespaces"},
				{Text: "networkpolicies"},
				{Text: "nodes"},
				{Text: "persistentvolumeclaims"},
				{Text: "persistentvolumes"},
				{Text: "pod"},
				{Text: "podsecuritypolicies"},
				{Text: "podtemplates"},
				{Text: "replicasets"},
				{Text: "replicationcontrollers"},
				{Text: "resourcequotas"},
				{Text: "secrets"},
				{Text: "serviceaccounts"},
				{Text: "services"},
				{Text: "statefulsets"},
				{Text: "storageclasses"},
				{Text: "thirdpartyresources"},
				// aliases
				{Text: "cs"},
				{Text: "cm"},
				{Text: "ds"},
				{Text: "deploy"},
				{Text: "ep"},
				{Text: "hpa"},
				{Text: "ing"},
				{Text: "limits"},
				{Text: "ns"},
				{Text: "no"},
				{Text: "pvc"},
				{Text: "pv"},
				{Text: "po"},
				{Text: "psp"},
				{Text: "rs"},
				{Text: "rc"},
				{Text: "quota"},
				{Text: "sa"},
				{Text: "svc"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(ctx, c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(ctx, c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(ctx, c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(ctx, c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(ctx, c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(ctx, c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(ctx, c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(ctx, c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(ctx, c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(ctx, c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(ctx, c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(ctx, c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(ctx, c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(ctx, c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(ctx, c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(ctx, c.client, namespace), third, true)
			}
		}
	case "describe":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(ctx, c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(ctx, c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(ctx, c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(ctx, c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(ctx, c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(ctx, c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(ctx, c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(ctx, c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(ctx, c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(ctx, c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(ctx, c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(ctx, c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(ctx, c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(ctx, c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(ctx, c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(ctx, c.client, namespace), third, true)
			}
		}
	case "create":
		subcommands := []prompt.Suggest{
			{Text: "configmap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "deployment", Description: "Create a deployment with the specified name."},
			{Text: "namespace", Description: "Create a namespace with the specified name"},
			{Text: "quota", Description: "Create a quota with the specified name."},
			{Text: "secret", Description: "Create a secret using specified subcommand"},
			{Text: "service", Description: "Create a service using specified subcommand."},
			{Text: "serviceaccount", Description: "Create a service account with the specified name"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "delete":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(ctx, c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(ctx, c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(ctx, c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(ctx, c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(ctx, c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(ctx, c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(ctx, c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(ctx, c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(ctx, c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(ctx, c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(ctx, c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(ctx, c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(ctx, c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(ctx, c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(ctx, c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(ctx, c.client, namespace), third, true)
			}
		}
	case "edit":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, args[1], true)
		}

		if len(args) == 3 {
			third := args[2]
			switch args[1] {
			case "componentstatuses", "cs":
				return prompt.FilterContains(getComponentStatusCompletions(ctx, c.client), third, true)
			case "configmaps", "cm":
				return prompt.FilterContains(getConfigMapSuggestions(ctx, c.client, namespace), third, true)
			case "daemonsets", "ds":
				return prompt.FilterContains(getDaemonSetSuggestions(ctx, c.client, namespace), third, true)
			case "deploy", "deployments":
				return prompt.FilterContains(getDeploymentSuggestions(ctx, c.client, namespace), third, true)
			case "endpoints", "ep":
				return prompt.FilterContains(getEndpointsSuggestions(ctx, c.client, namespace), third, true)
			case "ingresses", "ing":
				return prompt.FilterContains(getIngressSuggestions(ctx, c.client, namespace), third, true)
			case "limitranges", "limits":
				return prompt.FilterContains(getLimitRangeSuggestions(ctx, c.client, namespace), third, true)
			case "namespaces", "ns":
				return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
			case "no", "nodes":
				return prompt.FilterContains(getNodeSuggestions(ctx, c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumeclaims", "pvc":
				return prompt.FilterContains(getPersistentVolumeClaimSuggestions(ctx, c.client, namespace), third, true)
			case "persistentvolumes", "pv":
				return prompt.FilterContains(getPersistentVolumeSuggestions(ctx, c.client), third, true)
			case "podtemplates":
				return prompt.FilterContains(getPodTemplateSuggestions(ctx, c.client, namespace), third, true)
			case "replicasets", "rs":
				return prompt.FilterContains(getReplicaSetSuggestions(ctx, c.client, namespace), third, true)
			case "replicationcontrollers", "rc":
				return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), third, true)
			case "resourcequotas", "quota":
				return prompt.FilterContains(getResourceQuotasSuggestions(ctx, c.client, namespace), third, true)
			case "secrets":
				return prompt.FilterContains(getSecretSuggestions(ctx, c.client, namespace), third, true)
			case "sa", "serviceaccounts":
				return prompt.FilterContains(getServiceAccountSuggestions(ctx, c.client, namespace), third, true)
			case "svc", "services":
				return prompt.FilterContains(getServiceSuggestions(ctx, c.client, namespace), third, true)
			case "job", "jobs":
				return prompt.FilterContains(getJobSuggestions(ctx, c.client, namespace), third, true)
			}
		}

	case "namespace":
		if len(args) == 2 {
			return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), args[1], true)
		}
	case "logs":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), args[1], true)
		}
	case "rolling-update", "rollingupdate":
		if len(args) == 2 {
			return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), args[1], true)
		} else if len(args) == 3 {
			return prompt.FilterContains(getReplicationControllerSuggestions(ctx, c.client, namespace), args[2], true)
		}
	case "scale", "resize":
		if len(args) == 2 {
			// Deployment, ReplicaSet, Replication Controller, or Job.
			r := getDeploymentSuggestions(ctx, c.client, namespace)
			r = append(r, getReplicaSetSuggestions(ctx, c.client, namespace)...)
			r = append(r, getReplicationControllerSuggestions(ctx, c.client, namespace)...)
			return prompt.FilterContains(r, args[1], true)
		}
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	case "uncordon":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(getNodeSuggestions(ctx, c.client), args[1], true)
		}
	case "attach":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), args[1], true)
		}
	case "exec":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), args[1], true)
		}
	case "port-forward":
		if len(args) == 2 {
			return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), args[1], true)
		}
		if len(args) == 3 {
			return prompt.FilterHasPrefix(getPortsFromPodName(namespace, args[1]), args[2], true)
		}
	case "rollout":
		subCommands := []prompt.Suggest{
			{Text: "history", Description: "view rollout history"},
			{Text: "pause", Description: "Mark the provided resource as paused"},
			{Text: "resume", Description: "Resume a paused resource"},
			{Text: "undo", Description: "undoes a previous rollout"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "annotate":
	case "config":
		subCommands := []prompt.Suggest{
			{Text: "current-context", Description: "Displays the current-context"},
			{Text: "delete-cluster", Description: "Delete the specified cluster from the kubeconfig"},
			{Text: "delete-context", Description: "Delete the specified context from the kubeconfig"},
			{Text: "get-clusters", Description: "Display clusters defined in the kubeconfig"},
			{Text: "get-contexts", Description: "Describe one or many contexts"},
			{Text: "set", Description: "Sets an individual value in a kubeconfig file"},
			{Text: "set-cluster", Description: "Sets a cluster entry in kubeconfig"},
			{Text: "set-context", Description: "Sets a context entry in kubeconfig"},
			{Text: "set-credentials", Description: "Sets a user entry in kubeconfig"},
			{Text: "unset", Description: "Unsets an individual value in a kubeconfig file"},
			{Text: "use-context", Description: "Sets the current-context in a kubeconfig file"},
			{Text: "view", Description: "Display merged kubeconfig settings or a specified kubeconfig file"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
		if len(args) == 3 {
			third := args[2]
			switch args[1] {
			case "use-context":
				return prompt.FilterContains(getContextSuggestions(), third, true)
			}
		}
	case "cluster-info":
		subCommands := []prompt.Suggest{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "explain":
		return prompt.FilterHasPrefix(resourceTypes, args[1], true)
	case "top":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "nodes"},
				{Text: "pod"},
				// aliases
				{Text: "no"},
				{Text: "po"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "no", "node", "nodes":
				return prompt.FilterContains(getNodeSuggestions(ctx, c.client), third, true)
			case "po", "pod", "pods":
				return prompt.FilterContains(getPodSuggestions(ctx, c.client, namespace), third, true)
			}
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
