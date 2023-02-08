package core_test

import (
	"context"
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

var _ = Describe("lb_domains", func() {
	var (
		c      core.LBDomain
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.LBDomain
		slist  core.LBDomainList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.LBDomain{
			ID:               "b0000000000001",
			CommonConfigID:   1,
			ServiceCode:      "dpl0000001",
			State:            types.StateBeforeStart,
			Favorite:         types.FavoriteHighPriority,
			Name:             "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
			Description:      "lb_domains 1",
			RuleResourceName: "rule-1",
		}
		s2 = core.LBDomain{
			ID:               "b0000000000002",
			CommonConfigID:   2,
			ServiceCode:      "dpl0000002",
			State:            types.StateRunning,
			Favorite:         types.FavoriteLowPriority,
			Name:             "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
			Description:      "lb_domains 2",
			RuleResourceName: "rule-2",
		}
		slist = core.LBDomainList{
			Items: []core.LBDomain{s1, s2},
		}
	})
	Describe("Contract", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
							"id": "b0000000000001",
							"common_config_id": 1,
							"service_code": "dpl0000001",
							"state": 1,
							"favorite": 1,
							"name": "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
							"description": "lb_domains 1",
							"rule_resource_name": "rule-1"
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000002", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
						"id": "b0000000000002",
						"common_config_id": 2,
						"service_code": "dpl0000002",
						"state": 2,
						"favorite":2,
						"name": "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
						"description": "lb_domains 2",
						"rule_resource_name": "rule-2"
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns LBDomain m1", func() {
				BeforeEach(func() {
					c = core.LBDomain{
						ID: "b0000000000001",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("76D7462F50E3417C96898A827E0D1EC1"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns LBDomain m2", func() {
				BeforeEach(func() {
					c = core.LBDomain{
						ID: "b0000000000002",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("B10F0042282A46C18BFB440B5B58BF8C"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/lb_domains/b0000000000001", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Update(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/lb_domains/b0000000000001"]).To(MatchJSON(`{
							"favorite": 1,
							"description": "lb_domains 1",
							"rule_resource_name": "rule-1"
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
					err = s1.SetPathParams("b0000000000010")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractID", func() {
					Expect(s1.ID).To(Equal("b0000000000010"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.LBDomain
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "lb_domains")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/lb_domains/b0000000000001")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/lb_domains/b0000000000001")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("LBDomainList", func() {
		var c core.LBDomainList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
					"results": [
						{
							"id": "b0000000000001",
							"common_config_id": 1,
							"service_code": "dpl0000001",
							"state": 1,
							"favorite": 1,
							"name": "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
							"description": "lb_domains 1",
							"rule_resource_name": "rule-1"
						},
						{
							"id": "b0000000000002",
							"common_config_id": 2,
							"service_code": "dpl0000002",
							"state": 2,
							"favorite":2,
							"name": "lb.xxxxxxxxxxxxxxxx.x.d-16.jp.",
							"description": "lb_domains 2",
							"rule_resource_name": "rule-2"
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/count", httpmock.NewBytesResponder(200, []byte(`{
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
					c = core.LBDomainList{}
					reqId, err = cl.List(context.Background(), &c, nil)
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
			var nilSpec *core.LBDomainList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "lb_domains")

			testtool.TestGetPathMethodForCountableList(&slist, "/lb_domains")
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
	Describe("LBDomainListSearchKeywords", func() {
		Context("GetValues", func() {
			testcase := []struct {
				keyword core.LBDomainListSearchKeywords
				values  url.Values
			}{
				{
					core.LBDomainListSearchKeywords{
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
					core.LBDomainListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeOR,
							Offset: int32(10),
							Limit:  int32(100),
						},
						FullText:       api.KeywordsString{"hogehoge", "🐰"},
						ServiceCode:    api.KeywordsString{"mxxxxxx", "mxxxxxx1"},
						Name:           api.KeywordsString{"example.jp", "example.net"},
						State:          api.KeywordsState{types.StateBeforeStart, types.StateRunning},
						Favorite:       api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
						Description:    api.KeywordsString{"あああ", "🍺"},
						CommonConfigID: api.KeywordsID{100, 200},
						Label:          api.KeywordsLabels{api.Label{Key: "hogehoge", Value: "hugahuga"}},
					},
					/*
						_keywords_full_text[]
						_keywords_service_code[]
						_keywords_name[]
						_keywords_state[]
						_keywords_favorite[]
						_keywords_description[]
						_keywords_common_config_id[]
						_keywords_label[]
					*/
					url.Values{
						"type":                         []string{"OR"},
						"offset":                       []string{"10"},
						"limit":                        []string{"100"},
						"_keywords_full_text[]":        []string{"hogehoge", "🐰"},
						"_keywords_service_code[]":     []string{"mxxxxxx", "mxxxxxx1"},
						"_keywords_name[]":             []string{"example.jp", "example.net"},
						"_keywords_state[]":            []string{"1", "2"},
						"_keywords_favorite[]":         []string{"1", "2"},
						"_keywords_description[]":      []string{"あああ", "🍺"},
						"_keywords_common_config_id[]": []string{"100", "200"},
						"_keywords_label[]":            []string{"hogehoge=hugahuga"},
					},
				},
			}
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
