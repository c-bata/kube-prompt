package main

import (
	"github.com/c-bata/go-prompt-toolkit"
	"github.com/c-bata/kube-prompt/kube"
)

func main() {
	pt := prompt.NewPrompt(
		kube.Executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: powerful kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	pt.Run()
}
