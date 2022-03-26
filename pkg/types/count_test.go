package types_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("types.count", func() {
	Context("Count", func() {
		var count types.Count
		BeforeEach(func() {
			count.Count = 100
		})
		Context("GetCount", func() {
			It("returns count value", func() {
				Expect(count.GetCount()).To(Equal(int32(100)))
			})
		})
		Context("SetCount", func() {
			It("returns count value", func() {
				count.SetCount(50)
				Expect(count.GetCount()).To(Equal(int32(50)))
			})
		})
		Context("DeepCopy", func() {
			var (
				copy    *types.Count
				nilMeta *types.Count
			)
			When("AttributeMeta is not nil", func() {
				BeforeEach(func() {
					copy = count.DeepCopy()
				})
				It("returns copy ", func() {
					Expect(copy).To(Equal(&count))
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
})
