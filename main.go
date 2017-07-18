package main

import (
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt-toolkit"
	"github.com/c-bata/kube-prompt/kube"
)

func executor(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	args := strings.Split(s, " ")
	cmd := exec.Command("kubectl", args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func main() {
	pt := prompt.NewPrompt(
		executor,
		kube.Completer,
		prompt.OptionTitle("kube-prompt: powerful kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	pt.Run()
}
