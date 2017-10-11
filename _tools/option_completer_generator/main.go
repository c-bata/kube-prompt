package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func getFlags(help string) (flags, globalFlags string) {
	// flag options starts with "Flags:" and ends with "Global Flags:"
	x := strings.Split(help, "\nFlags:")
	if len(x) < 2 {
		fmt.Println("Cannot split with Flags\n" + help)
		os.Exit(1)
	}
	y := strings.Split(x[1], "\nGlobal Flags:")
	flags = strings.Trim(y[0], "\n")
	globalFlags = strings.Trim(y[1], "\n")
	return
}

func convertToSuggest(flagLine string) []prompt.Suggest {
	x := strings.Split(flagLine, ": ")
	key := x[0]
	val := strings.Join(x[1:], ": ")

	key = strings.Split(key, "[")[0]
	key = strings.Trim(key, " [=")
	keys := strings.Split(key, ", ")
	suggests := make([]prompt.Suggest, len(keys))

	for i := range keys {
		suggests[i] = prompt.Suggest{Text: keys[i], Description: val}
	}
	return suggests
}

func convertToSuggestions(flagText string) []prompt.Suggest {
	flags := strings.Split(strings.Trim(flagText, "\n"), "\n")
	suggestions := make([]prompt.Suggest, 0, len(flags))
	for i := range flags {
		x := convertToSuggest(flags[i])
		for j := range x {
			suggestions = append(suggestions, x[j])
		}
	}
	return suggestions
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	f, gf := getFlags(string(bytes))
	fmt.Printf("FLAGS:\n%#v\n", convertToSuggestions(f))
	fmt.Printf("GLOBAL FLAGS:\n%#v\n", convertToSuggestions(gf))
}
