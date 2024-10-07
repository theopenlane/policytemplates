package nistcsf

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/theopenlane/policytemplates/frameworks"
	"github.com/theopenlane/policytemplates/schema"
)

// Metadata contains the metadata for a NIST CSF control type
type Metadata struct {
	References []string `json:"references,omitempty"`
}

// nistCsfParseCSV parses the NIST CSF CSV file and returns a slice of controls in a standard format
func nistCsfParseCSV(file string) (s []schema.Control[Metadata], err error) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	// Parse the file
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return
	}

	// Parse the records
	// set outside the loop to keep track of the current settings when the
	// row is empty
	var (
		category        string
		subcategory     string
		description     string
		refCode         string
		childRefCode    string
		subChildRefCode string
	)

	for i, record := range records {
		// first row is the header
		if i == 0 {
			continue
		}

		// get the function (category) for the top-level control
		if record[0] != "" {
			// function is the top-level category but has no description
			category, refCode, _ = parseCategory(record[0])

			control := schema.Control[Metadata]{
				RefCode:  refCode,
				Category: category,
			}

			// append the control to the slice
			s = frameworks.AppendSubControl(refCode, control, s)
		}

		// get the category (subcategory) for the first child control
		// the category is the same as the parent control
		if record[1] != "" {
			subcategory, childRefCode, description = parseCategory(record[1])
			subControl := schema.Control[Metadata]{
				RefCode:     childRefCode,
				Category:    category,
				Subcategory: subcategory,
				Description: description,
			}

			s = frameworks.AppendSubControl(refCode, subControl, s)
		}

		if record[2] != "" {
			ref := strings.Split(record[2], ":")
			subChildRefCode = ref[0]

			control := schema.Control[Metadata]{
				RefCode:     subChildRefCode,
				Category:    category,
				Subcategory: subcategory,
				Description: strings.TrimSpace(ref[1]),
				Metadata: Metadata{
					References: []string{},
				},
			}

			s = frameworks.AppendSubControl(childRefCode, control, s)
		}

		if record[3] != "" {
			ref := record[3]

			reference := strings.TrimSpace(strings.ReplaceAll(ref, "·       ", ""))

			s = addReferencesToControl(reference, subChildRefCode, s)
		}
	}

	return
}

// parseCategory parses the category from the record and returns the category, ref code, and description
// example record: "Asset Management (ID.AM): The data, personnel, devices, systems, and ..."
// or as simple as "IDENTIFY (ID)"
func parseCategory(record string) (category string, recCode string, description string) {
	rec := strings.Split(record, " (")
	category = strcase.UpperCamelCase(rec[0])

	if len(rec) == 1 {
		return
	}

	rec = strings.Split(rec[1], ":")
	recCode = strings.TrimRight(rec[0], ")")

	if len(rec) > 1 {
		description = strings.TrimSpace(strings.TrimRight(rec[1], ")"))
	}

	return
}

// addReferencesToControl adds a reference to a control (or nested subcontrol) based on the ref code
func addReferencesToControl(reference string, refCode string, s []schema.Control[Metadata]) []schema.Control[Metadata] {
	for i, v := range s {
		if v.RefCode == refCode {
			v.Metadata.References = append(v.Metadata.References, reference)

			s[i] = v

			return s
		}

		for j, sub := range v.SubControls {
			out := addReferencesToControl(reference, refCode, []schema.Control[Metadata]{sub})

			s[i].SubControls[j] = out[0]
		}
	}

	return s
}
