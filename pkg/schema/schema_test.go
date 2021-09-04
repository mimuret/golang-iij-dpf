package schema_test

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ apis.Spec = &TestSpec{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type TestSpec struct {
	Id     string `read:"id"`
	Name   string `read:"name"`
	Number int64  `read:"number"`
}

func (t *TestSpec) GetGroup() string { return "test" }
func (t *TestSpec) GetName() string  { return "tests" }
func (t *TestSpec) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), "/tests"
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/tests/%s", t.Id)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/tests/%s/cancel", t.Id)
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/tests/%s/apply", t.Id)
	}
	return "", ""
}
func (t *TestSpec) DeepCopyTestSpec() *TestSpec {
	res := &TestSpec{}
	*res = *t
	return res
}

func (t *TestSpec) DeepCopyObject() api.Object {
	return t.DeepCopyTestSpec()
}

func (t *TestSpec) SetPathParams(...interface{}) error {
	return nil
}

type ErrSpec struct {
	Id string
}

func (t ErrSpec) GetGroup() string { return "test" }
func (t ErrSpec) GetName() string  { return "tests" }
func (t ErrSpec) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), "/tests"
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/tests/%s", t.Id)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/tests/%s/cancel", t.Id)
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/tests/%s/apply", t.Id)
	}
	return "", ""
}
func (t ErrSpec) SetPathParams(...interface{}) error {
	return nil
}

func (t ErrSpec) DeepCopyObject() api.Object {
	return nil
}

var _ = Describe("Register", func() {
	Context("Add", func() {
		var (
			register *schema.Register
		)
		BeforeEach(func() {
			schema.SchemaSet = schema.NewSchemaSet()
			register = schema.NewRegister("test")
		})
		It("can add first time", func() {
			register.Add(&TestSpec{})
		})
		It("can not add not pointer", func() {
			Expect(func() { register.Add(ErrSpec{}) }).To(Panic())
		})
		It("can not add same type", func() {
			register.Add(&TestSpec{})
			Expect(func() { register.Add(&TestSpec{}) }).To(Panic())
		})
	})
	Context("Parse", func() {
		var (
			register *schema.Register
			obj      api.Object
			err      error
		)
		BeforeEach(func() {
			schema.SchemaSet = schema.NewSchemaSet()
			register = schema.NewRegister("test")
			register.Add(&TestSpec{})
			err = nil
			obj = nil
		})
		When("data is not json", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse json"))
			})
		})
		When("data is json, but not found kind", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{}`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("kind value is not exist"))
			})
		})
		When("data is json, but not found apiVersion", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{"kind": "TestSpec"}`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("apiVersion value is not exist"))
			})
		})
		When("apiVersion is group, if group not found", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{"kind": "TestSpec", "apiVersion": "testtest"}`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("apiVersion `testtest` is not support"))
			})
		})
		When("kind is not found", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{"apiVersion": "test", "kind": "hogehoge"}`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("kind value `hogehoge` is not support"))
			})
		})
		When("failed to parse", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{"apiVersion": "test", "kind": "TestSpec", "resource": {"Id": 0}}`))
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse resource:"))
			})
		})
		When("successful", func() {
			BeforeEach(func() {
				obj, err = schema.SchemaSet.Parse([]byte(`{"apiVersion": "test", "kind": "TestSpec", "resource": {"Id": "hoge"}}`))
			})
			It("returns error", func() {
				Expect(err).To(Succeed())
				tc, ok := obj.(*TestSpec)
				Expect(ok).To(BeTrue())
				Expect(tc.Id).To(Equal("hoge"))
			})
		})
	})
})
