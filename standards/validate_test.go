package standards

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/santhosh-tekuri/jsonschema/v6"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateStandards(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	tests := []struct {
		name          string
		controls      Standard[any]
		errorExpected bool
		missingFields []string
	}{
		{
			name: "valid standards",
			controls: Standard[any]{
				Name: "SOC2",
				Controls: []Control[any]{
					{
						RefCode:  "C1",
						Category: "Category 1",
					},
				},
				Version: "2017",
			},

			errorExpected: false,
		},
		{
			name:          "empty standards",
			controls:      Standard[any]{},
			errorExpected: true,
		},
		{
			name: "missing required version",
			controls: Standard[any]{
				Name: "SOC2",
				Controls: []Control[any]{
					{
						RefCode:  "C1",
						Category: "Category 1",
					},
				},
			},
			errorExpected: true,
		},
		{
			name: "missing required name",
			controls: Standard[any]{
				Version: "2017",
				Controls: []Control[any]{
					{
						RefCode:  "C1",
						Category: "Category 1",
					},
				},
			},
			errorExpected: true,
		},
		{
			name: "missing required controls",
			controls: Standard[any]{
				Name:    "SOC2",
				Version: "2017",
			},
			errorExpected: true,
		},
		{
			name: "missing required category",
			controls: Standard[any]{
				Name:    "SOC2",
				Version: "2017",
				Controls: []Control[any]{
					{
						RefCode: "C1",
					},
				},
			},
			errorExpected: true,
		},
		{
			name: "missing required ref code",
			controls: Standard[any]{
				Name:    "SOC2",
				Version: "2017",
				Controls: []Control[any]{
					{
						Category: "Category 1",
					},
				},
			},
			errorExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStandards[any](tt.controls)
			if tt.errorExpected {
				require.Error(t, err)
				assert.IsType(t, &jsonschema.ValidationError{}, err)

				return
			}

			require.NoError(t, err)
		})
	}
}
