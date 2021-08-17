package testtool

import (
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ api.Spec = &TestSpec{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type TestSpec struct {
	Id     string `read:"id"`
	Name   string `read:"name" update:"name" create:"name" apply:"name"`
	Number int64  `read:"number" update:"number" create:"number" apply:"number"`
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
	if t == nil {
		return nil
	}
	res := &TestSpec{}
	*res = *t
	return res
}

func (t *TestSpec) DeepCopyObject() api.Object {
	return t.DeepCopyTestSpec()
}

var _ api.ListSpec = &TestSpecList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type TestSpecList struct {
	Items []TestSpec `read:"items"`
}

func (t *TestSpecList) DeepCopyObject() api.Object {
	if t == nil {
		return nil
	}
	res := &TestSpecList{}
	for _, item := range t.Items {
		s := item.DeepCopyTestSpec()
		res.Items = append(res.Items, *s)
	}
	return res
}

func (t *TestSpecList) GetGroup() string        { return "test" }
func (t *TestSpecList) GetName() string         { return "tests" }
func (t *TestSpecList) GetItems() interface{}   { return &t.Items }
func (c *TestSpecList) Len() int                { return len(c.Items) }
func (c *TestSpecList) Index(i int) interface{} { return c.Items[i] }
func (c *TestSpecList) GetMaxLimit() int32      { return 10000 }
func (c *TestSpecList) ClearItems()             { c.Items = []TestSpec{} }
func (c *TestSpecList) AddItem(v interface{}) bool {
	if a, ok := v.(TestSpec); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (t *TestSpecList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/tests"
	}
	return "", ""
}
func (t *TestSpecList) Init() {}

var _ api.CountableListSpec = &TestSpecCountableList{}

type TestSpecCountableList struct {
	api.Count
	TestSpecList
}

func (t *TestSpecCountableList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/tests"
	case api.ActionCount:
		return action.ToMethod(), "/tests/count"
	}
	return "", ""
}

func (t *TestSpecCountableList) DeepCopyObject() api.Object {
	if t == nil {
		return nil
	}
	res := &TestSpecCountableList{}
	for _, item := range t.Items {
		s := item.DeepCopyTestSpec()
		res.Items = append(res.Items, *s)
	}
	res.Count = t.Count
	return res
}

// for api.Object
func TestDeepCopyObject(s api.Object, nilSpec api.Object) {
	Context("DeepCopyObject", func() {
		var (
			o api.Object
		)
		BeforeEach(func() {
			o = nil
		})
		When("spec is not nil", func() {
			BeforeEach(func() {
				o = s.DeepCopyObject()
			})
			It("returns copy object", func() {
				Expect(o).To(Equal(s))
			})
		})
		When("spec is nil", func() {
			BeforeEach(func() {
				o = nilSpec.DeepCopyObject()
			})
			It("returns nil", func() {
				Expect(o).To(BeNil())
			})
		})
	})
	Context("DeepCopyObject", func() {
		var (
			o api.Object
		)
		BeforeEach(func() {
			o = nil
		})
		When("spec is not nil", func() {
			BeforeEach(func() {
				o = s.DeepCopyObject()
			})
			It("returns copy object", func() {
				Expect(o).To(Equal(s))
			})
		})
		When("spec is nil", func() {
			BeforeEach(func() {
				o = nilSpec.DeepCopyObject()
			})
			It("returns nil", func() {
				Expect(o).To(BeNil())
			})
		})
	})
}

// for api.Spec
func TestGetPathMethod(spec api.Spec, action api.Action, matchMethod string, matchPath string) {
	var (
		method, path string
	)
	When("action test", func() {
		BeforeEach(func() {
			method, path = spec.GetPathMethod(action)
		})
		It("returns method", func() {
			Expect(method).To(Equal(matchMethod), "action:"+string(action))
		})
		It("returns path", func() {
			Expect(path).To(Equal(matchPath), "action:"+string(action))
		})
	})
}

func TestGetPathMethodForSpec(spec api.Spec, createPath, getPath string) {
	When("action is ActionCreate", func() {
		TestGetPathMethod(spec, api.ActionCreate, http.MethodPost, createPath)
	})
	When("action is ActionRead", func() {
		TestGetPathMethod(spec, api.ActionRead, http.MethodGet, getPath)
	})
	When("action is ActionUpdate", func() {
		TestGetPathMethod(spec, api.ActionUpdate, http.MethodPatch, getPath)
	})
	When("action is ActionDelete", func() {
		TestGetPathMethod(spec, api.ActionDelete, http.MethodDelete, getPath)
	})
	When("action is other", func() {
		TestGetPathMethod(spec, api.ActionApply, "", "")
	})
}

func TestGetPathMethodForList(spec api.ListSpec, listPath string) {
	When("action is ActionList", func() {
		TestGetPathMethod(spec, api.ActionList, http.MethodGet, listPath)
	})
	When("action is other", func() {
		TestGetPathMethod(spec, api.ActionCount, "", "")
	})
}
func TestGetPathMethodForCountableList(spec api.CountableListSpec, listPath string) {
	When("action is ActionList", func() {
		TestGetPathMethod(spec, api.ActionList, http.MethodGet, listPath)
	})
	When("action is ActionCount", func() {
		TestGetPathMethod(spec, api.ActionCount, http.MethodGet, listPath+"/count")
	})
	When("action is other", func() {
		TestGetPathMethod(spec, api.ActionApply, "", "")
	})
}

func TestGetName(s api.Spec, name string) {
	Context("GetName", func() {
		It("returns name", func() {
			Expect(s.GetName()).To(Equal(name))
		})
	})
}
func TestGetGroup(s api.Spec, name string) {
	Context("GetGroup", func() {
		It("returns group name", func() {
			Expect(s.GetGroup()).To(Equal(name))
		})
	})
}

// List Spec
func TestGetItems(s api.ListSpec, items interface{}) {
	Context("GetItems", func() {
		It("returns ItemSlice", func() {
			Expect(s.GetItems()).To(Equal(items))
		})
	})
}

func TestLen(s api.ListSpec, num int) {
	Context("Len", func() {
		It("returns number of items", func() {
			Expect(s.Len()).To(Equal(num))
		})
	})
}

// api.CountableListSpec
func TestGetMaxLimit(s api.CountableListSpec, limit int32) {
	Context("GetMaxLimit", func() {
		It("returns limit", func() {
			Expect(s.GetMaxLimit()).To(Equal(limit))
		})
	})
}
func TestClearItems(s api.CountableListSpec) {
	Context("GetMaxLimit", func() {
		BeforeEach(func() {
			Expect(s.Len()).NotTo((Equal(0)))
			s.ClearItems()
		})
		It("can delete all items", func() {
			Expect(s.Len()).To((Equal(0)))
		})
	})
}
func TestAddItem(s api.CountableListSpec, validData interface{}) {
	Context("AddItem", func() {
		var (
			copy api.CountableListSpec
			len  int
			ok   bool
		)
		BeforeEach(func() {
			copy = s.DeepCopyObject().(api.CountableListSpec)
			len = copy.Len()
		})
		When("add Item", func() {
			BeforeEach(func() {
				ok = copy.AddItem(validData)
			})
			It("can add into list", func() {
				Expect(ok).To((BeTrue()))
				Expect(copy.Len()).To(Equal(len + 1))
			})
		})
		When("add other", func() {
			BeforeEach(func() {
				ok = copy.AddItem(&TestSpec{})
			})
			It("can not add", func() {
				Expect(ok).To((BeFalse()))
				Expect(copy.Len()).To(Equal(len))
			})
		})
	})
}
