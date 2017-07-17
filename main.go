package main

import "github.com/c-bata/go-prompt-toolkit/prompt"

func executor(b *prompt.Buffer) string {
	s := b.Text()
	return s
}

func completer(b *prompt.Buffer) []string {
	return []string{
		"foo",
		"foo",
		"foo",
	}
}

func main() {
	pt := prompt.NewPrompt(
		executor,
		completer,
	)
	pt.Run()
}
