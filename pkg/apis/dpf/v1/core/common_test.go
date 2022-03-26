package core_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AttributeMeta", func() {
	a := core.AttributeMeta{}
	Context("GetGroup", func() {
		It("returns core", func() {
			Expect(a.GetGroup()).To(Equal("core.api.dns-platform.jp/v1"))
		})
	})
	Context("DeepCopy", func() {
		var (
			copy    *core.AttributeMeta
			nilMeta *core.AttributeMeta
		)
		When("AttributeMeta is not nil", func() {
			BeforeEach(func() {
				copy = a.DeepCopy()
			})
			It("returns copy ", func() {
				Expect(copy).To(Equal(&a))
			})
		})
		When("AttributeMeta is nil", func() {
			BeforeEach(func() {
				copy = nilMeta.DeepCopy()
			})
			It("returns copy ", func() {
				Expect(copy).To(BeNil())
			})
		})
	})
})
