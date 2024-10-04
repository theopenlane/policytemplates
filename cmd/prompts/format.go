package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type format struct {
	Name  string
	Value string
}

var formats = []format{
	{
		Name:  "Save To File",
		Value: "file",
	},
	{
		Name:  "Standard Out - JSON",
		Value: "stdout",
	},
}

var selectFormat = &promptui.SelectTemplates{
	Label:    "{{ .Name }} ",
	Active:   "\U0001F449 {{ .Name | cyan }}",
	Inactive: "  {{ .Name | cyan }}",
	Selected: "\U0001F449 {{ .Name | green | cyan }}",
}

var searcherFormat = func(input string, index int) bool {
	format := formats[index]
	name := strings.ReplaceAll(strings.ToLower(format.Name), " ", "")
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")

	return strings.Contains(name, input)
}

func Formats() (string, error) {
	prompt := promptui.Select{
		Label:     "Output Format:",
		Templates: selectFormat,
		Items:     formats,
		Searcher:  searcherFormat,
	}

	i, _, err := prompt.Run()

	return formats[i].Value, err
}
