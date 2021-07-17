package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/stretchr/testify/assert"
)

var _ api.Spec = &JsonTest{}

type JsonTest struct {
	Name string `json:"name" read:"read_name" create:"create_name" update:"update_name" apply:"apply_name"`
}

func (j *JsonTest) DeepCopyObject() api.Object { return &JsonTest{Name: j.Name} }
func (JsonTest) GetName() string               { return "jsontests" }
func (JsonTest) GetGroup() string              { return "tests" }
func (JsonTest) GetPathMethod(a api.Action) (string, string) {
	return a.ToMethod(), "/tests/jsontests"
}

var _ = Describe("json", func() {
	var (
		bs    []byte
		err   error
		value JsonTest
	)
	BeforeEach(func() {
		value = JsonTest{Name: "hogehoge"}
	})
	Describe("for API", func() {
		Context("UnmarshalRead(tag name is `read`)", func() {
			BeforeEach(func() {
				err = api.UnmarshalRead([]byte(`{"read_name": "book"}`), &value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(value.Name).To(Equal("book"))
			})
		})
		Context("MarshalCreate(tag name is `create`)", func() {
			BeforeEach(func() {
				bs, err = api.MarshalCreate(&value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(string(bs)).To(Equal(`{"create_name":"hogehoge"}`))
			})
		})
		Context("MarshalUpdate(tag name is `update`)", func() {
			BeforeEach(func() {
				bs, err = api.MarshalUpdate(&value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(string(bs)).To(Equal(`{"update_name":"hogehoge"}`))
			})
		})
		Context("MarshalApply (tag name is `apply`)", func() {
			BeforeEach(func() {
				bs, err = api.MarshalApply(&value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(string(bs)).To(Equal(`{"apply_name":"hogehoge"}`))
			})
		})
	})
	Describe("for file", func() {
		Context("MarshalOutput(tag name is `json`)", func() {
			BeforeEach(func() {
				bs, err = api.MarshalOutput(&value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(string(bs)).To(Equal(`{"kind":"JsonTest","apiVersion":"tests","spec":{"name":"hogehoge"}}`))
			})
		})
		Context("UnMarshalInput(tag name is `json`)", func() {
			BeforeEach(func() {
				err = api.UnMarshalInput([]byte(`{"kind": "JsonTest", "apiVersion": "tests", "spec": {"name": "book"}}`), &value)
			})
			It("can read `name`", func() {
				Expect(err).To(Succeed())
				Expect(value.Name).To(Equal("book"))
			})
		})
	})
})

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
