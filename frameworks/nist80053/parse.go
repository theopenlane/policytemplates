package nist80053

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/theopenlane/policytemplates/frameworks"
	"github.com/theopenlane/policytemplates/schema"
)

// Metadata contains the metadata for a NIST 800-53 control type
type Metadata struct {
	Discussion      []string `json:"discussion,omitempty"`
	RelatedControls []string `json:"related_controls,omitempty"`
}

// parseCSV parses the NIST 800-53 CSV file and returns a slice of controls in a standard format
func parseCSV(file string) (s []schema.Control[Metadata], err error) {
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
	for i, record := range records {
		// first row is the header
		if i == 0 {
			continue
		}

		refCode := record[0]
		category := record[1]
		name := category

		if strings.Contains(name, "|") {
			parts := strings.Split(category, "|")
			category = strings.TrimSpace(parts[0])
			name = strings.TrimSpace(parts[1])
		}

		description := record[2]

		discussion := strings.Split(record[3], "\n")

		var relatedControls []string
		if record[4] != "None." {
			relatedControls = strings.Split(record[4], ",")

			for i, c := range relatedControls {
				relatedControls[i] = strings.Trim(strings.TrimSpace(c), ".")
			}
		}

		// remove withdrawn controls
		if !strings.Contains(description, "Withdrawn") {
			control := schema.Control[Metadata]{
				RefCode:     refCode,
				Category:    category,
				Name:        name,
				Description: description,
				Metadata: Metadata{
					Discussion:      discussion,
					RelatedControls: relatedControls,
				},
			}

			parentID := refCode
			if strings.Contains(refCode, ")") && strings.Contains(refCode, "(") {
				parentID = strings.Split(refCode, "(")[0]
			}

			s = frameworks.AppendSubControl(parentID, control, s)
		}
	}

	return
}
