package lb_domains_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("lb_domains", func() {
	Context("AttributeMeta", func() {
		var meta lb_domains.AttributeMeta
		BeforeEach(func() {
			meta = lb_domains.AttributeMeta{}
		})
		Context("GetGroup", func() {
			It("returns lb_domains", func() {
				Expect(meta.GetGroup()).To(Equal("lb-domains.api.dns-platform.jp/v1"))
			})
		})
		Context("GetLBDoaminID", func() {
			When("default", func() {
				It("returns ", func() {
					Expect(meta.GetLBDoaminID()).To(Equal(""))
				})
			})
			When("default", func() {
				BeforeEach(func() {
					meta.LBDomainID = "id1"
				})
				It("returns ContractID", func() {
					Expect(meta.GetLBDoaminID()).To(Equal("id1"))
				})
			})
			Context("SetLBDoaminID", func() {
				BeforeEach(func() {
					meta.SetLBDoaminID("id2")
				})
				It("can set ContractID", func() {
					Expect(meta.GetLBDoaminID()).To(Equal("id2"))
				})
			})
			Context("DeepCopy", func() {
				var (
					copy    *lb_domains.AttributeMeta
					nilMeta *lb_domains.AttributeMeta
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
