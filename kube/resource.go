package kube

import (
	"time"

	"github.com/c-bata/go-prompt-toolkit"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

const thresholdFetchInterval = 10 * time.Second

var resourceTypes = []string{
	"clusters",
	"componentstatuses",
	"cs",
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

/* Pod */

var (
	podList          *v1.PodList
	podLastFetchedAt time.Time
)

func fetchPods() {
	if time.Since(podLastFetchedAt) < thresholdFetchInterval {
		return
	}
	podList, _ = getClient().Pods(api.NamespaceDefault).List(v1.ListOptions{})
	return
}

func getPodCompletions() []prompt.Completion {
	go fetchPods()
	if podList == nil || len(podList.Items) == 0 {
		return []prompt.Completion{}
	}
	completions := make([]prompt.Completion, len(podList.Items))
	for i := range podList.Items {
		completions[i] = prompt.Completion{
			Text:        podList.Items[i].Name,
			Description: string(podList.Items[i].Status.Phase),
		}
	}
	return completions
}

/* Deployment */

var (
	deploymentList          *v1beta1.DeploymentList
	deploymentLastFetchedAt time.Time
)

func fetchDeployments() {
	if time.Since(deploymentLastFetchedAt) < thresholdFetchInterval {
		return
	}
	deploymentList, _ = getClient().Deployments(api.NamespaceDefault).List(v1.ListOptions{})
}

func getDeploymentNames() []string {
	go fetchDeployments()
	if deploymentList == nil || len(deploymentList.Items) == 0 {
		return []string{}
	}
	names := make([]string, len(deploymentList.Items))
	for i := range deploymentList.Items {
		names[i] = deploymentList.Items[i].Name
	}
	return names
}
