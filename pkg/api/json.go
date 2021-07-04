package api

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/mimuret/golang-iij-dpf/pkg/meta"
)

func MarshalMap(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
	}.Froze().Marshal(body)
}

func UnmarshalRead(bs []byte, o interface{}) error {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "read",
	}.Froze().Unmarshal(bs, o)
}

func MarshalCreate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "create",
	}.Froze().Marshal(body)
}

func MarshalUpdate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "update",
	}.Froze().Marshal(body)
}

func MarshalApply(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "apply",
	}.Froze().Marshal(body)
}

type OutputFrame struct {
	meta.KindVersion `json:",inline"`
	Spec             Spec `json:"spec"`
}

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
