package kube

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/c-bata/kube-prompt/internal/debug"
	"github.com/efekarakus/termcolor"

	"github.com/google/shlex"
)

func getOutputLexerType(kubectlCmd string) (lexerType string) {
	kubectlArgs, _ := shlex.Split(kubectlCmd)

	for index, arg := range kubectlArgs {
		if arg == "-o" {
			lexerType = strings.ToLower(kubectlArgs[index+1])
		} else if strings.HasPrefix(arg, "--output=") {
			lexerType = strings.ToLower(strings.Split(arg, "=")[1])
		}
	}
	return
}

func getOutputFormatter() (formatter string) {
	switch termColor := termcolor.SupportLevel(os.Stdout); termColor {
	case termcolor.Level16M:
		formatter = "terminal16m"
	case termcolor.Level256:
		formatter = "terminal256"
	case termcolor.LevelBasic:
		formatter = "terminal"
	case termcolor.LevelNone:
	default:
		formatter = ""
	}
	return
}

// Executor passed to go-prompt
func Executor(style string) func(string) {
	return func(s string) {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		} else if s == "quit" || s == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
			return
		}

		output := &bytes.Buffer{}
		cmd := exec.Command("/bin/sh", "-c", "kubectl "+s)
		cmd.Stdin = os.Stdin
		cmd.Stdout = output
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		} else {
			highlighted := &bytes.Buffer{}
			lexerType := getOutputLexerType(s)
			formatterType := getOutputFormatter()

			if formatterType != "" && lexerType != "" && (lexerType == "json" || lexerType == "yaml") {
				quick.Highlight(highlighted, output.String(), lexerType, formatterType, style)
				os.Stdout.Write(highlighted.Bytes())
			} else {
				os.Stdout.Write(output.Bytes())
			}
		}
		return
	}
}

// ExecuteAndGetResult Execute the command provided by the argument and return the stdout
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
