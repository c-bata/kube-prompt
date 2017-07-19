package kube

import (
	"github.com/c-bata/go-prompt-toolkit"
	"strings"
)

func optionCompleter(args []string, long bool) []prompt.Completion {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	switch args[0] {
	case "get":
		if long {
			return prompt.FilterContains(
				prompt.FilterHasPrefix(optionGet, "--", false),
				strings.TrimLeft(args[l-1], "--"),
				true,
			)
		}
		return prompt.FilterContains(optionGet, strings.TrimLeft(args[l-1], "-"), true)
	}
	return []prompt.Completion{}
}

var optionHelp = []prompt.Completion{
	{Text: "-h"},
	{Text: "--help"},
}

var optionGet = []prompt.Completion{
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

var optionDescribe = []prompt.Completion{
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
