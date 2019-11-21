package kube

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/kube-prompt/internal/debug"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func GetKubeContext() (string) {
	out, err := exec.Command("/bin/sh", "-c", "kubectl config current-context").Output()
	if err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return strings.TrimRight(string(out), "\n\r") + ">>> "
}

func ChangeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	if strings.Contains(s, "use-context") {
		LivePrefixState.LivePrefix = GetKubeContext()
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
