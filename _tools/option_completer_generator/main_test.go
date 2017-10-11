package main

import (
	"testing"

	"github.com/c-bata/go-prompt"
	"github.com/stretchr/testify/assert"
)

var exampleGetCommandHelpText = `
Display one or many resources.

Possible resource types include (case insensitive): pods (po), services (svc), deployments,
replicasets (rs), replicationcontrollers (rc), nodes (no), events (ev), limitranges (limits),
persistentvolumes (pv), persistentvolumeclaims (pvc), resourcequotas (quota), namespaces (ns),
serviceaccounts, ingresses (ing), horizontalpodautoscalers (hpa), daemonsets (ds), configmaps,
componentstatuses (cs), endpoints (ep), and secrets.

By specifying the output as 'template' and providing a Go template as the value
of the --template flag, you can filter the attributes of the fetched resource(s).

Usage:
  kubectl get [(-o|--output=)json|yaml|wide|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...] (TYPE [NAME | -l label] | TYPE/NAME ...) [flags] [flags]

Examples:
# List all pods in ps output format.
kubectl get pods

Flags:
      --all-namespaces[=false]: If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -w, --watch[=false]: After listing/getting the requested object, watch for changes.
      --watch-only[=false]: Watch for changes to the requested object(s), without listing/getting first.

Global Flags:
      --alsologtostderr[=false]: log to standard error as well as files
  -s, --server="": The address and port of the Kubernetes API server
      --vmodule=: comma-separated list of pattern=N settings for file-filtered logging
`

func TestGetFlags(t *testing.T) {
	actualFlags, actualGlobalFlags := getFlags(exampleGetCommandHelpText)
	expectedFlags := `      --all-namespaces[=false]: If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -w, --watch[=false]: After listing/getting the requested object, watch for changes.
      --watch-only[=false]: Watch for changes to the requested object(s), without listing/getting first.`
	assert.Equal(t, actualFlags, expectedFlags)

	expectedGlobalFlags := `      --alsologtostderr[=false]: log to standard error as well as files
  -s, --server="": The address and port of the Kubernetes API server
      --vmodule=: comma-separated list of pattern=N settings for file-filtered logging`
	assert.Equal(t, actualGlobalFlags, expectedGlobalFlags)
}

func TestConvertToSuggest(t *testing.T) {
	var scenarioTable = []struct {
		input    string
		expected []prompt.Suggest
	}{
		{
			input: `      --all-namespaces[=false]: If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`,
			expected: []prompt.Suggest{
				{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
			},
		},
		{
			input: `  -w, --watch[=false]: After listing/getting the requested object, watch for changes.`,
			expected: []prompt.Suggest{
				{Text: "-w", Description: "After listing/getting the requested object, watch for changes."},
				{Text: "--watch", Description: "After listing/getting the requested object, watch for changes."},
			},
		},
		{
			input: `      --vmodule=: comma-separated list of pattern=N settings for file-filtered logging`,
			expected: []prompt.Suggest{
				{Text: "--vmodule", Description: "comma-separated list of pattern=N settings for file-filtered logging"},
			},
		},
	}

	for _, s := range scenarioTable {
		actual := convertToSuggest(s.input)
		assert.Equal(t, s.expected, actual)
	}

}

func TestConvertToSuggestions(t *testing.T) {
	input := `      --all-namespaces[=false]: If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.
  -w, --watch[=false]: After listing/getting the requested object, watch for changes.
      --watch-only[=false]: Watch for changes to the requested object(s), without listing/getting first.`
	actual := convertToSuggestions(input)
	expected := []prompt.Suggest{
		{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
		{Text: "-w", Description: "After listing/getting the requested object, watch for changes."},
		{Text: "--watch", Description: "After listing/getting the requested object, watch for changes."},
		{Text: "--watch-only", Description: "Watch for changes to the requested object(s), without listing/getting first."},
	}
	assert.Equal(t, expected, actual)
}
