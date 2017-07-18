package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt-toolkit"
	"github.com/c-bata/kube-prompt/kube"
)

var (
	version  string
	revision string
)

func executor(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	args := strings.Split(s, " ")
	switch args[0] {
	case "version":
		res := fmt.Sprintf("Kube-Prompt Version: {Version: \"%s\", GitCommit: \"%s\"}\n", version, revision)
		cmd := exec.Command("kubectl", "version")
		out, err := cmd.Output()
		if err != nil {
			return err.Error()
		}
		return res + string(out)
	default:
	}
	cmd := exec.Command("kubectl", args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func main() {
	fmt.Println("kube-prompt: powerful interactive kubernetes client.")
	defer fmt.Println("Goodbye!")
	pt := prompt.NewPrompt(
		executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: powerful kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	pt.Run()
}
