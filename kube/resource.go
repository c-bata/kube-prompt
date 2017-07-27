package kube

import (
	"sync/atomic"
	"time"

	"github.com/c-bata/go-prompt-toolkit"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

const thresholdFetchInterval = 10 * time.Second

var resourceTypes = []string{
	"clusters",
	"componentstatuses",        // aka 'cs'
	"configmaps",               // aka 'cm'
	"daemonsets",               // aka 'ds'
	"deployments",              // aka 'deploy'
	"endpoints",                // aka 'ep'
	"events",                   // aka 'ev'
	"horizontalpodautoscalers", // aka 'hpa'
	"ingresses",                // aka 'ing'
	"jobs",
	"limitranges", // aka 'limits'
	"namespaces",  // aka 'ns'
	"networkpolicies",
	"nodes",                  // aka 'no'
	"persistentvolumeclaims", // aka 'pvc'
	"persistentvolumes",      // aka 'pv'
	"pods",                   // aka 'po'
	"podsecuritypolicies",    // aka 'psp'
	"podtemplates",
	"replicasets",            // aka 'rs'
	"replicationcontrollers", // aka 'rc'
	"resourcequotas",         // aka 'quota'
	"secrets",
	"serviceaccounts", // aka 'sa'
	"services",        // aka 'svc'
	"statefulsets",
	"storageclasses",
	"thirdpartyresources",
	// shorten aliases
	"cs",
	"cm",
	"ds",
	"deploy",
	"ep",
	"ev",
	"hpa",
	"ing",
	"limits",
	"ns",
	"no",
	"pvc",
	"pv",
	"po",
	"psp",
	"rs",
	"rc",
	"quota",
	"sa",
	"svc",
}

/* Config Maps */

var (
	configMapsList       atomic.Value
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

func getConfigMapCompletions() []prompt.Completion {
	go fetchConfigMapList()
	l, ok := configMapsList.Load().(*v1.ConfigMapList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text: l.Items[i].Name,
		}
	}
	return completions
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

func getPodCompletions() []prompt.Completion {
	go fetchPods()
	l, ok := podList.Load().(*v1.PodList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text:        l.Items[i].Name,
			Description: string(l.Items[i].Status.Phase),
		}
	}
	return completions
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

func getDeploymentNames() []prompt.Completion {
	go fetchDeployments()
	l, ok := podList.Load().(*v1beta1.DeploymentList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text: l.Items[i].Name,
		}
	}
	return completions
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

func getNodeCompletions() []prompt.Completion {
	go fetchNodeList()
	l, ok := podList.Load().(*v1.NodeList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text: l.Items[i].Name,
		}
	}
	return completions
}

/* Secret */

var (
	secretList       atomic.Value
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

func getSecretCompletions() []prompt.Completion {
	go fetchSecretList()
	l, ok := secretList.Load().(*v1.SecretList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text: l.Items[i].Name,
		}
	}
	return completions
}

/* Service Account */

var (
	serviceAccountList       atomic.Value
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

func getServiceAccountCompletions() []prompt.Completion {
	go fetchServiceAccountList()
	l, ok := serviceAccountList.Load().(*v1.ServiceAccountList)
	if !ok || len(l.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(l.Items))
	for i := range l.Items {
		completions[i] = prompt.Completion{
			Text: l.Items[i].Name,
		}
	}
	return completions
}
