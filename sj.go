package sj

import "encoding/json"

func search(needle string, haystack string) (*string, error) {
	// validate json
	var j json.RawMessage
	err := json.Unmarshal([]byte(haystack), &j)

	if err != nil {
		return nil, err
	}

	var res string
	return &res, nil
}
