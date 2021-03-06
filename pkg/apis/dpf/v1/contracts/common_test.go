package contracts_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("contracts", func() {
	Context("AttributeMeta", func() {
		var meta contracts.AttributeMeta
		BeforeEach(func() {
			meta = contracts.AttributeMeta{}
		})
		Context("GetGroup", func() {
			It("returns contracts", func() {
				Expect(meta.GetGroup()).To(Equal("contracts.api.dns-platform.jp/v1"))
			})
		})
		Context("GetContractID", func() {
			When("default", func() {
				It("returns ", func() {
					Expect(meta.GetContractID()).To(Equal(""))
				})
			})
			When("default", func() {
				BeforeEach(func() {
					meta.ContractID = "id1"
				})
				It("returns ContractID", func() {
					Expect(meta.GetContractID()).To(Equal("id1"))
				})
			})
			Context("SetContractID", func() {
				BeforeEach(func() {
					meta.SetContractID("id2")
				})
				It("can set ContractID", func() {
					Expect(meta.GetContractID()).To(Equal("id2"))
				})
			})
			Context("DeepCopy", func() {
				var (
					copy    *contracts.AttributeMeta
					nilMeta *contracts.AttributeMeta
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
