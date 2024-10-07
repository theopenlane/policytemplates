package iso27001

import (
	"net/url"

	"github.com/theopenlane/policytemplates/schema"
)

const (
	// Name of the framework
	Name = "ISO 27001"
	// Framework is the slug representation of the framework
	Framework = "iso27001"
	// Version of the framework standards
	Version = "2022"
)

// Generate generates ISO27001 standards from the CSV file and validates them against the schema
func Generate() (std schema.Framework[Metadata], err error) {
	link, err := url.Parse("https://www.iso.org/standard/27001")
	if err != nil {
		return
	}

	std = schema.Framework[Metadata]{
		Name:      Name,
		Framework: Framework,
		Version:   Version,
		WebLink:   link.String(),
	}

	std.Controls, err = parseCSV("references/iso27001/controls-2022.csv")
	if err != nil {
		return
	}

	if err := schema.Validate(std); err != nil {
		return std, err
	}

	return
}
