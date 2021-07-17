package types_test

import (
	"testing"

	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "types Suite")
}

var _ = Describe("types.count", func() {
	Context("Count", func() {
		var (
			count types.Count
		)
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
	})
})
