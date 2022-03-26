package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("contract_zones", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.Zone
		slist  contracts.ContractZoneList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Zone{
			Id:               "m1",
			CommonConfigID:   1,
			ServiceCode:      "dpm0000001",
			State:            types.StateBeforeStart,
			Favorite:         types.FavoriteHighPriority,
			Name:             "example.jp.",
			Network:          "",
			Description:      "zone 1",
			ZoneProxyEnabled: types.Disabled,
		}
		s2 = core.Zone{
			Id:               "m2",
			CommonConfigID:   2,
			ServiceCode:      "dpm0000002",
			State:            types.StateRunning,
			Favorite:         types.FavoriteLowPriority,
			Name:             "168.192.in-addr.arpa.",
			Network:          "192.168.0.0/16",
			Description:      "zone 2",
			ZoneProxyEnabled: types.Enabled,
		}

		slist = contracts.ContractZoneList{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			Items: []core.Zone{s1, s2},
		}
	})
	Describe("ContractZoneList", func() {
		var c contracts.ContractZoneList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/zones", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "E8FCA53D138A4AAD8DCED5514A5AF7CB",
					"results": [
						{
							"id": "m1",
							"common_config_id": 1,
							"service_code": "dpm0000001",
							"state": 1,
							"favorite": 1,
							"name": "example.jp.",
							"network": "",
							"description": "zone 1",
							"zone_proxy_enabled": 0
						},
						{
							"id": "m2",
							"common_config_id": 2,
							"service_code": "dpm0000002",
							"state": 2,
							"favorite":2,
							"name": "168.192.in-addr.arpa.",
							"network": "192.168.0.0/16",
							"description": "zone 2",
							"zone_proxy_enabled": 1
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.ContractZoneList{
						AttributeMeta: contracts.AttributeMeta{
							ContractID: "f1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("E8FCA53D138A4AAD8DCED5514A5AF7CB"))
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
				It("can set ContractID", func() {
					Expect(slist.GetContractID()).To(Equal("id12"))
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
			When("arguments type missmatch (ContractID)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.ContractZoneList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "zones")
			testtool.TestGetPathMethodForCountableList(&slist, "/contracts/f1/zones")
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
	})
})
