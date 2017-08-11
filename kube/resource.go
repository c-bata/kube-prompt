package kube

import (
	"sync/atomic"
	"time"

	"github.com/c-bata/go-prompt"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

const thresholdFetchInterval = 10 * time.Second

var resourceTypes = []prompt.Suggest{
	//{Text: "clusters"},  // valid only for federation apiservers
	{Text: "componentstatuses"},
	{Text: "configmaps"},
	{Text: "daemonsets"},
	{Text: "deployments"},
	{Text: "endpoints"},
	{Text: "events"},
	{Text: "horizontalpodautoscalers"},
	{Text: "ingresses"},
	{Text: "jobs"},
	{Text: "limitranges"},
	{Text: "namespaces"},
	{Text: "networkpolicies"},
	{Text: "nodes"},
	{Text: "persistentvolumeclaims"},
	{Text: "persistentvolumes"},
	{Text: "pods"},
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

/* Component Status */

var (
	componentStatusList          atomic.Value
	componentStatusLastFetchedAt time.Time
)

func fetchComponentStatusList() {
	if time.Since(componentStatusLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ComponentStatuses().List(v1.ListOptions{})
	componentStatusList.Store(l)
	return
}

func getComponentStatusCompletions() []prompt.Suggest {
	go fetchComponentStatusList()
	l, ok := componentStatusList.Load().(*v1.ComponentStatusList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Config Maps */

var (
	configMapsList          atomic.Value
	configMapsLastFetchedAt time.Time
)

func fetchConfigMapList() {
	if time.Since(configMapsLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ConfigMaps(api.NamespaceDefault).List(v1.ListOptions{})
	configMapsList.Store(l)
	return
}

func getConfigMapSuggestions() []prompt.Suggest {
	go fetchConfigMapList()
	l, ok := configMapsList.Load().(*v1.ConfigMapList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Pod */

var (
	podList          atomic.Value
	podLastFetchedAt time.Time
)

func fetchPods() {
	if time.Since(podLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Pods(api.NamespaceDefault).List(v1.ListOptions{})
	podList.Store(l)
	return
}

func getPodSuggestions() []prompt.Suggest {
	go fetchPods()
	l, ok := podList.Load().(*v1.PodList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text:        l.Items[i].Name,
			Description: string(l.Items[i].Status.Phase),
		}
	}
	return s
}

/* Daemon Sets */

var (
	daemonSetList          atomic.Value
	daemonSetLastFetchedAt time.Time
)

func fetchDaemonSetList() {
	if time.Since(daemonSetLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().DaemonSets(api.NamespaceDefault).List(v1.ListOptions{})
	daemonSetList.Store(l)
	return
}

func getDaemonSetSuggestions() []prompt.Suggest {
	go fetchDaemonSetList()
	l, ok := daemonSetList.Load().(*v1beta1.DaemonSetList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Deployment */

var (
	deploymentList          atomic.Value
	deploymentLastFetchedAt time.Time
)

func fetchDeployments() {
	if time.Since(deploymentLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Deployments(api.NamespaceDefault).List(v1.ListOptions{})
	deploymentList.Store(l)
	return
}

func getDeploymentNames() []prompt.Suggest {
	go fetchDeployments()
	l, ok := deploymentList.Load().(*v1beta1.DeploymentList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Endpoint */

var (
	endpointList          atomic.Value
	endpointLastFetchedAt time.Time
)

func fetchEndpoints() {
	if time.Since(endpointLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Endpoints(api.NamespaceDefault).List(v1.ListOptions{})
	endpointList.Store(l)
	return
}

func getEndpointsSuggestions() []prompt.Suggest {
	go fetchEndpoints()
	l, ok := endpointList.Load().(*v1.EndpointsList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Events */

var (
	eventList          atomic.Value
	eventLastFetchedAt time.Time
)

func fetchEvents() {
	if time.Since(eventLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Events(api.NamespaceDefault).List(v1.ListOptions{})
	eventList.Store(l)
	return
}

func getEventsSuggestions() []prompt.Suggest {
	go fetchEvents()
	l, ok := eventList.Load().(*v1.EventList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Node */

var (
	nodeList          atomic.Value
	nodeLastFetchedAt time.Time
)

func fetchNodeList() {
	if time.Since(nodeLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Nodes().List(v1.ListOptions{})
	nodeList.Store(l)
	return
}

func getNodeSuggestions() []prompt.Suggest {
	go fetchNodeList()
	l, ok := nodeList.Load().(*v1.NodeList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Secret */

var (
	secretList          atomic.Value
	secretLastFetchedAt time.Time
)

func fetchSecretList() {
	if time.Since(secretLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Secrets(api.NamespaceDefault).List(v1.ListOptions{})
	secretList.Store(l)
	return
}

func getSecretSuggestions() []prompt.Suggest {
	go fetchSecretList()
	l, ok := secretList.Load().(*v1.SecretList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Service Account */

var (
	serviceAccountList          atomic.Value
	serviceAccountLastFetchedAt time.Time
)

func fetchServiceAccountList() {
	if time.Since(serviceAccountLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ServiceAccounts(api.NamespaceDefault).List(v1.ListOptions{})
	serviceAccountList.Store(l)
	return
}

func getServiceAccountSuggestions() []prompt.Suggest {
	go fetchServiceAccountList()
	l, ok := serviceAccountList.Load().(*v1.ServiceAccountList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}

/* Service */

var (
	serviceList          atomic.Value
	serviceLastFetchedAt time.Time
)

func fetchServiceList() {
	if time.Since(serviceLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Services(api.NamespaceDefault).List(v1.ListOptions{})
	serviceList.Store(l)
	return
}

func getServiceSuggestions() []prompt.Suggest {
	go fetchServiceList()
	l, ok := serviceList.Load().(*v1.ServiceAccountList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l.Items))
	for i := range l.Items {
		s[i] = prompt.Suggest{
			Text: l.Items[i].Name,
		}
	}
	return s
}
