package testtool

import jsoniter "github.com/json-iterator/go"

func MarshalMap(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
	}.Froze().Marshal(body)
}
