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
	fmt.Printf("kube-prompt %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
	defer fmt.Println("Bye!")
	p := prompt.New(
		kube.Executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
