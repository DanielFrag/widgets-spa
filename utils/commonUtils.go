package utils

import (
	"bytes"
	"encoding/json"
)

//FormatJSON format the map[string]interface{} to json
func FormatJSON(m map[string]interface{}) []byte {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(m)
	return b.Bytes()
}
