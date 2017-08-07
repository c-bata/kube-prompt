package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/kube-prompt/kube"
)

var (
	version  string
	revision string
)

func main() {
	fmt.Printf("kube-prompt: interactive kubernetes client. (version: '%s', revision: '%s')\n", version, revision)
	defer fmt.Println("Goodbye!")
	p := prompt.New(
		kube.Executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
