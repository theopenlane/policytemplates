package standards

import "encoding/json"

// contains checks if a control with the given ref code exists in the slice of controls
func contains[T any](s []Control[T], i string) bool {
	for _, v := range s {
		if v.RefCode == i {
			return true
		}
	}

	return false
}

// appendSubControl appends a control to the slice of controls based on the parent control ref code
func appendSubControl[T any](parentID string, c Control[T], s []Control[T]) []Control[T] {
	for _, sub := range s {
		if contains(sub.SubControls, parentID) {
			sub.SubControls = appendSubControl(parentID, c, sub.SubControls)

			return s
		}
	}

	if !contains(s, parentID) {
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

// structToMap converts a struct to a map
func structToMap(obj any) (map[string]any, error) {
	var result map[string]any

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
