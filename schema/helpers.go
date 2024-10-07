package schema

import "encoding/json"

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
