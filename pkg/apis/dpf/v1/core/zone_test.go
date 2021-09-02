package core_test

import (
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("zones", func() {
	var (
		c      core.Zone
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.Zone
		slist  core.ZoneList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Zone{
			Id:               "m1",
			CommonConfigId:   1,
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
			CommonConfigId:   2,
			ServiceCode:      "dpm0000002",
			State:            types.StateRunning,
			Favorite:         types.FavoriteLowPriority,
			Name:             "168.192.in-addr.arpa.",
			Network:          "192.168.0.0/16",
			Description:      "zone 2",
			ZoneProxyEnabled: types.Enabled,
		}
		slist = core.ZoneList{
			Items: []core.Zone{s1, s2},
		}
	})
	Describe("Contract", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
							"id": "m1",
							"common_config_id": 1,
							"service_code": "dpm0000001",
							"state": 1,
							"favorite": 1,
							"name": "example.jp.",
							"network": "",
							"description": "zone 1",
							"zone_proxy_enabled": 0
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
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
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns Zone m1", func() {
				BeforeEach(func() {
					c = core.Zone{
						Id: "m1",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("76D7462F50E3417C96898A827E0D1EC1"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns Zone m2", func() {
				BeforeEach(func() {
					c = core.Zone{
						Id: "m2",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("B10F0042282A46C18BFB440B5B58BF8C"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Update(&s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/zones/m1"]).To(MatchJSON(`{
							"favorite": 1,
							"description": "zone 1"
						}`))
			})
		})
		Context("SetPathParams", func() {
			When("no arguments, nothing to do", func() {
				BeforeEach(func() {
					err = s1.SetPathParams()
				})
				It("returns error", func() {
					Expect(err).To(Succeed())
				})
			})
			When("enough arguments", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m100")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.Id).To(Equal("m100"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m100", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.Zone
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "zones")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("ZoneList", func() {
		var (
			c core.ZoneList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
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
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "9C518C729E5541D999389C686FE8987D",
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
					c = core.ZoneList{}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("2C18922432DC48D485613F5383A7ED8E"))
					Expect(c).To(Equal(slist))
				})
			})
		})
		Context("SetPathParams", func() {
			BeforeEach(func() {
				err = slist.SetPathParams()
			})
			It("nothing todo", func() {
				Expect(err).To(Succeed())
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.ZoneList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "zones")

			testtool.TestGetPathMethodForCountableList(&slist, "/zones")
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
	Describe("ZoneListSearchKeywords", func() {
		Context("GetValues", func() {
			var (
				testcase = []struct {
					keyword core.ZoneListSearchKeywords
					values  url.Values
				}{
					{
						core.ZoneListSearchKeywords{
							CommonSearchParams: api.CommonSearchParams{
								Type:   api.SearchTypeAND,
								Offset: int32(10),
								Limit:  int32(100),
							},
						},
						url.Values{
							"type":   []string{"AND"},
							"offset": []string{"10"},
							"limit":  []string{"100"},
						},
					},
					{
						core.ZoneListSearchKeywords{
							CommonSearchParams: api.CommonSearchParams{
								Type:   api.SearchTypeOR,
								Offset: int32(10),
								Limit:  int32(100),
							},
							FullText:         api.KeywordsString{"hogehoge", "🐰"},
							ServiceCode:      api.KeywordsString{"mxxxxxx", "mxxxxxx1"},
							Name:             api.KeywordsString{"example.jp", "example.net"},
							Network:          api.KeywordsString{"192.168.0.0/24"},
							State:            api.KeywordsState{types.StateBeforeStart, types.StateRunning},
							Favorite:         api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
							Description:      api.KeywordsString{"あああ", "🍺"},
							CommonConfigId:   api.KeywordsId{100, 200},
							ZoneProxyEnabled: api.KeywordsBoolean{types.Enabled, types.Disabled},
						},
						/*
							_keywords_full_text[]
							_keywords_service_code[]
							_keywords_name[]
							_keywords_network[]
							_keywords_state[]
							_keywords_favorite[]
							_keywords_description[]
							_keywords_common_config_id[]
							_keywords_zone_proxy_enabled[]
						*/
						url.Values{
							"type":                           []string{"OR"},
							"offset":                         []string{"10"},
							"limit":                          []string{"100"},
							"_keywords_full_text[]":          []string{"hogehoge", "🐰"},
							"_keywords_service_code[]":       []string{"mxxxxxx", "mxxxxxx1"},
							"_keywords_name[]":               []string{"example.jp", "example.net"},
							"_keywords_network[]":            []string{"192.168.0.0/24"},
							"_keywords_state[]":              []string{"1", "2"},
							"_keywords_favorite[]":           []string{"1", "2"},
							"_keywords_description[]":        []string{"あああ", "🍺"},
							"_keywords_common_config_id[]":   []string{"100", "200"},
							"_keywords_zone_proxy_enabled[]": []string{"1", "0"},
						},
					},
				}
			)
			It("can convert url.Value", func() {
				for _, tc := range testcase {
					s, err := tc.keyword.GetValues()
					Expect(err).To(Succeed())
					Expect(s).To(Equal(tc.values))
				}
			})
		})
	})
})
