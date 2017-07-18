package kube

import (
	"fmt"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"time"
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
	podList *v1.PodList
	podLastFetchedAt time.Time
)

func fetchPods() {
	if time.Since(podLastFetchedAt) < thresholdFetchInterval {
		return
	}
	podList, _ = getClient().Pods(api.NamespaceDefault).List(v1.ListOptions{})
	return
}

func getPodNames() []string {
	go fetchPods()
	if podList == nil || len(podList.Items) == 0 {
		return []string{}
	}
	names := make([]string, len(podList.Items))
	for i := range podList.Items {
		names[i] = podList.Items[i].Name
	}
	return names
}

func formatPod(pod *v1.Pod) string {
	var res string
	res += "Status:\n"
	res += fmt.Sprintf("  Phase   = %s\n", pod.Status.Phase)
	res += fmt.Sprintf("  Message = %s\n", pod.Status.Message)
	res += fmt.Sprintf("  Start   = %s\n", pod.Status.StartTime.String())
	res += "\nLabels:\n"
	for k := range pod.Labels {
		res += fmt.Sprintf("  %s=%s\n", k, pod.Labels[k])
	}
	return res
}

func describePod(name string) string {
	pod, err := client.Pods(api.NamespaceDefault).Get(name)
	if err != nil {
		return err.Error()
	}
	return formatPod(pod)
}

/* Deployment */

var (
	deploymentList *v1beta1.DeploymentList
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

func formatDeployment(deployment *v1beta1.Deployment) string {
	var res string
	res += fmt.Sprintf("Name:\t%s\n", deployment.Name)
	res += fmt.Sprintf("NameSpace:\t%s\n", deployment.Namespace)
	res += fmt.Sprintf("CreationTimestamp:\t%s\n", deployment.CreationTimestamp.String())
	if len(deployment.Labels) == 0 {
		res += "Labels:\tnone\n"
	} else {
		res += "Labels:"
		for k := range deployment.Labels {
			res += fmt.Sprintf("\t%s=%s\n", k, deployment.Labels[k])
		}
	}
	res += fmt.Sprintf("Selector:\t%s\n", deployment.Spec.Selector)
	res += fmt.Sprintf(
		"Replicas:\t%d updated | %d total | %d avairable | %d unavairable\n",
		deployment.Status.UpdatedReplicas,
		deployment.Status.Replicas,
		deployment.Status.AvailableReplicas,
		deployment.Status.UnavailableReplicas,
	)
	res += fmt.Sprintf("StrategyType:\t%s\n", deployment.Spec.Strategy.Type)
	res += fmt.Sprintf("MinReadySeconds:\t%d\n", deployment.Spec.MinReadySeconds)
	res += "Conditions:\n"
	res += "  Type\tStatus\tReason\n"
	res += "  ----\t------\t------\n"
	conditions := deployment.Status.Conditions
	for i := range conditions {
		res += fmt.Sprintf("  %s\t%s\t%s\t", conditions[i].Type, conditions[i].Status, conditions[i].Reason)
	}
	return res
}

func describeDeployment(name string) string {
	d, err := client.Deployments(api.NamespaceDefault).Get(name)
	if err != nil {
		return err.Error()
	}
	return formatDeployment(d)
}
