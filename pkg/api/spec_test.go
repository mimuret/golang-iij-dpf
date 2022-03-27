package api_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type countableListErr testtool.TestSpecCountableList

func (c *countableListErr) DeepCopyObject() api.Object {
	return &dummyObject{}
}

type dummyObject struct{}

func (c *dummyObject) DeepCopyObject() api.Object {
	return &dummyObject{}
}

var _ = Describe("spec.go", func() {
	Describe("DeepCopySpec", func() {
		var (
			spec *testtool.TestSpec
			ret  api.Spec
		)
		When("param is nil value", func() {
			BeforeEach(func() {
				ret = api.DeepCopySpec(nil)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("param is nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopySpec(spec)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("DeepCopyObject returns invalid", func() {
			It("raise panic", func() {
				Expect(func() { _ = api.DeepCopySpec(&countableListErr{}) }).To(Panic())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopySpec(&testtool.TestSpec{})
			})
			It("returns ret", func() {
				Expect(ret).NotTo(BeNil())
			})
		})
	})
	Describe("DeepCopyListSpec", func() {
		var (
			spec *testtool.TestSpecList
			ret  api.Spec
		)
		When("param is nil svalue", func() {
			BeforeEach(func() {
				ret = api.DeepCopyListSpec(nil)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("param is nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopyListSpec(spec)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("DeepCopyObject returns invalid", func() {
			It("raise panic", func() {
				Expect(func() { _ = api.DeepCopyListSpec(&countableListErr{}) }).To(Panic())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopyListSpec(&testtool.TestSpecList{})
			})
			It("returns ret", func() {
				Expect(ret).NotTo(BeNil())
			})
		})
	})
	Describe("DeepCopyCountableListSpec", func() {
		var (
			spec *testtool.TestSpecCountableList
			ret  api.Spec
		)
		When("param is nil value", func() {
			BeforeEach(func() {
				ret = api.DeepCopyCountableListSpec(nil)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("param is nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopyCountableListSpec(spec)
			})
			It("retruns nil", func() {
				Expect(ret).To(BeNil())
			})
		})
		When("DeepCopyObject returns invalid", func() {
			It("raise panic", func() {
				Expect(func() { _ = api.DeepCopyCountableListSpec(&countableListErr{}) }).To(Panic())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				ret = api.DeepCopyCountableListSpec(&testtool.TestSpecCountableList{})
			})
			It("returns ret", func() {
				Expect(ret).NotTo(BeNil())
			})
		})
	})
})
