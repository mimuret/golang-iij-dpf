package zones_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("zones", func() {
	Context("AttributeMeta", func() {
		var meta zones.AttributeMeta
		BeforeEach(func() {
			meta = zones.AttributeMeta{}
		})
		Context("GetGroup", func() {
			It("returns zones", func() {
				Expect(meta.GetGroup()).To(Equal("zones.api.dns-platform.jp/v1"))
			})
		})
		Context("GetZoneId", func() {
			When("default", func() {
				It("returns ", func() {
					Expect(meta.GetZoneId()).To(Equal(""))
				})
			})
			When("default", func() {
				BeforeEach(func() {
					meta.ZoneId = "id1"
				})
				It("returns ZoneId", func() {
					Expect(meta.GetZoneId()).To(Equal("id1"))
				})
			})
			Context("SetZoneId", func() {
				BeforeEach(func() {
					meta.SetZoneId("id2")
				})
				It("can set ZoneId", func() {
					Expect(meta.GetZoneId()).To(Equal("id2"))
				})
			})
			Context("DeepCopy", func() {
				var (
					copy    *zones.AttributeMeta
					nilMeta *zones.AttributeMeta
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
})
