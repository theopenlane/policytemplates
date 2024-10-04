package standards

import (
	"encoding/csv"
	"os"
	"strings"
)

// NIST80053Metadata contains the metadata for a NIST 800-53 control type
type NIST80053Metadata struct {
	Discussion      []string `json:"discussion,omitempty"`
	RelatedControls []string `json:"related_controls,omitempty"`
}

// nist80053ParseCSV parses the NIST 800-53 CSV file and returns a slice of controls in a standard format
func nist80053ParseCSV(file string) (s []Control[NIST80053Metadata], err error) {
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
			control := Control[NIST80053Metadata]{
				RefCode:     refCode,
				Category:    category,
				Name:        name,
				Description: description,
				MetaData: NIST80053Metadata{
					Discussion:      discussion,
					RelatedControls: relatedControls,
				},
			}

			parentID := refCode
			if strings.Contains(refCode, ")") && strings.Contains(refCode, "(") {
				parentID = strings.Split(refCode, "(")[0]
			}

			s = appendSubControl(parentID, control, s)
		}
	}

	return
}
