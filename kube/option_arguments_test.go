package kube

import (
	"testing"

	"github.com/c-bata/go-prompt"
)

func TestGetPreviousOption(t *testing.T) {
	b1 := prompt.NewBuffer()
	b2 := prompt.NewBuffer()
	b3 := prompt.NewBuffer()
	b1.InsertText("apply -f ./", false, true)
	b2.InsertText("apply ./", false, true)
	b3.InsertText("apply", false, true)
	var scenarioTable = []struct {
		buf             *prompt.Buffer
		expectedCommand string
		expectedFound   bool
		expectedOption  string
	}{
		{
			buf:             b1,
			expectedCommand: "apply",
			expectedFound:   true,
			expectedOption:  "-f",
		},
		{
			buf:             b2,
			expectedCommand: "",
			expectedFound:   false,
			expectedOption:  "",
		},
		{
			buf:             b3,
			expectedCommand: "",
			expectedFound:   false,
			expectedOption:  "",
		},
	}

	for _, s := range scenarioTable {
		cmd, actual, found := getPreviousOption(*s.buf.Document())
		if cmd != s.expectedCommand {
			t.Errorf("Should be %s, but got %s", s.expectedOption, actual)
		}
		if found != s.expectedFound {
			t.Errorf("Should be %v, but got found=%v", s.expectedFound, found)
		}
		if actual != s.expectedOption {
			t.Errorf("Should be %s, but got %s", s.expectedOption, actual)
		}
	}
}
