package utils_test

import (
	"testing"

	"github.com/mimuret/golang-iij-dpf/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "utils Suite")
}

var _ = Describe("utils", func() {
	Context("ToInt64", func() {
		When("convertable types", func() {
			When("int", func() {
				It("return int64", func() {
					Expect(utils.ToInt64(int(10))).To(Equal(int64(10)))
				})
			})
			When("int8", func() {
				It("return int64", func() {
					Expect(utils.ToInt64(int8(10))).To(Equal(int64(10)))
				})
			})
			When("int16", func() {
				It("return int64", func() {
					Expect(utils.ToInt64(int16(10))).To(Equal(int64(10)))
				})
			})
			When("int32", func() {
				It("return int64", func() {
					Expect(utils.ToInt64(int32(10))).To(Equal(int64(10)))
				})
			})
			When("int64", func() {
				It("return int64", func() {
					Expect(utils.ToInt64(int64(10))).To(Equal(int64(10)))
				})
			})
		})
		When("not convertable types", func() {
			When("uint", func() {
				It("return err", func() {
					_, err := utils.ToInt64(uint(10))
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
	Context("ToString", func() {
		When("convertable types", func() {
			When("string", func() {
				It("return string", func() {
					Expect(utils.ToString("hogehoge")).To(Equal("hogehoge"))
				})
			})
		})
		When("not convertable types", func() {
			When("int", func() {
				It("return error", func() {
					_, err := utils.ToString(uint(10))
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
