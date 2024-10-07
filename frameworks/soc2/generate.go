package soc2

import (
	"net/url"

	"github.com/theopenlane/policytemplates/schema"
)

const (
	// Name of the framework
	Name = "SOC 2"
	// Framework is the slug representation of the framework
	Framework = "soc2"
	// Version of the framework standards
	Version = "2017 - with revised points of focus 2022"
)

// Generate generates SOC 2 standards from the CSV file and validates them against the schema
func Generate() (std schema.Framework[Metadata], err error) {
	link, err := url.Parse("https://www.aicpa-cima.com/resources/download/2017-trust-services-criteria-with-revised-points-of-focus-2022")
	if err != nil {
		return
	}

	std = schema.Framework[Metadata]{
		Name:      Name,
		Framework: Framework,
		Version:   Version,
		WebLink:   link.String(),
	}

	std.Controls, err = parseCSV("references/soc2/controls.csv")
	if err != nil {
		return
	}

	if err := schema.Validate(std); err != nil {
		return std, err
	}

	return
}
