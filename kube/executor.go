package kube

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bytes"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func ExecuteAndReturnList(s string) (list []string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	ss := string(out.Bytes())
	return strings.Split(ss, "\n")

}
