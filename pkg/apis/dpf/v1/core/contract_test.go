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

var _ = Describe("contract", func() {
	var (
		c      core.Contract
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.Contract
		slist  core.ContractList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Contract{
			ID:          "f1",
			ServiceCode: "dpf0000001",
			State:       types.StateBeforeStart,
			Favorite:    types.FavoriteHighPriority,
			Plan:        core.PlanBasic,
			Description: "contract 1",
		}
		s2 = core.Contract{
			ID:          "f2",
			ServiceCode: "dpf0000002",
			State:       types.StateRunning,
			Favorite:    types.FavoriteLowPriority,
			Plan:        core.PlanPremium,
			Description: "contract 2",
		}
		slist = core.ContractList{
			Items: []core.Contract{s1, s2},
		}
	})
	Describe("Contract", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "D7155DA183ED47C48437183873DAC765",
					"result": {
						"id": "f1",
						"service_code": "dpf0000001",
						"state": 1,
						"favorite": 1,
						"plan": 1,
						"description": "contract 1"
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EECF5BC0FF834BFFA4F0253224667D10",
					"result": {
						"id": "f2",
						"service_code": "dpf0000002",
						"state": 2,
						"favorite": 2,
						"plan": 2,
						"description": "contract 2"
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns Contract f1", func() {
				BeforeEach(func() {
					c = core.Contract{
						ID: "f1",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("D7155DA183ED47C48437183873DAC765"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns Contract f2", func() {
				BeforeEach(func() {
					c = core.Contract{
						ID: "f2",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("EECF5BC0FF834BFFA4F0253224667D10"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/contracts/f1"]).To(MatchJSON(`{
							"favorite": 1,
							"description": "contract 1"
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
					err = s1.SetPathParams("f10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractID", func() {
					Expect(s1.ID).To(Equal("f10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("f10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractID)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.Contract
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "contracts")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/contracts/f1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/contracts/f1")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("CommonConfigList", func() {
		var c core.ContractList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "1773D7B5F415464C944837053CE8BBF2",
					"results": [
						{
							"id": "f1",
							"service_code": "dpf0000001",
							"state": 1,
							"favorite": 1,
							"plan": 1,
							"description": "contract 1"
							},
						{
							"id": "f2",
							"service_code": "dpf0000002",
							"state": 2,
							"favorite": 2,
							"plan": 2,
							"description": "contract 2"
							}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "BD88B5E38397477989002302583C7FCC",
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
					c = core.ContractList{}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("1773D7B5F415464C944837053CE8BBF2"))
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
			var nilSpec *core.ContractList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "contracts")

			testtool.TestGetPathMethodForCountableList(&slist, "/contracts")
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
	Describe("ContractListSearchKeywords", func() {
		Context("GetValues", func() {
			testcase := []struct {
				keyword core.ContractListSearchKeywords
				values  url.Values
			}{
				{
					core.ContractListSearchKeywords{
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
					core.ContractListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeOR,
							Offset: int32(10),
							Limit:  int32(100),
						},
						FullText:    api.KeywordsString{"hogehoge", "🐰"},
						ServiceCode: api.KeywordsString{"fxxxxxx", "fxxxxxx1"},
						Plan:        core.KeywordsPlan{core.PlanBasic, core.PlanPremium},
						State:       api.KeywordsState{types.StateBeforeStart, types.StateRunning},
						Favorite:    api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
						Description: api.KeywordsString{"あああ", "🍺"},
					},
					url.Values{
						"type":                     []string{"OR"},
						"offset":                   []string{"10"},
						"limit":                    []string{"100"},
						"_keywords_full_text[]":    []string{"hogehoge", "🐰"},
						"_keywords_service_code[]": []string{"fxxxxxx", "fxxxxxx1"},
						"_keywords_plan[]":         []string{"1", "2"},
						"_keywords_state[]":        []string{"1", "2"},
						"_keywords_favorite[]":     []string{"1", "2"},
						"_keywords_description[]":  []string{"あああ", "🍺"},
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
	Context("PlanToString", func() {
		Context("String", func() {
			It("returns plan string", func() {
				Expect(core.PlanBasic.String()).To(Equal("basic"))
				Expect(core.PlanPremium.String()).To(Equal("premium"))
			})
		})
	})
})
