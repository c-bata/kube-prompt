package kube

import (
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

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
