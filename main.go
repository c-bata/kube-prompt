package main

import (
	"fmt"

	"github.com/c-bata/go-prompt-toolkit"
	"github.com/c-bata/kube-prompt/kube"
)

var (
	version  string
	revision string
)

func main() {
	fmt.Printf("kube-prompt: powerful interactive kubernetes client. (version: '%s', revision: '%s')\n", version, revision)
	defer fmt.Println("Goodbye!")
	pt := prompt.NewPrompt(
		kube.Executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: powerful kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	pt.Run()
}
