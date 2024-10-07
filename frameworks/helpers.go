package frameworks

import (
	"github.com/theopenlane/policytemplates/schema"
)

// Contains checks if a control with the given ref code exists in the slice of controls
func Contains[T any](s []schema.Control[T], i string) bool {
	for _, v := range s {
		if v.RefCode == i {
			return true
		}
	}

	return false
}

// AppendSubControl appends a control to the slice of controls based on the parent control ref code
func AppendSubControl[T any](parentID string, c schema.Control[T], s []schema.Control[T]) []schema.Control[T] {
	for _, sub := range s {
		if Contains(sub.SubControls, parentID) {
			sub.SubControls = AppendSubControl(parentID, c, sub.SubControls)

			return s
		}
	}

	if !Contains(s, parentID) {
		s = append(s, c)

		return s
	}

	for i, v := range s {
		if v.RefCode == parentID {
			s[i].SubControls = append(s[i].SubControls, c)
		}
	}

	return s
}
