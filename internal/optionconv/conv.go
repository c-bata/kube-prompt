package optionconv

import (
	"errors"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

func GetOptionsFromHelpText(help string) (options string, err error) {
	x := strings.Split(help, "\nOptions:")
	if len(x) < 2 {
		return "", errors.New("parse error")
	}
	y := strings.Split(x[1], "\n\n")
	options = strings.Trim(y[0], "\n")
	return options, nil
}

func SplitOptions(options string) []string {
	lines := strings.Split(options, "\n")
	results := make([]string, 0, len(lines))
	var index int
	for _, l := range lines {
		if strings.HasPrefix(l, "  ") || index == 0 {
			results = append(results, strings.TrimSpace(l))
			index++
		} else {
			results[index-1] += " " + l
		}
	}
	return results
}

func convertToSuggest(flagLine string) []prompt.Suggest {
	x := strings.SplitN(flagLine, ": ", 2)
	key := x[0]
	description := x[1]

	var keys []string
	if strings.Contains(key, ", ") {
		keys = strings.Split(key, ", ")
	} else {
		keys = []string{key}
	}
	suggests := make([]prompt.Suggest, len(keys))
	for i := range keys {
		if strings.Contains(keys[i], "=") {
			keys[i] = strings.Split(keys[i], "=")[0]
		}
		keys[i] = strings.TrimSpace(keys[i])
		suggests[i] = prompt.Suggest{Text: keys[i], Description: description}
	}
	return suggests
}

func ConvertToSuggestions(options []string) []prompt.Suggest {
	suggestions := make([]prompt.Suggest, 0, len(options))
	for i := range options {
		x := convertToSuggest(options[i])
		for j := range x {
			suggestions = append(suggestions, x[j])
		}
	}
	return suggestions
}
