package nist80053

import (
	"net/url"

	"github.com/theopenlane/policytemplates/schema"
)

const (
	// Name of the framework
	Name = "NIST SP 800-53"
	// Framework is the slug representation of the framework
	Framework = "nist_800_53"
	// Version of the framework standards
	Version = "Rev 5"
)

// Generate generates NIST 800-53 standards from the CSV file and validates them against the schema
func Generate() (std schema.Framework[Metadata], err error) {
	link, err := url.Parse("https://csrc.nist.gov/pubs/sp/800/53/r5/upd1/final")
	if err != nil {
		return
	}

	std = schema.Framework[Metadata]{
		Name:      Name,
		Framework: Framework,
		Version:   Version,
		WebLink:   link.String(),
	}

	std.Controls, err = parseCSV("references/nist-80053/controls-r5.csv")
	if err != nil {
		return
	}

	if err := schema.Validate(std); err != nil {
		return std, err
	}

	return
}
