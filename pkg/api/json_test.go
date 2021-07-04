package api_test

import (
	"encoding/json"
	"testing"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestUnMarshalInput(t *testing.T) {
	tc := json.RawMessage(`{"kind": "TestSpec","apiVersion":"test","spec":{"Id": "bbbb", "Name":"hoge", "Number": 2}}`)
	unm := &TestSpec{}
	err := api.UnMarshalInput(tc, unm)
	if assert.NoError(t, err) {
		spec := &TestSpec{
			Id:     "bbbb",
			Name:   "hoge",
			Number: 2,
		}
		assert.Equal(t, unm, spec)
	}
}

func TestMarshalOutput(t *testing.T) {

	spec := &TestSpec{
		Id:     "10202",
		Name:   "www",
		Number: 1,
	}
	tc := json.RawMessage(`{"kind":"TestSpec","apiVersion":"test","spec":{"Id":"10202","Name":"www","Number":1}}`)

	bs, err := api.MarshalOutput(spec)
	if assert.NoError(t, err) {
		assert.Equal(t, string(bs), string(tc))
	}

}
