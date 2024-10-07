package soc2

import (
	"testing"
)

func TestGetCategory(t *testing.T) {
	tests := []struct {
		refCode string
		want    string
	}{
		{"CC.1.1", "Security"},
		{"A.2.3", "Availability"},
		{"PI.3.4", "Processing Integrity"},
		{"C.4.5", "Confidentiality"},
		{"P.5.6", "Privacy"},
		{"X.1.1", ""},     // Test case for an unknown category
		{"CCC.1.1.1", ""}, // Test case for an unknown category
	}

	for _, tt := range tests {
		t.Run(tt.refCode, func(t *testing.T) {
			if got := getCategory(tt.refCode); got != tt.want {
				t.Errorf("getCategory(%q) = %v, want %v", tt.refCode, got, tt.want)
			}
		})
	}
}
