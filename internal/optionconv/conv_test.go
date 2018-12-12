package optionconv_test

import (
	"fmt"
	"reflect"
	"testing"

	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/kube-prompt/internal/optionconv"
)

func ExampleGetOptionsFromHelpText() {
	input := `Expose a resource as a new Kubernetes service.

Looks up a deployment, service, replica set, replication controller or pod by name and uses the selector for that
resource as the selector for a new service on the specified port. A deployment or replica set will be exposed as a
service only if its selector is convertible to a selector that service supports, i.e. when the selector contains only
the matchLabels component. Note that if no port is specified via --port and the exposed resource has multiple ports, all
will be re-used by the new service. Also if no labels are specified, the new service will re-use the labels from the
resource it exposes.

Possible resources include (case insensitive):

pod (po), service (svc), replicationcontroller (rc), deployment (deploy), replicaset (rs)

Examples:
  # Create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000.
  kubectl expose rc nginx --port=80 --target-port=8000

  # Create a service for a replication controller identified by type and name specified in "nginx-controller.yaml",
which serves on port 80 and connects to the containers on port 8000.
  kubectl expose -f nginx-controller.yaml --port=80 --target-port=8000

  # Create a service for a pod valid-pod, which serves on port 444 with the name "frontend"
  kubectl expose pod valid-pod --port=444 --name=frontend

  # Create a second service based on the above service, exposing the container port 8443 as port 443 with the name
"nginx-https"
  kubectl expose service nginx --port=443 --target-port=8443 --name=nginx-https

Options:
      --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
the template. Only applies to golang and jsonpath output formats.
      --cluster-ip='': ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create
a headless service.
      --dry-run=false: If true, only print the object that would be sent, without sending it.
      --external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP
is routed to a node, the service can be accessed by this IP in addition to its generated service IP.
  -f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service
      --generator='service/v2': The name of the API generator to use. There are 2 generators: 'service/v1' and
'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed
in v2. Default is 'service/v2'.
  -l, --labels='': Labels to apply to the service created by this call.
      --load-balancer-ip='': IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used
(cloud-provider specific).
      --name='': The name for the newly created object.
      --no-headers=false: When using the default or custom-column output format, don't print headers (default print
headers).
  -o, --output='': Output format. One of:
json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...
See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template
[http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template
[http://kubernetes.io/docs/user-guide/jsonpath].
      --overrides='': An inline JSON override for the generated object. If this is non-empty, it is used to override the
generated object. Requires that the object supply a valid apiVersion field.
      --port='': The port that the service should serve on. Copied from the resource being exposed, if unspecified

Usage:
  kubectl expose (-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP] [--target-port=number-or-name]
[--name=name] [--external-ip=external-ip-of-service] [--type=type] [options]

Use "kubectl options" for a list of global command-line options (applies to all commands).`
	got, _ := optionconv.GetOptionsFromHelpText(input)
	fmt.Println(got)
	// Output:
	// --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
	// the template. Only applies to golang and jsonpath output formats.
	//       --cluster-ip='': ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create
	// a headless service.
	//       --dry-run=false: If true, only print the object that would be sent, without sending it.
	//       --external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP
	// is routed to a node, the service can be accessed by this IP in addition to its generated service IP.
	//   -f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service
	//       --generator='service/v2': The name of the API generator to use. There are 2 generators: 'service/v1' and
	// 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed
	// in v2. Default is 'service/v2'.
	//   -l, --labels='': Labels to apply to the service created by this call.
	//       --load-balancer-ip='': IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used
	// (cloud-provider specific).
	//       --name='': The name for the newly created object.
	//       --no-headers=false: When using the default or custom-column output format, don't print headers (default print
	// headers).
	//   -o, --output='': Output format. One of:
	// json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=...
	// See custom columns [http://kubernetes.io/docs/user-guide/kubectl-overview/#custom-columns], golang template
	// [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template
	// [http://kubernetes.io/docs/user-guide/jsonpath].
	//       --overrides='': An inline JSON override for the generated object. If this is non-empty, it is used to override the
	// generated object. Requires that the object supply a valid apiVersion field.
	//       --port='': The port that the service should serve on. Copied from the resource being exposed, if unspecified
}

func ExampleSplitOption() {
	in := `      --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in
the template. Only applies to golang and jsonpath output formats.
      --cluster-ip='': ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create
a headless service.
      --dry-run=false: If true, only print the object that would be sent, without sending it.
      --external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP
is routed to a node, the service can be accessed by this IP in addition to its generated service IP.
  -f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service
      --generator='service/v2': The name of the API generator to use. There are 2 generators: 'service/v1' and
'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed
in v2. Default is 'service/v2'.
      --port='': The port that the service should serve on. Copied from the resource being exposed, if unspecified`
	for _, o := range optionconv.SplitOptions(in) {
		fmt.Println(o)
	}
	// Output:
	// --allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.
	// --cluster-ip='': ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service.
	// --dry-run=false: If true, only print the object that would be sent, without sending it.
	// --external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.
	// -f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service
	// --generator='service/v2': The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'.
	// --port='': The port that the service should serve on. Copied from the resource being exposed, if unspecified
}

func TestConvertToSuggestions(t *testing.T) {
	input := []string{
		"--allow-missing-template-keys=true: If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.",
		"--external-ip='': Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.",
		"-f, --filename=[]: Filename, directory, or URL to files identifying the resource to expose a service",
		"--generator='service/v2': The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'.",
		"--port='': The port that the service should serve on. Copied from the resource being exposed, if unspecified",
	}
	actual := optionconv.ConvertToSuggestions(input)
	expected := []prompt.Suggest{
		{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
		{Text: "--external-ip", Description: "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP."},
		{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
		{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
		{Text: "--generator", Description: "The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'."},
		{Text: "--port", Description: "The port that the service should serve on. Copied from the resource being exposed, if unspecified"},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected:\n%#v\n\ngot:\n%#v\n", expected, actual)
	}
}
