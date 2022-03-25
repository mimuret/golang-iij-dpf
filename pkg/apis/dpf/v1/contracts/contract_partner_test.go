package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("contract_partners", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.ContractPartner
		slist  contracts.ContractPartnerList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.ContractPartner{"svc001"}
		s2 = contracts.ContractPartner{"svc002"}

		slist = contracts.ContractPartnerList{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Items: []contracts.ContractPartner{s1, s2},
		}
	})
	Describe("ContractPartnerList", func() {
		var (
			c contracts.ContractPartnerList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/contract_partners", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0EF26E92E5D54352A78C9EA4551BE502",
					"results": [
						{
							"service_code": "svc001"
						},
						{
							"service_code": "svc002"
						}
					]
				}`)))

			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.ContractPartnerList{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0EF26E92E5D54352A78C9EA4551BE502"))
					Expect(c).To(Equal(slist))
				})
			})
		})
		Context("SetPathParams", func() {
			When("no arguments, nothing to do", func() {
				BeforeEach(func() {
					err = slist.SetPathParams()
				})
				It("returns error", func() {
					Expect(err).To(Succeed())
				})
			})
			When("enough arguments", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("id12")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(slist.GetContractId()).To(Equal("id12"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("f1", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.ContractPartnerList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "contract_partners")
			testtool.TestGetPathMethodForList(&slist, "/contracts/f1/contract_partners")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 2)
			Context("Index", func() {
				When("index exist", func() {
					It("returns item", func() {
						Expect(slist.Index(0)).To(Equal(s1))
					})
				})
			})
		})
	})
	Describe("ContractPartner", func() {
		var (
			s, copy, nilSpec *contracts.ContractPartner
		)
		BeforeEach(func() {
			s = &contracts.ContractPartner{}
		})
		Context("DeepCopyObject", func() {
			When("object is not nil", func() {
				BeforeEach(func() {
					copy = s.DeepCopy()
				})
				It("returns copy", func() {
					Expect(copy).NotTo(BeNil())
					Expect(copy).To(Equal(s))
				})
			})
			When("object is not nil", func() {
				BeforeEach(func() {
					copy = nilSpec.DeepCopy()
				})
				It("returns nil", func() {
					Expect(copy).To(BeNil())
				})
			})
		})
	})
})
