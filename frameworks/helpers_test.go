package frameworks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/theopenlane/policytemplates/schema"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		controls []schema.Control[any]
		refCode  string
		expected bool
	}{
		{
			name: "control exists",
			controls: []schema.Control[any]{
				{RefCode: "C1"},
				{RefCode: "C2"},
			},
			refCode:  "C1",
			expected: true,
		},
		{
			name: "control does not exist",
			controls: []schema.Control[any]{
				{RefCode: "C1"},
				{RefCode: "C2"},
			},
			refCode:  "C3",
			expected: false,
		},
		{
			name:     "empty controls",
			controls: []schema.Control[any]{},
			refCode:  "C1",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.controls, tt.refCode)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestAppendSubControl(t *testing.T) {
	tests := []struct {
		name     string
		parentID string
		control  schema.Control[any]
		controls []schema.Control[any]
		expected []schema.Control[any]
	}{
		{
			name:     "append to empty controls",
			parentID: "P1",
			control:  schema.Control[any]{RefCode: "C1"},
			controls: []schema.Control[any]{},
			expected: []schema.Control[any]{
				{RefCode: "C1"},
			},
		},
		{
			name:     "append to existing parent control",
			parentID: "P1",
			control:  schema.Control[any]{RefCode: "C2"},
			controls: []schema.Control[any]{
				{RefCode: "P1"},
			},
			expected: []schema.Control[any]{
				{
					RefCode: "P1",
					SubControls: []schema.Control[any]{
						{RefCode: "C2"},
					},
				},
			},
		},
		{
			name:     "append to nested sub control",
			parentID: "C1",
			control:  schema.Control[any]{RefCode: "C3"},
			controls: []schema.Control[any]{
				{
					RefCode: "P1",
					SubControls: []schema.Control[any]{
						{RefCode: "C1"},
					},
				},
			},
			expected: []schema.Control[any]{
				{
					RefCode: "P1",
					SubControls: []schema.Control[any]{
						{
							RefCode: "C1",
							SubControls: []schema.Control[any]{
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
			control:  schema.Control[any]{RefCode: "C4"},
			controls: []schema.Control[any]{
				{RefCode: "P1"},
			},
			expected: []schema.Control[any]{
				{RefCode: "P1"},
				{RefCode: "C4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AppendSubControl(tt.parentID, tt.control, tt.controls)
			assert.Equal(t, tt.expected, result)
		})
	}
}
