package test

import (
	"encoding/json"
)

func UnmarshalToMapString(a json.RawMessage) map[string]interface{} {
	o := map[string]interface{}{}
	if err := json.Unmarshal(a, &o); err != nil {
		return map[string]interface{}{"err": err}
	}
	return o
}
