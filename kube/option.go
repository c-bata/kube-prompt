package kube

import (
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs := excludeOptions(args)
	switch commandArgs[0] {
	case "get":
		suggests = getOptions
	case "describe":
		suggests = describeOptions
	case "create":
		suggests = createOptions
	case "replace":
		suggests = replaceOptions
	case "patch":
		suggests = patchOptions
	case "delete":
		suggests = deleteOptions
	case "edit":
		suggests = editOptions
	case "apply":
		suggests = applyOptions
	case "logs":
		suggests = logsOptions
	case "rolling-update":
		suggests = rollingUpdateOptions
	case "scale", "resize":
		suggests = scaleOptions
	case "attach":
		suggests = attachOptions
	case "exec":
		suggests = execOptions
	case "port-forward":
		suggests = portForwardOptions
	case "proxy":
		suggests = proxyOptions
	case "run", "run-container":
		suggests = runOptions
	case "expose":
		suggests = exposeOptions
	case "autoscale":
		suggests = autoscaleOptions
	case "rollout":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "history":
				suggests = rolloutHistoryOptions
			case "pause":
				suggests = rolloutPauseOptions
			case "resume":
				suggests = rolloutResumeOptions
			case "status":
				suggests = rolloutStatusOptions
			case "undo":
				suggests = rolloutUndoOptions
			}
		}
	case "label":
		suggests = labelOptions
	case "cluster-info":
		suggests = clusterInfoOptions
	case "explain":
		suggests = explainOptions
	case "cordon":
		suggests = cordonOptions
	case "drain":
		suggests = drainOptions
	case "uncordon":
		suggests = uncordonOptions
	case "annotate":
		suggests = annotateOptions
	case "convert":
		suggests = convertOptions
	case "top":
		if len(commandArgs) >= 2 {
			switch commandArgs[1] {
			case "no", "node", "nodes":
				suggests = topNodeOptions
			case "po", "pod", "pods":
				suggests = topPodOptions
			}
		}
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "get-contexts":
				suggests = configGetContextsOptions
			case "view":
				suggests = configViewOptions
			case "set-cluster":
				suggests = configSetClusterOptions
			case "set-credentials":
				suggests = configSetCredentialsOptions
			case "set":
				suggests = configSetOptions
			}
		}
	default:
		suggests = optionHelp
	}

	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

/* Option arguments */

var yamlFileCompleter = completer.FilePathCompleter{
	IgnoreCase: true,
	Filter: func(fi os.FileInfo) bool {
		if fi.IsDir() {
			return true
		}
		if strings.HasSuffix(fi.Name(), ".yaml") || strings.HasSuffix(fi.Name(), ".yml") {
			return true
		}
		return false
	},
}

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}
	switch cmd {
	case "get", "describe", "create", "delete", "replace", "patch",
		"edit", "apply", "expose", "rolling-update", "rollout",
		"label", "annotate", "scale", "convert", "autoscale", "top":
		switch option {
		case "-f", "--filename":
			return yamlFileCompleter.Complete(d), true
		case "-n", "--namespace":
			return getNameSpaceSuggestions(), true
		}
	}
	return []prompt.Suggest{}, false
}
