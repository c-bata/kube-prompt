package kube

import (
	"strings"

	"github.com/c-bata/go-prompt-toolkit"
)

func Completer(s string) []string {
	args := strings.Split(s, " ")
	if len(args) == 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	if len(args) == 2 {
		return secondArgsCompleter(args[0], args[1])
	}

	if len(args) == 3 {
		return thirdArgsCompleter(args[0], args[1], args[2])
	}

	return []string{}
}

var commands = []string{
	"get",
	"describe",
	"create",
	"replace",
	"patch",
	"delete",
	"edit",
	"apply",
	"namespace",
	"logs",
	"rolling-update",
	"scale",
	"cordon",
	"drain",
	"uncordon",
	"attach",
	"exec",
	"port-forward",
	"proxy",
	"run",
	"expose",
	"autoscale",
	"rollout",
	"label",
	"annotate",
	"config",
	"cluster-info",
	"api-versions",
	"version",
	"explain",
	"convert",
}

var operationSpecies = []string{
	"pods",
	"replicationcontroller",
	"rc",
	"services",
}

func secondArgsCompleter(first, second string) []string {
	switch first {
	case "get":
		return prompt.FilterHasPrefix(operationSpecies, second, true)
	case "describe":
		return prompt.FilterHasPrefix(operationSpecies, second, true)
	case "create":
		return prompt.FilterHasPrefix(operationSpecies, second, true)
	case "replace":
	case "patch":
	case "delete":
	case "edit":
	case "apply":
	case "namespace":
	case "logs":
	case "rolling-update":
	case "scale":
	case "cordon":
	case "drain":
	case "uncordon":
	case "attach":
	case "exec":
	case "port-forward":
	case "proxy":
	case "run":
	case "expose":
	case "autoscale":
	case "rollout":
	case "label":
	case "annotate":
	case "config":
	case "cluster-info":
	case "api-versions":
	case "version":
	case "explain":
	case "convert":
	default:
		return []string{}
	}
	return []string{}
}

func thirdArgsCompleter(first, second, third string) []string {
	if first == "get" && second == "pods" {
		return prompt.FilterContains(getPods(), third, true)
	}
	if first == "describe" && second == "pods" {
		return prompt.FilterContains(getPods(), third, true)
	}
	return []string{}
}

// utilities

func filterHasPrefix(completions []string, sub string) []string {
	if sub == "" {
		return completions
	}
	ret := make([]string, 0, len(completions))
	for _, n := range completions {
		if strings.HasPrefix(n, sub) {
			ret = append(ret, n)
		}
	}
	return ret
}
