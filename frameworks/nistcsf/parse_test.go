package nistcsf

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/theopenlane/policytemplates/schema"
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
		initialControls  []schema.Control[Metadata]
		refCode          string
		expectedControls []schema.Control[Metadata]
	}{
		{
			name: "add reference to existing subcontrol",
			initialControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					SubControls: []schema.Control[Metadata]{
						{
							RefCode:     "ID.AM",
							Category:    "Identify",
							Subcategory: "Asset Management",
						},
					},
				},
			},
			refCode: "ID.AM",
			expectedControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					SubControls: []schema.Control[Metadata]{
						{
							RefCode:     "ID.AM",
							Category:    "Identify",
							Subcategory: "Asset Management",
							Metadata: Metadata{
								References: []string{reference},
							},
						},
					},
				},
			},
		},
		{
			name: "add reference to existing control",
			initialControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
				},
			},
			refCode: "ID",
			expectedControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					Metadata: Metadata{
						References: []string{reference},
					},
				},
			},
		},
		{
			name: "append reference to existing references",
			initialControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					Metadata: Metadata{
						References: []string{"Other Reference"},
					},
				},
			},
			refCode: "ID",
			expectedControls: []schema.Control[Metadata]{
				{
					RefCode:  "ID",
					Category: "Identify",
					Metadata: Metadata{
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
