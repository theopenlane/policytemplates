package standards

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		controls []Control[any]
		refCode  string
		expected bool
	}{
		{
			name: "control exists",
			controls: []Control[any]{
				{RefCode: "C1"},
				{RefCode: "C2"},
			},
			refCode:  "C1",
			expected: true,
		},
		{
			name: "control does not exist",
			controls: []Control[any]{
				{RefCode: "C1"},
				{RefCode: "C2"},
			},
			refCode:  "C3",
			expected: false,
		},
		{
			name:     "empty controls",
			controls: []Control[any]{},
			refCode:  "C1",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := contains(tt.controls, tt.refCode)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestAppendSubControl(t *testing.T) {
	tests := []struct {
		name     string
		parentID string
		control  Control[any]
		controls []Control[any]
		expected []Control[any]
	}{
		{
			name:     "append to empty controls",
			parentID: "P1",
			control:  Control[any]{RefCode: "C1"},
			controls: []Control[any]{},
			expected: []Control[any]{
				{RefCode: "C1"},
			},
		},
		{
			name:     "append to existing parent control",
			parentID: "P1",
			control:  Control[any]{RefCode: "C2"},
			controls: []Control[any]{
				{RefCode: "P1"},
			},
			expected: []Control[any]{
				{
					RefCode: "P1",
					SubControls: []Control[any]{
						{RefCode: "C2"},
					},
				},
			},
		},
		{
			name:     "append to nested sub control",
			parentID: "C1",
			control:  Control[any]{RefCode: "C3"},
			controls: []Control[any]{
				{
					RefCode: "P1",
					SubControls: []Control[any]{
						{RefCode: "C1"},
					},
				},
			},
			expected: []Control[any]{
				{
					RefCode: "P1",
					SubControls: []Control[any]{
						{
							RefCode: "C1",
							SubControls: []Control[any]{
								{RefCode: "C3"},
							},
						},
					},
				},
			},
		},
		{
			name:     "append to non-existing parent control",
			parentID: "P2",
			control:  Control[any]{RefCode: "C4"},
			controls: []Control[any]{
				{RefCode: "P1"},
			},
			expected: []Control[any]{
				{RefCode: "P1"},
				{RefCode: "C4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := appendSubControl(tt.parentID, tt.control, tt.controls)
			assert.Equal(t, tt.expected, result)
		})
	}
}
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
