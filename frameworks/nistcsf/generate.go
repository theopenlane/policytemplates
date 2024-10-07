package nistcsf

import (
	"net/url"

	"github.com/theopenlane/policytemplates/schema"
)

const (
	// Name of the framework
	Name = "NIST Cybersecurity Framework"
	// Framework is the slug representation of the framework
	Framework = "nist_csf"
	// Version of the framework standards
	Version = "1.1"
)

// GenerateNISTCSFStandards generates NIST CSF standards from the CSV file and validates them against the schema
func Generate() (std schema.Framework[Metadata], err error) {
	link, err := url.Parse("https://nvlpubs.nist.gov/nistpubs/CSWP/NIST.CSWP.04162018.pdf")
	if err != nil {
		return
	}

	std = schema.Framework[Metadata]{
		Name:      Name,
		Framework: Framework,
		Version:   Version,
		WebLink:   link.String(),
	}

	std.Controls, err = nistCsfParseCSV("references/nist-csf/controls-1.1.csv")
	if err != nil {
		return
	}

	if err := schema.Validate(std); err != nil {
		return std, err
	}

	return
}
