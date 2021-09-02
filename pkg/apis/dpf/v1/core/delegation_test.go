package core_test

import (
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("delegations", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.Delegation
		slist  core.DelegationList
		atTime types.Time
	)
	BeforeEach(func() {
		atTime, err = types.ParseTime(time.RFC3339Nano, "2021-06-20T13:21:12.881Z")
		Expect(err).To(Succeed())
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Delegation{
			Id:                    "m1",
			ServiceCode:           "dpm000001",
			Name:                  "example.jp.",
			Network:               "",
			Description:           "zone 1",
			DelegationRequestedAt: atTime,
		}
		s2 = core.Delegation{
			Id:                    "m2",
			ServiceCode:           "dpm000002",
			Name:                  "168.192.in-addr.arpa.",
			Network:               "192.168.0.0/16",
			Description:           "zone 2",
			DelegationRequestedAt: types.Time{},
		}
		slist = core.DelegationList{
			Items: []core.Delegation{s1, s2},
		}
	})
	Describe("Delegation", func() {
		Context("DeepCopyObject", func() {
			var nilSpec *core.Delegation
			testtool.TestDeepCopyObject(&s1, nilSpec)
		})
	})
	Describe("DelegationList", func() {
		var (
			c core.DelegationList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/delegations", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "51367A80A2B447EFA490AF379F12DD48",
					"results": [
						{
							"id": "m1",
							"service_code": "dpm000001",
							"name": "example.jp.",
							"network": "",
							"description": "zone 1",
							"delegation_requested_at": "2021-06-20T13:21:12.881Z"
						},
						{
							"id": "m2",
							"service_code": "dpm000002",
							"name": "168.192.in-addr.arpa.",
							"network": "192.168.0.0/16",
							"description": "zone 2",
							"delegation_requested_at": ""
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/delegations/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "D371175B1BB248BE812A845C351FDE82",
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
					c = core.DelegationList{}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("51367A80A2B447EFA490AF379F12DD48"))
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
			var nilSpec *core.DelegationList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "delegations")

			testtool.TestGetPathMethodForCountableList(&slist, "/delegations")
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
	Describe("DelegationListSearchKeywords", func() {
		Context("GetValues", func() {
			var (
				testcase = []struct {
					keyword core.DelegationListSearchKeywords
					values  url.Values
				}{
					{
						core.DelegationListSearchKeywords{
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
						core.DelegationListSearchKeywords{
							CommonSearchParams: api.CommonSearchParams{
								Type:   api.SearchTypeOR,
								Offset: int32(10),
								Limit:  int32(100),
							},
							FullText:    api.KeywordsString{"hogehoge", "🐰"},
							ServiceCode: api.KeywordsString{"mxxxxxx", "mxxxxxx1"},
							Name:        api.KeywordsString{"example.jp", "example.net"},
							Network:     api.KeywordsString{"192.168.0.0/24"},
							Favorite:    api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
							Description: api.KeywordsString{"あああ", "🍺"},
							Requested:   api.KeywordsBoolean{types.Disabled, types.Enabled},
						},
						url.Values{
							"type":                     []string{"OR"},
							"offset":                   []string{"10"},
							"limit":                    []string{"100"},
							"_keywords_full_text[]":    []string{"hogehoge", "🐰"},
							"_keywords_service_code[]": []string{"mxxxxxx", "mxxxxxx1"},
							"_keywords_name[]":         []string{"example.jp", "example.net"},
							"_keywords_network[]":      []string{"192.168.0.0/24"},
							"_keywords_favorite[]":     []string{"1", "2"},
							"_keywords_description[]":  []string{"あああ", "🍺"},
							"_keywords_requested[]":    []string{"0", "1"},
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
