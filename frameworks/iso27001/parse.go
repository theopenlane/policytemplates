package iso27001

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/theopenlane/policytemplates/frameworks"
	"github.com/theopenlane/policytemplates/schema"
)

// Metadata contains the metadata for ISO27001 Control types
type Metadata struct {
}

// themeMapping maps the ISO27001 themes to the appropriate category
var themeMapping = map[string]string{
	"A.5": "Organization Controls",
	"A.6": "People Controls",
	"A.7": "Physical Controls",
	"A.8": "Technology Controls",
}

// parseCSV parses the ISO27001 CSV file and returns a slice of controls in a standard format
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
	var (
		category    string
		parentRefID string
	)

	for i, record := range records {
		if i == 0 {
			continue
		}

		if record[0] != "" {
			theme := strings.Split(record[0], "-")
			parentRefID = strings.TrimSpace(theme[0])
			desc := strings.TrimSpace(theme[1])
			category = themeMapping[parentRefID]

			// append the parent control
			control := schema.Control[Metadata]{
				RefCode:     parentRefID,
				Category:    category,
				Description: desc,
				Name:        category,
				Metadata:    Metadata{},
			}

			s = frameworks.AppendSubControl(parentRefID, control, s)
		}

		refCode := record[1]
		control := schema.Control[Metadata]{
			RefCode:     refCode,
			Category:    category,
			Name:        record[2],
			Description: record[3],
			Metadata:    Metadata{},
		}

		s = frameworks.AppendSubControl(parentRefID, control, s)
	}

	return
}
