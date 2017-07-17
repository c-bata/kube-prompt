package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var client *kubernetes.Clientset

func NewClient(config *rest.Config) (err error) {
	client, err = kubernetes.NewForConfig(config)
	return err
}
