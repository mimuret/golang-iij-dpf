package contracts_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("contracts", func() {
	Context("AttributeMeta", func() {
		var (
			meta contracts.AttributeMeta
		)
		BeforeEach(func() {
			meta = contracts.AttributeMeta{}
		})
		Context("GetGroup", func() {
			It("returns contracts", func() {
				Expect(meta.GetGroup()).To(Equal("contracts"))
			})
		})
		Context("GetContractId", func() {
			When("default", func() {
				It("returns ", func() {
					Expect(meta.GetContractId()).To(Equal(""))
				})
			})
			When("default", func() {
				BeforeEach(func() {
					meta.ContractId = "id1"
				})
				It("returns ContractId", func() {
					Expect(meta.GetContractId()).To(Equal("id1"))
				})
			})
			Context("SetContractId", func() {
				BeforeEach(func() {
					meta.SetContractId("id2")
				})
				It("can set ContractId", func() {
					Expect(meta.GetContractId()).To(Equal("id2"))
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
