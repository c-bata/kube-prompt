package kube

import (
	"github.com/c-bata/go-prompt"
	"k8s.io/client-go/kubernetes"
)

func loadApiResources(client *kubernetes.Clientset) error {
	resources, err := client.ServerResources()
	if err != nil {
		return err
	}

	result := make([]prompt.Suggest, 0, len(resources)*2)
	for _, resource := range resources {
		for _, apiResource := range resource.APIResources {
			result = append(result, prompt.Suggest{
				Text: apiResource.Name,
			})
			for _, name := range apiResource.ShortNames {
				result = append(result, prompt.Suggest{
					Text: name,
				})
			}
		}
	}

	resourceTypes = result
	subcommands = result
	return nil
}
