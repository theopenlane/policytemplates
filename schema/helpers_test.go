package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected map[string]any
		wantErr  bool
	}{
		{
			name: "simple struct",
			input: struct {
				Field1 string
				Field2 float64
			}{
				Field1: "value1",
				Field2: 2,
			},
			expected: map[string]any{
				"Field1": "value1",
				"Field2": float64(2),
			},
			wantErr: false,
		},
		{
			name: "nested struct",
			input: struct {
				Field1 string
				Field2 struct {
					SubField1 string
					SubField2 string
				}
			}{
				Field1: "value1",
				Field2: struct {
					SubField1 string
					SubField2 string
				}{
					SubField1: "subvalue1",
					SubField2: "subvalue2",
				},
			},
			expected: map[string]any{
				"Field1": "value1",
				"Field2": map[string]any{
					"SubField1": "subvalue1",
					"SubField2": "subvalue2",
				},
			},
			wantErr: false,
		},
		{
			name:     "invalid input",
			input:    func() {},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := structToMap(tt.input)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
