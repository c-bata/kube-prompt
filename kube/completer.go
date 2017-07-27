package kube

import (
	"strings"

	"github.com/c-bata/go-prompt-toolkit"
)

func Completer(s string) []prompt.Completion {
	if s == "" {
		return []prompt.Completion{}
	}
	args := strings.Split(s, " ")
	l := len(args)

	if strings.HasPrefix(args[l-1], "-") {
		return optionCompleter(args, strings.HasPrefix(args[l-1], "--"))
	}

	return argumentsCompleter(excludeOptions(args))
}

func strToCompletionList(x []string) []prompt.Completion {
	l := len(x)
	y := make([]prompt.Completion, l)
	for i := 0; i < l; i++ {
		y[i] = prompt.Completion{Text: x[i]}
	}
	return y
}

var commands = []prompt.Completion{
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
	// {Text: "attach", Description: "Attach to a running container."},  // still not supported
	// {Text: "exec", Description: "Execute a command in a container."}, // still not supported
	// {Text: "port-forward", Description: "Forward one or more local ports to a pod."}, // still not supported
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
}

func argumentsCompleter(args []string) []prompt.Completion {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "get":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), args[1], true)
		}
	case "describe":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), second, true)
		}

		third := args[2]
		switch second {
		case "configmaps", "cm":
			return prompt.FilterContains(getConfigMapCompletions(), third, true)
		case "po", "pod", "pods":
			return prompt.FilterContains(getPodCompletions(), third, true)
		case "deploy", "deployments":
			return prompt.FilterContains(getDeploymentNames(), third, true)
		case "no", "nodes":
			return prompt.FilterContains(getNodeCompletions(), third, true)
		case "secrets":
			return prompt.FilterContains(getSecretCompletions(), third, true)
		case "sa", "serviceaccounts":
			return prompt.FilterContains(getServiceAccountCompletions(), third, true)
		}
	case "create":
		subcommands := []prompt.Completion{
			{Text: "configmap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "deployment", Description: "Create a deployment with the specified name."},
			{Text: "namespace", Description: "Create a namespace with the specified name"},
			{Text: "quota", Description: "Create a quota with the specified name."},
			{Text: "secret", Description: "Create a secret using specified subcommand"},
			{Text: "service", Description: "Create a service using specified subcommand."},
			{Text: "serviceaccount", Description: "Create a service account with the specified name"},
		}
		return prompt.FilterHasPrefix(subcommands, args[1], true)
	case "replace":
	case "patch":
	case "delete":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), args[1], true)
	case "edit":
	case "apply":
	case "namespace":
	case "logs":
	case "rolling-update":
	case "scale":
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	case "uncordon":
		return prompt.FilterHasPrefix(getNodeCompletions(), args[1], true)
	//case "attach": // still not supported
	//case "exec":   // still not supported
	//case "port-forward": // still not supported
	case "proxy":
	case "run":
	case "expose":
	case "autoscale":
	case "rollout":
	case "label":
	case "annotate":
	case "config":
		subCommands := []prompt.Completion{
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
	case "cluster-info":
		subCommands := []prompt.Completion{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "api-versions":
	case "version":
	case "explain":
		return prompt.FilterHasPrefix(strToCompletionList(resourceTypes), args[1], true)
	case "convert":
	default:
		return []prompt.Completion{}
	}
	return []prompt.Completion{}
}

