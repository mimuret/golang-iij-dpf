package api

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/mimuret/golang-iij-dpf/pkg/meta"
)

// Unmarshal api response
func UnmarshalRead(bs []byte, o interface{}) error {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "read",
	}.Froze().Unmarshal(bs, o)
}

// Marshal for create request
func MarshalCreate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "create",
	}.Froze().Marshal(body)
}

// Marshal for update request
func MarshalUpdate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "update",
	}.Froze().Marshal(body)
}

// Marshal for apply request
func MarshalApply(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "apply",
	}.Froze().Marshal(body)
}

// file format frame
type OutputFrame struct {
	meta.KindVersion `json:",inline"`
	Spec             Spec `json:"spec"`
}

// Marshal for file format
func MarshalOutput(spec Spec) ([]byte, error) {
	t := reflect.TypeOf(spec)
	t = t.Elem()
	out := &OutputFrame{
		KindVersion: meta.KindVersion{
			Kind:       t.Name(),
			APIVersion: spec.GetGroup(),
		},
		Spec: spec,
	}
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        false,
		TagKey:                 "json",
	}.Froze().Marshal(out)
}

// UnMarshal for file format
func UnMarshalInput(bs []byte, spec Spec) error {
	t := reflect.TypeOf(spec)
	t = t.Elem()
	out := &OutputFrame{
		KindVersion: meta.KindVersion{
			Kind:       t.Name(),
			APIVersion: spec.GetGroup(),
		},
		Spec: spec,
	}

	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        false,
		TagKey:                 "json",
	}.Froze().Unmarshal(bs, out)
}
