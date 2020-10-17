package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/styles"
	prompt "github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/c-bata/kube-prompt/internal/debug"
	"github.com/c-bata/kube-prompt/kube"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	_ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var (
	version  string
	revision string
	style    string
)

func isStyleSupported(item string) bool {
	for _, availableStyle := range styles.Names() {
		if item == availableStyle {
			return true
		}
	}
	return false
}

func main() {
	styleUsage := fmt.Sprintf("The chroma style to use when coloring json/yaml output. Available options: %s", styles.Names())
	flag.StringVar(&style, "style", "algol", styleUsage)
	flag.Parse()

	if !isStyleSupported(style) {
		fmt.Fprintf(os.Stderr, "%s style does not exist\n", style)
		os.Exit(1)
	}

	c, err := kube.NewCompleter()
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	defer debug.Teardown()
	fmt.Printf("kube-prompt %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		kube.Executor(style),
		c.Complete,
		prompt.OptionTitle("kube-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
