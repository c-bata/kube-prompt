package kube

import (
	"os/exec"
	"strings"
)

var unSupportedCommand = []string{
	"exec",
	"attach",
	"port-forward",
}

func Executor(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	args := strings.Split(s, " ")
	for _, c := range unSupportedCommand {
		if c == args[0] {
			return "Sorry, this command is still not supported... :-("
		}
	}

	cmd := exec.Command("kubectl", args...)
	out, _ := cmd.CombinedOutput()
	return string(out)
}
