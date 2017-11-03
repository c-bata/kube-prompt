package kube

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync/atomic"
	"time"

	"github.com/c-bata/go-prompt"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

const thresholdFetchInterval = 10 * time.Second

var resourceTypes = []prompt.Suggest{
	{Text: "clusters"}, // valid only for federation apiservers
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
	{Text: "pod"},
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
	l, _ := getClient().ConfigMaps(api.NamespaceAll).List(v1.ListOptions{})
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

/* Contexts */

var (
	contextList        atomic.Value
	contextLastFetchAt time.Time
)

func fetchContextList() {
	if time.Since(contextLastFetchAt) < thresholdFetchInterval {
		return
	}
	r, err := ExecuteAndGetResult("config get-contexts --no-headers -o name")
	if err != nil {
		log.Printf("[WARN] Got Error when fetchContextList: %s", err.Error())
	}
	contextList.Store(strings.Split(r, "\n"))
}

func getContextSuggestions() []prompt.Suggest {
	go fetchContextList()
	l, ok := contextList.Load().([]string)
	if !ok || len(l) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(l))
	for i := range l {
		s[i] = prompt.Suggest{
			Text: l[i],
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
	l, _ := getClient().Pods(api.NamespaceAll).List(v1.ListOptions{})
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

func getPod(podName string) (v1.Pod, bool) {
	l, ok := podList.Load().(*v1.PodList)
	if !ok || len(l.Items) == 0 {
		return v1.Pod{}, false
	}
	for i := range l.Items {
		if podName == l.Items[i].Name {
			return l.Items[i], true
		}
	}
	return v1.Pod{}, false
}

func getPortsFromPodName(podName string) []prompt.Suggest {
	pod, found := getPod(podName)
	if !found {
		return []prompt.Suggest{}
	}

	// Extract unique ports
	portSet := make(map[int32]struct{})
	for i := range pod.Spec.Containers {
		ports := pod.Spec.Containers[i].Ports
		for j := range ports {
			portSet[ports[j].ContainerPort] = struct{}{}
		}
	}

	// Sort
	var ports []int
	for k := range portSet {
		ports = append(ports, int(k))
	}
	sort.Ints(ports)

	// Prepare suggestions
	suggests := make([]prompt.Suggest, 0, len(ports))
	for i := range ports {
		suggests = append(suggests, prompt.Suggest{
			Text: fmt.Sprintf("%d:%d", ports[i], ports[i]),
		})
	}
	return suggests
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
	l, _ := getClient().DaemonSets(api.NamespaceAll).List(v1.ListOptions{})
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
	l, _ := getClient().Deployments(api.NamespaceAll).List(v1.ListOptions{})
	deploymentList.Store(l)
	return
}

func getDeploymentSuggestions() []prompt.Suggest {
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
	l, _ := getClient().Endpoints(api.NamespaceAll).List(v1.ListOptions{})
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
	l, _ := getClient().Events(api.NamespaceAll).List(v1.ListOptions{})
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
	l, _ := getClient().Secrets(api.NamespaceAll).List(v1.ListOptions{})
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

/* Ingress */

var (
	ingressList          atomic.Value
	ingressLastFetchedAt time.Time
)

func fetchIngressList() {
	if time.Since(ingressLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Ingresses(api.NamespaceAll).List(v1.ListOptions{})
	ingressList.Store(l)
	return
}

func getIngressSuggestions() []prompt.Suggest {
	go fetchIngressList()
	l, ok := ingressList.Load().(*v1.NamespaceList)
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

/* LimitRange */

var (
	limitRangeList          atomic.Value
	limitRangeLastFetchedAt time.Time
)

func fetchLimitRangeList() {
	if time.Since(limitRangeLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().LimitRanges(api.NamespaceAll).List(v1.ListOptions{})
	limitRangeList.Store(l)
	return
}

func getLimitRangeSuggestions() []prompt.Suggest {
	go fetchLimitRangeList()
	l, ok := limitRangeList.Load().(*v1.NamespaceList)
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

/* NameSpaces */

var (
	namespaceList          atomic.Value
	namespaceLastFetchedAt time.Time
)

func fetchNameSpaceList() {
	if time.Since(namespaceLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().Namespaces().List(v1.ListOptions{})
	namespaceList.Store(l)
	return
}

func getNameSpaceSuggestions() []prompt.Suggest {
	go fetchNameSpaceList()
	l, ok := namespaceList.Load().(*v1.NamespaceList)
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

/* Persistent Volume Claims */

var (
	persistentVolumeClaimsList          atomic.Value
	persistentVolumeClaimsLastFetchedAt time.Time
)

func fetchPersistentVolumeClaimsList() {
	if time.Since(persistentVolumeClaimsLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().PersistentVolumeClaims(api.NamespaceAll).List(v1.ListOptions{})
	persistentVolumeClaimsList.Store(l)
	return
}

func getPersistentVolumeClaimSuggestions() []prompt.Suggest {
	go fetchPersistentVolumeClaimsList()
	l, ok := persistentVolumeClaimsList.Load().(*v1.PersistentVolumeClaimList)
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

/* Persistent Volumes */

var (
	persistentVolumesList          atomic.Value
	persistentVolumesLastFetchedAt time.Time
)

func fetchPersistentVolumeList() {
	if time.Since(persistentVolumesLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().PersistentVolumes().List(v1.ListOptions{})
	persistentVolumesList.Store(l)
	return
}

func getPersistentVolumeSuggestions() []prompt.Suggest {
	go fetchPersistentVolumeList()
	l, ok := persistentVolumesList.Load().(*v1.PersistentVolumeList)
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

/* Pod Security Policies */

var (
	podSecurityPolicyList          atomic.Value
	podSecurityPolicyLastFetchedAt time.Time
)

func fetchPodSecurityPolicyList() {
	if time.Since(podSecurityPolicyLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().PodSecurityPolicies().List(v1.ListOptions{})
	podSecurityPolicyList.Store(l)
	return
}

func getPodSecurityPolicySuggestions() []prompt.Suggest {
	go fetchPodSecurityPolicyList()
	l, ok := podSecurityPolicyList.Load().(*v1beta1.PodSecurityPolicyList)
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

/* Pod Templates */

var (
	podTemplateList          atomic.Value
	podTemplateLastFetchedAt time.Time
)

func fetchPodTemplateList() {
	if time.Since(podTemplateLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().PodTemplates(api.NamespaceAll).List(v1.ListOptions{})
	podTemplateList.Store(l)
	return
}

func getPodTemplateSuggestions() []prompt.Suggest {
	go fetchPodTemplateList()
	l, ok := podTemplateList.Load().(*v1.PodTemplateList)
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

/* Replica Sets */

var (
	replicaSetList          atomic.Value
	replicaSetLastFetchedAt time.Time
)

func fetchReplicaSetList() {
	if time.Since(replicaSetLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ReplicaSets(api.NamespaceAll).List(v1.ListOptions{})
	replicaSetList.Store(l)
	return
}

func getReplicaSetSuggestions() []prompt.Suggest {
	go fetchReplicaSetList()
	l, ok := replicaSetList.Load().(*v1beta1.ReplicaSetList)
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

/* Replication Controller */

var (
	replicationControllerList          atomic.Value
	replicationControllerLastFetchedAt time.Time
)

func fetchReplicationControllerList() {
	if time.Since(replicationControllerLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ReplicationControllers(api.NamespaceAll).List(v1.ListOptions{})
	replicationControllerList.Store(l)
	return
}

func getReplicationControllerSuggestions() []prompt.Suggest {
	go fetchReplicationControllerList()
	l, ok := replicationControllerList.Load().(*v1.ReplicationControllerList)
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

/* Resource quotas */

var (
	resourceQuotaList          atomic.Value
	resourceQuotaLastFetchedAt time.Time
)

func fetchResourceQuotaList() {
	if time.Since(resourceQuotaLastFetchedAt) < thresholdFetchInterval {
		return
	}
	l, _ := getClient().ResourceQuotas(api.NamespaceAll).List(v1.ListOptions{})
	resourceQuotaList.Store(l)
	return
}

func getResourceQuotasSuggestions() []prompt.Suggest {
	go fetchResourceQuotaList()
	l, ok := resourceQuotaList.Load().(*v1.ResourceQuotaList)
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
	l, _ := getClient().ServiceAccounts(api.NamespaceAll).List(v1.ListOptions{})
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
	l, _ := getClient().Services(api.NamespaceAll).List(v1.ListOptions{})
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
