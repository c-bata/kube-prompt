package kube

import (
	"fmt"
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/v1"
	"time"
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
	podList []v1.Pod
	podLastFetchedAt time.Time
)

func fetchPods() {
	if time.Since(podLastFetchedAt) < thresholdFetchInterval {
		return
	}
	client := getClient()
	pod, err := client.Pods(api.NamespaceDefault).List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	podList = pod.Items
	return
}

func getPodNames() []string {
	go fetchPods()
	names := make([]string, len(podList))
	for i := range podList {
		names[i] = podList[i].Name
	}
	return names
}

func describePod(podname string) string {
	pod, err := client.Pods(api.NamespaceDefault).Get(podname)
	if err != nil {
		panic(err)
	}

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
