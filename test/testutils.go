package test

import (
	"bytes"
	"encoding/json"
)

func CompactJson(jsonData []byte) (string, error) {
	var compactedData bytes.Buffer
	err := json.Compact(&compactedData, jsonData)

	return compactedData.String(), err
}