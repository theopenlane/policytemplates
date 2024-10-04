package standards

import (
	"encoding/csv"
	"os"
	"strings"
)

// SOC2Metadata contains the metadata for SOC2 Control types
type SOC2Metadata struct {
	PointsOfFocus []string `json:"points_of_focus,omitempty"`
}

// socCategoryMapping maps the SOC 2 categories to the appropriate category
var socCategoryMapping = map[string]string{
	"CC": "Security",
	"A":  "Availability",
	"PI": "Processing Integrity",
	"C":  "Confidentiality",
	"P":  "Privacy",
}

// soc2ParseCSV parses the SOC2 CSV file and returns a slice of controls in a standard format
func soc2ParseCSV(file string) (s []Control[SOC2Metadata], err error) {
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
		if i == 0 {
			continue
		}

		seriesID := record[0]

		pof := strings.Split(record[5], "\n– ")
		pof = pof[1:]

		refCode := record[3]
		control := Control[SOC2Metadata]{
			RefCode:     refCode,
			Category:    getCategory(refCode),
			Subcategory: record[1],
			Description: record[4],
			MetaData: SOC2Metadata{
				PointsOfFocus: pof,
			},
		}

		s = appendSubControl(seriesID, control, s)
	}

	return
}

// getCategory returns the category for the SOC2 control based on the ref code
func getCategory(refCode string) string {
	part := strings.Split(refCode, ".")

	// remove the ID from the ref code
	code := strings.TrimFunc(part[0], func(r rune) bool {
		return r >= '0' && r <= '9'
	})

	for c, cat := range socCategoryMapping {
		if strings.EqualFold(code, c) {
			return cat
		}
	}

	return ""
}
