package main

import (
	"flag"

	"github.com/c-bata/go-prompt-toolkit/prompt"
	"github.com/c-bata/kube-prompt/kube"
	"k8s.io/client-go/tools/clientcmd"
)

func executor(s string) string {
	return s
}

func main() {
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
	if *kubeconfig == "" {
		panic("-kubeconfig not specified")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	if err := kube.NewClient(config); err != nil {
		panic(err)
	}
	pt := prompt.NewPrompt(
		executor,
		kube.Completer,
	)
	pt.Run()
}
