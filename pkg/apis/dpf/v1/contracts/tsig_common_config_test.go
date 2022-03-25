package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("tsigs/common_configs", func() {
	var (
		c      contracts.TsigCommonConfigList
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.CommonConfig
		slist  contracts.TsigCommonConfigList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.CommonConfig{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                1,
			Name:              "共通設定1",
			ManagedDNSEnabled: types.Enabled,
			Default:           types.Enabled,
			Description:       "common config 1",
		}
		s2 = contracts.CommonConfig{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                2,
			Name:              "共通設定2",
			ManagedDNSEnabled: types.Disabled,
			Default:           types.Disabled,
			Description:       "common config 2",
		}
		slist = contracts.TsigCommonConfigList{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Items: []contracts.CommonConfig{s1, s2},
			Id:    1,
		}
	})
	Describe("TsigCommonConfigList", func() {
		Context("List", func() {
			BeforeEach(func() {

				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1/common_configs", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "C38FBA61C657432695BC798DD96B0D9A",
					"results": [
						{
							"id": 1,
							"name": "共通設定1",
							"managed_dns_enabled": 1,
							"default": 1,
							"description": "common config 1"
						},
						{
							"id": 2,
							"name": "共通設定2",
							"managed_dns_enabled": 0,
							"default": 0,
							"description": "common config 2"
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1/common_configs/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "1073C4F38C0A4027B4776F0F3E8197BB",
					"result": {
						"count": 2
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.TsigCommonConfigList{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
						Id: 1,
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("C38FBA61C657432695BC798DD96B0D9A"))
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
					err = slist.SetPathParams("id12", 122)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(slist.GetContractId()).To(Equal("id12"))
				})
				It("can set Id (tsig_id)", func() {
					Expect(slist.Id).To(Equal(int64(122)))
				})
			})
			When("not enough arguments", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("id12")
				})
				It("not returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("f1", 2, 3)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2, 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (Id)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("F1", "f2")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.TsigCommonConfigList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "tsigs/1/common_configs")

			testtool.TestGetPathMethodForCountableList(&slist, "/contracts/f1/tsigs/1/common_configs")
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
		Context("api.CountableListSpec common test", func() {
			testtool.TestGetMaxLimit(&slist, 10000)
			testtool.TestClearItems(&slist)
			testtool.TestAddItem(&slist, s1)
		})
		Context("contracts.ChildSpec common test", func() {
			Context("GetId", func() {
				It("returns Id", func() {
					Expect(slist.GetId()).To(Equal(slist.Id))
				})
			})
			Context("SetId", func() {
				BeforeEach(func() {
					slist.SetId(2)
				})
				It("can set Id", func() {
					Expect(slist.GetId()).To(Equal(int64(2)))
				})
			})
		})
	})
})
