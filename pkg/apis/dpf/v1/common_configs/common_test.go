package common_configs_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/common_configs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("common_configs", func() {
	Context("AttributeMeta", func() {
		var meta common_configs.AttributeMeta
		BeforeEach(func() {
			meta = common_configs.AttributeMeta{}
		})
		Context("GetGroup", func() {
			It("returns name", func() {
				Expect(meta.GetGroup()).To(Equal("common-configs.api.dns-platform.jp/v1"))
			})
		})
		Context("GetCommonConfigId", func() {
			When("default", func() {
				It("returns 0", func() {
					Expect(meta.GetCommonConfigId()).To(Equal(int64(0)))
				})
			})
			When("default", func() {
				BeforeEach(func() {
					meta.CommonConfigId = 1
				})
				It("returns CommonConfigId", func() {
					Expect(meta.GetCommonConfigId()).To(Equal(int64(1)))
				})
			})
		})
		Context("SetCommonConfigId", func() {
			BeforeEach(func() {
				meta.SetCommonConfigId(2)
			})
			It("can set CommonConfigId", func() {
				Expect(meta.GetCommonConfigId()).To(Equal(int64(2)))
			})
		})
		Context("DeepCopy", func() {
			var (
				copy    *common_configs.AttributeMeta
				nilMeta *common_configs.AttributeMeta
			)
			When("AttributeMeta is not nil", func() {
				BeforeEach(func() {
					copy = meta.DeepCopy()
				})
				It("returns copy ", func() {
					Expect(copy).To(Equal(&meta))
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
