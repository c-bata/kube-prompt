package kube

import (
	"strings"
	"os/exec"
)

var unSupportedCommand = []string{
	"exec",
}

func Executor(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	args := strings.Split(s, " ")
	for _, c := range unSupportedCommand {
		if c == args[0] {
			return "Sorry, this command still not supported."
		}
	}

	cmd := exec.Command("kubectl", args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
