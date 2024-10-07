package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type framework struct {
	Name        string
	Description string
	Value       string
}

var frameworks = []framework{
	{
		Name:        "SOC2",
		Description: "2017 Trust Services Criteria for Security, Availability, Processing Integrity, Confidentiality, and Privacy (with Revised Points of Focus – 2022)",
		Value:       "soc2",
	},
	{
		Name:        "NIST CSF",
		Description: "Cybersecurity Framework (CSF) Version 1.1",
		Value:       "nist-csf-1.1",
	},
	{
		Name:        "NIST 800-53",
		Description: "Security and Privacy Controls for Federal Information Systems and Organizations (Revision 5)",
		Value:       "nist-800-53-rev5",
	},
	{
		Name:        "ISO 27001:2022",
		Description: "International Organization for Standardization (ISO) 27001:2022",
		Value:       "iso27001:2022",
	},
}

var selectFramework = &promptui.SelectTemplates{
	Label:    "{{ .Name }} ",
	Active:   "\U0001F449 {{ .Name | cyan }}",
	Inactive: "  {{ .Name | cyan }}",
	Selected: "\U0001F449 {{ .Name | green | cyan }}",
	Details: `
{{ "Description:" | faint }}	{{ .Description }}`,
}

var searcherFramework = func(input string, index int) bool {
	framework := frameworks[index]
	name := strings.ReplaceAll(strings.ToLower(framework.Name), " ", "")
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")

	return strings.Contains(name, input)
}

func Frameworks() (string, error) {
	prompt := promptui.Select{
		Label:     "Frameworks:",
		Templates: selectFramework,
		Items:     frameworks,
		Searcher:  searcherFramework,
	}

	i, _, err := prompt.Run()

	return frameworks[i].Value, err
}
