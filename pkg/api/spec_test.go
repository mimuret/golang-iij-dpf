package api_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("spec.go", func() {
	Describe("DeepCopySpec", func() {
		var (
			spec *testtool.TestSpec
			copy api.Spec
		)
		When("param is nil", func() {
			BeforeEach(func() {
				spec = nil
				copy = api.DeepCopySpec(spec)
			})
			It("retruns nil", func() {
				Expect(copy).To(BeNil())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				spec = &testtool.TestSpec{}
				copy = api.DeepCopySpec(spec)
			})
			It("returns copy", func() {
				Expect(copy).NotTo(BeNil())
			})
		})
	})
	Describe("DeepCopyListSpec", func() {
		var (
			spec *testtool.TestSpecList
			copy api.Spec
		)
		When("param is nil", func() {
			BeforeEach(func() {
				spec = nil
				copy = api.DeepCopyListSpec(spec)
			})
			It("retruns nil", func() {
				Expect(copy).To(BeNil())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				spec = &testtool.TestSpecList{}
				copy = api.DeepCopyListSpec(spec)
			})
			It("returns copy", func() {
				Expect(copy).NotTo(BeNil())
			})
		})
	})
	Describe("DeepCopyCountableListSpec", func() {
		var (
			spec *testtool.TestSpecCountableList
			copy api.Spec
		)
		When("param is nil", func() {
			BeforeEach(func() {
				spec = nil
				copy = api.DeepCopyCountableListSpec(spec)
			})
			It("retruns nil", func() {
				Expect(copy).To(BeNil())
			})
		})
		When("param is not nil", func() {
			BeforeEach(func() {
				spec = &testtool.TestSpecCountableList{}
				copy = api.DeepCopyCountableListSpec(spec)
			})
			It("returns copy", func() {
				Expect(copy).NotTo(BeNil())
			})
		})
	})
})
