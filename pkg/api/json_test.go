package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
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
