package kube

import (
	"strings"
)

func Executor(s string) string {
	args := strings.Split(s, " ")
	if len(args) == 3 {
		return thirdArgsExecutor(args[0], args[1], args[2])
	}
	return s
}

func thirdArgsExecutor(first, second, third string) string {
	if first == "describe" && second == "pods" {
		return describePod(third)
	}
	return first + " " + second + " " + third
}
