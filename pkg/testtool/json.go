package testtool

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

func MarshalMap(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
	}.Froze().Marshal(body)
}

func UnmarshalToMapString(a json.RawMessage) map[string]interface{} {
	o := map[string]interface{}{}
	if err := json.Unmarshal(a, &o); err != nil {
		return map[string]interface{}{"err": err}
	}
	return o
}
