package standards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCategory(t *testing.T) {
	tests := []struct {
		input           string
		expectedCat     string
		expectedRecCode string
		expectedDesc    string
	}{
		{
			input:           "Identify (ID): The data ...",
			expectedCat:     "Identify",
			expectedRecCode: "ID",
			expectedDesc:    "The data ...",
		},
		{
			input:       "Identify",
			expectedCat: "Identify",
		},
		{
			input:           "Protect (PR)",
			expectedCat:     "Protect",
			expectedRecCode: "PR",
			expectedDesc:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			category, recCode, description := parseCategory(tt.input)
			assert.Equal(t, tt.expectedCat, category)
			assert.Equal(t, tt.expectedRecCode, recCode)
			assert.Equal(t, tt.expectedDesc, description)
		})
	}
}

func TestAddReferencesToControl(t *testing.T) {
	reference := "Reference 1"

	tests := []struct {
		name             string
		initialControls  []Control[NISTCSFMetadata]
		refCode          string
		expectedControls []Control[NISTCSFMetadata]
	}{
		{
			name: "add reference to existing subcontrol",
			initialControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					SubControls: []Control[NISTCSFMetadata]{
						{
							RefCode:     "ID.AM",
							Category:    "Identify",
							Subcategory: "Asset Management",
						},
					},
				},
			},
			refCode: "ID.AM",
			expectedControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					SubControls: []Control[NISTCSFMetadata]{
						{
							RefCode:     "ID.AM",
							Category:    "Identify",
							Subcategory: "Asset Management",
							MetaData: NISTCSFMetadata{
								References: []string{reference},
							},
						},
					},
				},
			},
		},
		{
			name: "add reference to existing control",
			initialControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
				},
			},
			refCode: "ID",
			expectedControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					MetaData: NISTCSFMetadata{
						References: []string{reference},
					},
				},
			},
		},
		{
			name: "append reference to existing references",
			initialControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					MetaData: NISTCSFMetadata{
						References: []string{"Other Reference"},
					},
				},
			},
			refCode: "ID",
			expectedControls: []Control[NISTCSFMetadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					MetaData: NISTCSFMetadata{
						References: []string{"Other Reference", reference},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addReferencesToControl(reference, tt.refCode, tt.initialControls)
			assert.Equal(t, tt.expectedControls, result)
		})
	}
}
