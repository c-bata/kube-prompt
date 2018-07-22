package kube

import (
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // register gcp auth helper
	"k8s.io/client-go/tools/clientcmd"
)

var client *kubernetes.Clientset

func getClient() *kubernetes.Clientset {
	if client != nil {
		return client
	}
	client = NewClient()
	return client
}

func NewClient() *kubernetes.Clientset {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	config, err := loader.ClientConfig()
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return client
}
