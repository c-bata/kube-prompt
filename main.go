package main

import (
	"github.com/c-bata/go-prompt-toolkit/prompt"
	"github.com/c-bata/kube-prompt/kube"
)

func executor(s string) string {
	return s
}

func main() {
	pt := prompt.NewPrompt(
		executor,
		kube.Completer,
	)
	pt.Run()
}
