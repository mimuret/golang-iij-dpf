package api

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/mimuret/golang-iij-dpf/pkg/meta"
)

var (
	jsonAdapter    = &JsonAPIAdapter{}
	UnmarshalRead  = jsonAdapter.UnmarshalRead
	MarshalCreate  = jsonAdapter.MarshalCreate
	MarshalUpdate  = jsonAdapter.MarshalUpdate
	MarshalApply   = jsonAdapter.MarshalApply
	MarshalOutput  = jsonAdapter.MarshalOutput
	UnMarshalInput = jsonAdapter.UnMarshalInput
)

type JsonApiInterface interface {
	UnmarshalRead(bs []byte, o interface{}) error
	MarshalCreate(body interface{}) ([]byte, error)
	MarshalUpdate(body interface{}) ([]byte, error)
	MarshalApply(body interface{}) ([]byte, error)
}

type JsonFileInerface interface {
	MarshalOutput(spec Spec) ([]byte, error)
	UnMarshalInput(bs []byte, obj Object) error
}

type JsonAPIAdapter struct {
}

// Unmarshal api response
func (j *JsonAPIAdapter) UnmarshalRead(bs []byte, o interface{}) error {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "read",
	}.Froze().Unmarshal(bs, o)
}

// Marshal for create request
func (j *JsonAPIAdapter) MarshalCreate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "create",
	}.Froze().Marshal(body)
}

// Marshal for update request
func (j *JsonAPIAdapter) MarshalUpdate(body interface{}) ([]byte, error) {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        true,
		TagKey:                 "update",
	}.Froze().Marshal(body)
}

// Marshal for apply request
func (j *JsonAPIAdapter) MarshalApply(body interface{}) ([]byte, error) {
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
	Resource         Object `json:"resource"`
}

// Marshal for file format
func (j *JsonAPIAdapter) MarshalOutput(spec Spec) ([]byte, error) {
	t := reflect.TypeOf(spec)
	t = t.Elem()
	out := &OutputFrame{
		KindVersion: meta.KindVersion{
			Kind:       t.Name(),
			APIVersion: spec.GetGroup(),
		},
		Resource: spec,
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
func (j *JsonAPIAdapter) UnMarshalInput(bs []byte, obj Object) error {
	out := &OutputFrame{
		Resource: obj,
	}

	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		OnlyTaggedField:        false,
		TagKey:                 "json",
	}.Froze().Unmarshal(bs, out)
}
