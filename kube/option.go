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
		suggests = optionGet
	case "describe":
		suggests = optionDescribe
	case "create":
		suggests = optionCreate
	case "replace":
		suggests = optionReplace
	case "patch":
		suggests = optionPatch
	case "delete":
		suggests = optionDelete
	case "edit":
		suggests = optionEdit
	case "cluster-info":
		suggests = optionClusterInfo
	case "explain":
		suggests = optionExplain
	case "cordon":
		suggests = optionCordon
	case "drain":
		suggests = optionDrain
	case "uncordon":
		suggests = optionHelp
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "view":
				suggests = optionConfigView
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

var optionGet = []prompt.Suggest{
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

var optionDescribe = []prompt.Suggest{
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

var optionCreate = []prompt.Suggest{
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

var optionReplace = []prompt.Suggest{
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

var optionPatch = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to a file identifying the resource to update"},
	{Text: "--filename", Description: "Filename, directory, or URL to a file identifying the resource to update"},
	{Text: "-o", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "--output", Description: "Output mode. Use '-o name' for shorter output (resource/name)."},
	{Text: "-p", Description: "The patch to be applied to the resource JSON file."},
	{Text: "--patch", Description: "The patch to be applied to the resource JSON file."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--type", Description: "The type of patch being provided; one of [json merge strategic]"},
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

var optionDelete = []prompt.Suggest{
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

var optionEdit = []prompt.Suggest{
	{Text: "-f", Description: "Filename, directory, or URL to file to use to edit the resource"},
	{Text: "--filename", Description: "Filename, directory, or URL to file to use to edit the resource"},
	{Text: "-o", Description: "Output format. One of: yaml|json."},
	{Text: "--output", Description: "Output format. One of: yaml|json."},
	{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1')."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation."},
	{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future."},
	{Text: "--windows-line-endings", Description: "Use Windows line-endings (default Unix line-endings)"},
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

var optionCordon = []prompt.Suggest{
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
	{Text: "--server", Description: "The address and port of the Kubernetes API server"},
	{Text: "--stderrthreshold", Description: "logs at or above this threshold go to stderr"},
	{Text: "--token", Description: "Bearer token for authentication to the API server."},
	{Text: "--user", Description: "The name of the kubeconfig user to use"},
	{Text: "--username", Description: "Username for basic authentication to the API server."},
	{Text: "--v", Description: "log level for V logs"},
	{Text: "--vmodule", Description: "comma-separated list of pattern=N settings for file-filtered logging"},
}

var optionDrain = []prompt.Suggest{
	{Text: "--force", Description: "Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, or DaemonSet."},
	{Text: "--grace-period", Description: "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used."},
	{Text: "--ignore-daemonsets", Description: "Ignore DaemonSet-managed pods."},
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
	{Text: "--server", Description: "The address and port of the Kubernetes API server"},
	{Text: "--stderrthreshold", Description: "logs at or above this threshold go to stderr"},
	{Text: "--token", Description: "Bearer token for authentication to the API server."},
	{Text: "--user", Description: "The name of the kubeconfig user to use"},
	{Text: "--username", Description: "Username for basic authentication to the API server."},
	{Text: "--v", Description: "log level for V logs"},
	{Text: "--vmodule", Description: "comma-separated list of pattern=N settings for file-filtered logging"},
}

var optionClusterInfo = []prompt.Suggest{
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
}

var optionExplain = []prompt.Suggest{
	{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	{Text: "--recursive", Description: "Print the fields of fields (Currently only 1 level deep)"},
}

var optionConfigView = []prompt.Suggest{
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
