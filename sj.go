package sj

import (
	"bytes"
	"encoding/json"
)

func Search(needle string, rawHaystack string) (*[]string, error) {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, []byte(rawHaystack))
	if err != nil {
		return nil, err
	}

	return search(needle, buffer.String())
}

func search(needle string, haystack string) (*[]string, error) {
	if haystack == "" {
		return nil, nil
	}

	var result []string
	var decode map[string]json.RawMessage
	err := json.Unmarshal([]byte(haystack), &decode)
	if err != nil {
		return nil, err
	}

	for k, v := range decode {
		if k == needle {
			result = append(result, string(v))
		} else {
			sr, err := search(needle, string(v))
			if err != nil {
				return nil, err
			}
			result = append(result, *sr...)
		}
	}

	return &result, nil
}
