package kube

import (
	"strings"

	"github.com/c-bata/go-prompt"
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
		suggests = append(flagGet, flagGlobal...)
	case "describe":
		suggests = append(flagDescribe, flagGlobal...)
	case "create":
		suggests = append(flagCreate, flagGlobal...)
	case "replace":
		suggests = append(flagReplace, flagGlobal...)
	case "patch":
		suggests = append(flagPatch, flagGlobal...)
	case "delete":
		suggests = append(flagDelete, flagGlobal...)
	case "edit":
		suggests = append(flagEdit, flagGlobal...)
	case "apply":
		suggests = append(flagApply, flagGlobal...)
	case "namespace":
		suggests = flagGlobal
	case "logs":
		suggests = append(flagLogs, flagGlobal...)
	case "rolling-update", "rollingupdate":
		suggests = append(flagRollingUpdate, flagGlobal...)
	case "scale", "resize":
		suggests = append(flagScale, flagGlobal...)
	case "attach":
		suggests = append(flagAttach, flagGlobal...)
	case "exec":
		suggests = append(flagExec, flagGlobal...)
	case "port-forward":
		suggests = append(flagPortForward, flagGlobal...)
	case "proxy":
		suggests = append(flagProxy, flagGlobal...)
	case "run", "run-container":
		suggests = append(flagRun, flagGlobal...)
	case "expose":
		suggests = append(flagExpose, flagGlobal...)
	case "auto-scale":
		suggests = append(flagAutoScale, flagGlobal...)
	case "cluster-info":
		suggests = flagClusterInfo
	case "explain":
		suggests = flagExplain
	case "cordon":
		suggests = append(flagCordon, flagGlobal...)
	case "drain":
		suggests = flagCordon
	case "uncordon":
		suggests = optionHelp
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "view":
				suggests = flagConfigView
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

var flagGet = []prompt.Suggest{
	{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "--export", Description: "If true, use 'export' for the resources.  Exported resources are stripped of cluster-specific information."},
	{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "-L", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	{Text: "--label-columns", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--raw", Description: "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file."},
	{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "-l", Description: "Selector (label query) to filter on"},
	{Text: "--selector", Description: "Selector (label query) to filter on"},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-kind", Description: "If present, list the resource type for the requested object(s)."},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates."},
	{Text: "-w", Description: "After listing/getting the requested object, watch for changes."},
	{Text: "--watch", Description: "After listing/getting the requested object, watch for changes."},
	{Text: "--watch-only", Description: "Watch for changes to the requested object(s), without listing/getting first."},
}

var flagDescribe = []prompt.Suggest{
	{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	{Text: "-f", Description: "Filename, directory, or URL to files containing the resource to describe"},
	{Text: "--filename", Description: "Filename, directory, or URL to files containing the resource to describe"},
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "-l", Description: "Selector (label query) to filter on"},
	{Text: "--selector", Description: "Selector (label query) to filter on"},
	{Text: "--show-events", Description: "If true, display events related to the described object."},
}

var flagCreate = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	{Text: "--edit", Description: "Edit the API resource before creating"},
	{Text: "-f", Description: "Filename, directory, or URL to files to use to create the resource"},
	{Text: "--filename", Description: "Filename, directory, or URL to files to use to create the resource"},
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=..."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "--schema-cache-dir", Description: "If non-empty, load/store cached API schemas in this directory, default is '$HOME/.kube/schema'"},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
	{Text: "--windows-line-endings", Description: "Only relevant if --edit=true. Use Windows line-endings (default Unix line-endings)"},
}

var flagReplace = []prompt.Suggest{
	{Text: "--cascade", Description: "Only relevant during a force replace. If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController)."},
	{Text: "-f", Description: "Filename, directory, or URL to file to use to replace the resource."},
	{Text: "--filename", Description: "Filename, directory, or URL to file to use to replace the resource."},
	{Text: "--force", Description: "Delete and re-create the specified resource"},
	{Text: "--grace-period", Description: "Only relevant during a force replace. Period of time in seconds given to the old resource to terminate gracefully. Ignored if negative."},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "--schema-cache-dir", Description: "If non-empty, load/store cached API schemas in this directory, default is '$HOME/.kube/schema'"},
	{Text: "--timeout", Description: "Only relevant during a force replace. The length of time to wait before giving up on a delete of the old resource, zero means determine a timeout from the size of the object"},
	{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var flagPatch = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to update"},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to update"},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "-p", Description: "The patch to be applied to the resource JSON file."},
	{Text: "--patch", Description: "The patch to be applied to the resource JSON file."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--type", Description: "The type of patch being provided; one of [json merge strategic]"},
}

var flagDelete = []prompt.Suggest{
	{Text: "--all", Description: "to select all the specified resources."},
	{Text: "--cascade", Description: "If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController).  Default true."},
	{Text: "-f", Description: "Filename, directory, or URL to a file containing the resource to delete."},
	{Text: "--filename", Description: "Filename, directory, or URL to a file containing the resource to delete."},
	{Text: "--grace-period", Description: "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative."},
	{Text: "--ignore-not-found", Description: "Treat 'resource not found' as a successful delete. Defaults to 'true' when --all is specified."},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "-l", Description: "Selector (label query) to filter on."},
	{Text: "--selector", Description: "Selector (label query) to filter on."},
	{Text: "--timeout", Description: "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object"},
}

var flagEdit = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to file to use to edit the resource"},
	{Text: "--filename", Description: "Filename, directory, or URL to file to use to edit the resource"},
	{Text: "-o", Description: "Output format. One of: yaml|json."},
	{Text: "--output", Description: "Output format. One of: yaml|json."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "--windows-line-endings", Description: "Use Windows line-endings (default Unix line-endings)"},
}

var flagApply = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to file that contains the configuration to apply"},
	{Text: "--filename", Description: "Filename, directory, or URL to file that contains the configuration to apply"},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--schema-cache-dir", Description: "If non-empty, load/store cached API schemas in this directory, default is '$HOME/.kube/schema'"},
	{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var flagLogs = []prompt.Suggest{
	{Text: "-c", Description: "Print the logs of this container"},
	{Text: "--container", Description: "Print the logs of this container"},
	{Text: "-f", Description: "Specify if the logs should be streamed."},
	{Text: "--follow", Description: "Specify if the logs should be streamed."},
	{Text: "--limit-bytes", Description: "Maximum bytes of logs to return. Defaults to no limit."},
	{Text: "-p", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "--previous", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "--since", Description: "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--since-time", Description: "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--tail", Description: "Lines of recent log file to display. Defaults to -1, showing all log lines."},
	{Text: "--timestamps", Description: "Include timestamps on each line in the log output"},
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

var flagCordon = []prompt.Suggest{
	{Text: "--force", Description: "Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, or DaemonSet."},
	{Text: "--grace-period", Description: "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used."},
	{Text: "--ignore-daemonsets", Description: "Ignore DaemonSet-managed pods."},
}

var flagScale = []prompt.Suggest{
	{Text: "--current-replicas", Description: "Precondition for current size. Requires that the current size of the resource match this value in order to scale."},
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to set a new size"},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to set a new size"},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--replicas", Description: "The new desired number of replicas. Required."},
	{Text: "--resource-version", Description: "Precondition for resource version. Requires that the current resource version match this value in order to scale."},
	{Text: "--timeout", Description: "The length of time to wait before giving up on a scale operation, zero means don't wait."},
}

var flagAttach = []prompt.Suggest{
	{Text: "-c", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	{Text: "--container", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	{Text: "-i", Description: "Pass stdin to the container"},
	{Text: "--stdin", Description: "Pass stdin to the container"},
	{Text: "-t", Description: "Stdin is a TTY"},
	{Text: "--tty", Description: "Stdin is a TTY"},
}

var flagExec = []prompt.Suggest{
	{Text: "-c", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	{Text: "--container", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	{Text: "-p", Description: "Pod name"},
	{Text: "--pod", Description: "Pod name"},
	{Text: "-i", Description: "Pass stdin to the container"},
	{Text: "--stdin", Description: "Pass stdin to the container"},
	{Text: "-t", Description: "Stdin is a TTY"},
	{Text: "--tty", Description: "Stdin is a TTY"},
}

var flagPortForward = []prompt.Suggest{
	{Text: "-p", Description: "Pod name"},
	{Text: "--pod", Description: "Pod name"},
}

var flagProxy = []prompt.Suggest{
	{Text: "--accept-hosts", Description: "Regular expression for hosts that the proxy should accept."},
	{Text: "--accept-paths", Description: "Regular expression for paths that the proxy should accept. (default '^/.*')"},
	{Text: "--address", Description: "The IP address on which to serve on. (default '127.0.0.1')"},
	{Text: "--api-prefix", Description: "Prefix to serve the proxied API under. (default '/')"},
	{Text: "--disable-filter", Description: "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port."},
	{Text: "-p", Description: "The port on which to run the proxy. Set to 0 to pick a random port. (default: 8001)"},
	{Text: "--port", Description: "The port on which to run the proxy. Set to 0 to pick a random port. (default: 8001)"},
	{Text: "--reject-methods", Description: "Regular expression for HTTP methods that the proxy should reject. (default 'POST,PUT,PATCH')"},
	{Text: "--reject-paths", Description: "Regular expression for paths that the proxy should reject. (default '^/api/.*/exec,^/api/.*/run,^/api/.*/attach')"},
	{Text: "-u", Description: "Unix socket on which to run the proxy."},
	{Text: "--unix-socket", Description: "Unix socket on which to run the proxy."},
	{Text: "-w", Description: "Also serve static files from the given directory under the specified prefix."},
	{Text: "--www", Description: "Also serve static files from the given directory under the specified prefix."},
	{Text: "-P", Description: "Prefix to serve static files under, if static file directory is specified. (default '/static/')"},
	{Text: "--www-prefix", Description: "Prefix to serve static files under, if static file directory is specified. (default '/static/')"},
}

var flagRun = []prompt.Suggest{
	{Text: "--attach", Description: "If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--interactive' is set, in which case the default is true."},
	{Text: "--command", Description: "If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default."},
	{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	{Text: "--env", Description: "Environment variables to set in the container"},
	{Text: "--expose", Description: "If true, a public, external service is created for the container(s) which are run"},
	{Text: "--generator", Description: "The name of the API generator to use.  Default is 'deployment/v1beta1' if --restart=Always, otherwise the default is 'job/v1'.  This will happen only for cluster version at least 1.2, for olders we will fallback to 'run/v1' for --restart=Always, 'run-pod/v1' for others."},
	{Text: "--hostport", Description: "The host port mapping for the container port. To demonstrate a single-machine container."},
	{Text: "--image", Description: "The image for the container to run."},
	{Text: "-l", Description: "Labels to apply to the pod(s)."},
	{Text: "--labels", Description: "Labels to apply to the pod(s)."},
	{Text: "--leave-stdin-open ", Description: "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes."},
	{Text: "--limits", Description: "The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'"},
	{Text: "--no-headers", Description: "When using the default output, don't print headers."},
	{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--overrides", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	{Text: "--port", Description: "The port that this container exposes.  If --expose is true, this is also the port used by the service that is created."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "-r", Description: "Number of replicas to create for this container. Default is 1."},
	{Text: "--replicas", Description: "Number of replicas to create for this container. Default is 1."},
	{Text: "--requests", Description: "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'"},
	{Text: "--restart", Description: "The restart policy for this Pod.  Legal values [Always, OnFailure, Never].  If set to 'Always' a deployment is created for this pod, if set to OnFailure or Never, a job is created for this pod and --replicas must be 1.  Default 'Always'"},
	{Text: "--rm", Description: "If true, delete resources created in this command for attached containers."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "--service-generator", Description: "The name of the generator to use for creating a service.  Only used if --expose is true"},
	{Text: "--service-overrides", Description: "An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true."},
	{Text: "-a", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-all", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "-i", Description: "Keep stdin open on the container(s) in the pod, even if nothing is attached."},
	{Text: "--stdin", Description: "Keep stdin open on the container(s) in the pod, even if nothing is attached."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--tty", Description: "Allocated a TTY for each container in the pod.  Because -t is currently shorthand for --template, -t is not supported for --tty. This shorthand is deprecated and we expect to adopt -t for --tty soon."},
}

var flagExpose = []prompt.Suggest{
	{Text: "--container-port", Description: "IP to assign to to the Load Balancer. If empty, an ephemeral IP will be created and used (cloud-provider specific)."},
	{Text: "--dry-run", Description: "The name for the newly created object."},
	{Text: "--external-ip", Description: "When using the default output, don't print headers."},
	{Text: "-f", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--filename", Description: "Output format. One of: json|yaml|wide|name|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://releases.k8s.io/release-1.2/docs/user-guide/jsonpath.md]."},
	{Text: "--generator", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "-l", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	{Text: "--labels", Description: "The port that the service should serve on. Copied from the resource being exposed, if unspecified"},
	{Text: "--load-balancer-ip", Description: "The network protocol for the service to be created. Default is 'tcp'."},
	{Text: "--name", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--no-headers", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "-o", Description: "A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set."},
	{Text: "--output", Description: "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'"},
	{Text: "--output-version", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--overrides", Description: "When printing, show all resources (default hide terminated pods.)"},
	{Text: "--port", Description: "When printing, show all labels as the last column (default hide labels column)"},
	{Text: "--protocol", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	{Text: "--record", Description: "Name or number for the port on the container that the service should direct traffic to. Optional."},
	{Text: "--save-config", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	{Text: "--selector", Description: "Type for this service: ClusterIP, NodePort, or LoadBalancer. Default is 'ClusterIP'."},
	{Text: "--session-affinity", Description: "Synonym for --target-port"},
	{Text: "-a", Description: "If true, only print the object that would be sent, without creating it."},
	{Text: "--show-all", Description: "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP."},
	{Text: "--show-labels", Description: "Filename, directory, or URL to a file identifying the resource to expose a service"},
	{Text: "--sort-by", Description: "Filename, directory, or URL to a file identifying the resource to expose a service"},
	{Text: "--target-port", Description: "The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'."},
	{Text: "--template", Description: "Labels to apply to the service created by this call."},
	{Text: "--type", Description: "Labels to apply to the service created by this call."},
}

var flagClusterInfo = []prompt.Suggest{
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
}

var flagExplain = []prompt.Suggest{
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "--recursive", Description: "Print the fields of fields (Currently only 1 level deep)"},
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
