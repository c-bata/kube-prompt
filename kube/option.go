package kube

import (
	"strings"

	prompt "github.com/c-bata/go-prompt"
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
	case "namespace":
		suggests = flagGlobal
	case "logs":
		suggests = logsOptions
	case "rolling-update":
		suggests = flagRollingUpdate
	case "scale", "resize":
		suggests = scaleOptions
	case "attach":
		suggests = attachOptions
	case "exec":
		suggests = execOptions
	case "port-forward":
		suggests = append(flagPortForward, flagGlobal...)
	case "proxy":
		suggests = proxyOptions
	case "run", "run-container":
		suggests = runOptions
	case "expose":
		suggests = append(exposeOptions, flagGlobal...)
	case "auto-scale":
		suggests = append(flagAutoScale, flagGlobal...)
	case "rollout":
		suggests = flagGlobal
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "history":
				suggests = append(suggests, flagRolloutHistory...)
			case "pause":
				suggests = append(suggests, flagRolloutPause...)
			case "resume":
				suggests = append(suggests, flagRolloutResume...)
			case "undo":
				suggests = append(suggests, flagRolloutUndo...)
			}
		}
	case "label":
		suggests = labelOptions
	case "cluster-info":
		suggests = flagClusterInfo
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
		suggests = flagGlobal
		if len(commandArgs) >= 2 {
			switch commandArgs[1] {
			case "no", "node", "nodes":
				suggests = append(suggests, flagTopNode...)
			case "po", "pod", "pods":
				suggests = append(suggests, flagTopPod...)
			}
		}
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "view":
				suggests = flagConfigView
			case "set-cluster":
				suggests = flagConfigSetCluster
			case "set-credentials":
				suggests = flagConfigSetCredentials
			case "set-context":
				suggests = flagConfigSetContext
			case "set":
				suggests = flagConfigSet
			case "unset":
				suggests = flagConfigUnset
			case "current-context":
				suggests = flagConfigCurrentContext
			case "use-context":
				suggests = flagConfigUseContext
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

var flagGlobal = []prompt.Suggest{
	{Text: "--alsologtostderr", Description: "log to standard error as well as files"},
	{Text: "--certificate-authority", Description: "Path to a cert. file for the certificate authority."},
	{Text: "--client-certificate", Description: "Path to a client certificate file for TLS."},
	{Text: "--client-key", Description: "Path to a client key file for TLS."},
	{Text: "--cluster", Description: "The name of the kubeconfig cluster to use"},
	{Text: "--context", Description: "The name of the kubeconfig context to use"},
	{Text: "--insecure-skip-tls-verify", Description: "If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure."},
	{Text: "--kubeconfig", Description: "Path to the kubeconfig file to use for CLI requests."},
	{Text: "--log-backtrace-at", Description: "when logging hits line file:N, emit a stack trace"},
	{Text: "--log-dir", Description: "If non-empty, write log files in this directory"},
	{Text: "--log-flush-frequency", Description: "Maximum number of seconds between log flushes"},
	{Text: "--logtostderr", Description: "log to standard error instead of files"},
	{Text: "--match-server-version", Description: "Require server version to match client version"},
	{Text: "--namespace", Description: "If present, the namespace scope for this CLI request."},
	{Text: "--password", Description: "Password for basic authentication to the API server."},
	{Text: "-s", Description: "The address and port of the Kubernetes API server"},
	{Text: "--server", Description: "The address and port of the Kubernetes API server"},
	{Text: "--stderrthreshold", Description: "logs at or above this threshold go to stderr"},
	{Text: "--token", Description: "Bearer token for authentication to the API server."},
	{Text: "--user", Description: "The name of the kubeconfig user to use"},
	{Text: "--username", Description: "Username for basic authentication to the API server."},
	{Text: "--v", Description: "log level for V logs"},
	{Text: "--vmodule", Description: "comma-separated list of pattern=N settings for file-filtered logging"},
}
var flagRollingUpdate = []prompt.Suggest{
	{Text: "--container", Description: "Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod"},
	{Text: "--deployment-label-key", Description: "The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise"},
	{Text: "--dry-run", Description: "If true, print out the changes that would be made, but don't actually make them."},
	{Text: "-f", Description: "Filename or URL to file to use to create the new replication controller."},
	{Text: "--filename", Description: "Filename or URL to file to use to create the new replication controller."},
	{Text: "--image", Description: "Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f"},
	{Text: "--no-headers", Description: "When using the default output, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template and jsonpath template."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template and jsonpath template."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--poll-interval", Description: "Time delay between polling for replication controller status after the update. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'."},
	{Text: "--rollback", Description: "If true, this is a request to abort an existing rollout that is partially rolled out. It effectively reverses current and next and runs a rollout"},
	{Text: "--schema-cache-dir", Description: "If non-empty, load/store cached API schemas in this directory, default is '$HOME/.kube/schema'"},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--timeout", Description: "Max time to wait for a replication controller to update before giving up. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'."},
	{Text: "--update-period", Description: "Time to wait between updating pods. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'."},
	{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var flagPortForward = []prompt.Suggest{
	{Text: "-p", Description: "Pod name"},
	{Text: "--pod", Description: "Pod name"},
}

var flagRolloutHistory = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--revision", Description: "See the details, including podTemplate of the revision specified"},
}

var flagRolloutPause = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
}

var flagRolloutResume = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
}

var flagRolloutUndo = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to get from a server."},
	{Text: "--to-revision", Description: "The revision to rollback to. Default to 0 (last revision)."},
}

var flagClusterInfo = []prompt.Suggest{
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
}

var flagAutoScale = []prompt.Suggest{
	{Text: "--cpu-percent", Description: "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, the server will apply a default value."},
	{Text: "--dry-run", Description: "If true, only print the object that would be sent, without creating it."},
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to autoscale."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to autoscale."},
	{Text: "--generator", Description: "The name of the API generator to use. Currently there is only 1 generator."},
	{Text: "--max", Description: "The upper limit for the number of pods that can be set by the autoscaler. Required."},
	{Text: "--min", Description: "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value."},
	{Text: "--name", Description: "The name for the newly created object. If not specified, the name of the input resource will be used."},
	{Text: "--no-headers", Description: "When using the default output, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var flagTopNode = []prompt.Suggest{
	{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in."},
	{Text: "--heapster-port", Description: "Port name in service to use."},
	{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as."},
	{Text: "--heapster-service", Description: "Name of Heapster service."},
	{Text: "-l", Description: "Selector (label query) to filter on"},
	{Text: "--selector", Description: "Selector (label query) to filter on"},
}

var flagTopPod = []prompt.Suggest{
	{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	{Text: "--containers", Description: "If present, print usage of containers within a pod."},
	{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in."},
	{Text: "--heapster-port", Description: "Port name in service to use."},
	{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as."},
	{Text: "--heapster-service", Description: "Name of Heapster service."},
	{Text: "-l", Description: "Selector (label query) to filter on"},
	{Text: "--selector", Description: "Selector (label query) to filter on"},
}

var flagConfigView = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "--flatten", Description: "flatten the resulting kubeconfig file into self-contained output (useful for creating portable kubeconfig files)"},
	{Text: "--merge", Description: "merge the full hierarchy of kubeconfig files"},
	{Text: "--minify", Description: "remove all information not used by current-context from the output"},
	{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--raw", Description: "display raw byte data"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var flagConfigSetCluster = []prompt.Suggest{
	{Text: "--api-version", Description: "api-version for the cluster entry in kubeconfig"},
	{Text: "--certificate-authority", Description: "path to certificate-authority for the cluster entry in kubeconfig"},
	{Text: "--embed-certs", Description: "embed-certs for the cluster entry in kubeconfig"},
	{Text: "--insecure-skip-tls-verify", Description: "insecure-skip-tls-verify for the cluster entry in kubeconfig"},
	{Text: "--server", Description: "server for the cluster entry in kubeconfig"},
}

var flagConfigSetCredentials = []prompt.Suggest{
	{Text: "--client-certificate", Description: "path to client-certificate for the user entry in kubeconfig"},
	{Text: "--client-key", Description: "path to client-key for the user entry in kubeconfig"},
	{Text: "--embed-certs", Description: "embed client cert/key for the user entry in kubeconfig"},
	{Text: "--password", Description: "password for the user entry in kubeconfig"},
	{Text: "--token", Description: "token for the user entry in kubeconfig"},
	{Text: "--username", Description: "username for the user entry in kubeconfig"},
}

var flagConfigSetContext = []prompt.Suggest{
	{Text: "--cluster", Description: "cluster for the context entry in kubeconfig"},
	{Text: "--namespace", Description: "namespace for the context entry in kubeconfig"},
	{Text: "--user", Description: "user for the context entry in kubeconfig"},
}

var flagConfigSet = []prompt.Suggest{}

var flagConfigUnset = []prompt.Suggest{}

var flagConfigCurrentContext = []prompt.Suggest{}

var flagConfigUseContext = []prompt.Suggest{}
