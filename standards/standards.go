package standards

// Standard represents a standard with a name, version, web link, and a list of controls
// All controls must have a name, version, and at least one control
// This will be validated by the jsonschema
type Standard[T any] struct {
	Name     string       `json:"name"`
	Version  string       `json:"version"`
	WebLink  string       `json:"web_link,omitempty"`
	Controls []Control[T] `json:"controls"`
}

// Control is the fields that a control may have, all controls must have a ref code (the unique identifier) and a category
// all other fields are optional
// Framework specific metadata can be added to the metadata field using the appropriate type
type Control[T any] struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	RefCode     string       `json:"ref_code"`
	Category    string       `json:"category"`
	Subcategory string       `json:"subcategory,omitempty"`
	SubControls []Control[T] `json:"sub_controls,omitempty"`
	MetaData    T            `json:"metadata,omitempty"`
}

// GenerateSOC2Standards generates SOC 2 standards from the CSV file and validates them against the schema
func GenerateSOC2Standards() (std Standard[SOC2Metadata], err error) {
	std = Standard[SOC2Metadata]{
		Name:    "SOC 2",
		Version: "2017 - with revised points of focus 2022",
		WebLink: "https://www.aicpa-cima.com/resources/download/2017-trust-services-criteria-with-revised-points-of-focus-2022",
	}

	std.Controls, err = soc2ParseCSV("references/soc2/controls.csv")
	if err != nil {
		return
	}

	if err := validateStandards(std); err != nil {
		return std, err
	}

	return
}

// GenerateNISTCSFStandards generates NIST CSF standards from the CSV file and validates them against the schema
func GenerateNISTCSFStandards() (std Standard[NISTCSFMetadata], err error) {
	std = Standard[NISTCSFMetadata]{
		Name:    "NIST Cybersecurity Framework",
		Version: "1.1",
		WebLink: "https://nvlpubs.nist.gov/nistpubs/CSWP/NIST.CSWP.04162018.pdf",
	}

	std.Controls, err = nistCsfParseCSV("references/nist-csf/controls-1.1.csv")
	if err != nil {
		return
	}

	if err := validateStandards(std); err != nil {
		return std, err
	}

	return
}

// GenerateNist80053Standards generates NIST 800-53 standards from the CSV file and validates them against the schema
func GenerateNist80053Standards() (std Standard[NIST80053Metadata], err error) {
	std = Standard[NIST80053Metadata]{
		Name:    "NIST SP 800-53",
		Version: "Rev 5",
		WebLink: "https://csrc.nist.gov/pubs/sp/800/53/r5/upd1/final",
	}

	std.Controls, err = nist80053ParseCSV("references/nist-80053/controls-r5.csv")
	if err != nil {
		return
	}

	if err := validateStandards(std); err != nil {
		return std, err
	}

	return
}
