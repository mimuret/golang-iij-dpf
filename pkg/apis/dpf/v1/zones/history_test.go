package zones_test

import (
	"context"
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("history text", func() {
	var (
		c      zones.HistoryList
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.History
		slist  zones.HistoryList
	)
	BeforeEach(func() {
		atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:50:37.396Z")
		Expect(err).To(Succeed())

		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.History{
			Id:          1,
			CommittedAt: atTime,
			Description: "commit 1",
			Operator:    "user1",
		}
		s2 = zones.History{
			Id:          2,
			CommittedAt: atTime,
			Description: "commit 2",
			Operator:    "user2",
		}
		slist = zones.HistoryList{
			AttributeMeta: zones.AttributeMeta{
				ZoneID: "m1",
			},
			Items: []zones.History{s1, s2},
		}
	})
	Describe("HistoryList", func() {
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "71B4D2B268744DBB8C3220C3A38C29E5",
					"results": [
						{
							"id": 1,
							"committed_at": "2021-06-20T10:50:37.396Z",
							"description": "commit 1",
							"operator": "user1"
							},
						{
							"id": 2,
							"committed_at": "2021-06-20T10:50:37.396Z",
							"description": "commit 2",
							"operator": "user2"
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "8BAB2071CF2C4509A93E485D7F368A71",
					"result": {
						"count": 2
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns m1", func() {
				BeforeEach(func() {
					c = zones.HistoryList{
						AttributeMeta: zones.AttributeMeta{
							ZoneID: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("71B4D2B268744DBB8C3220C3A38C29E5"))
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
					err = slist.SetPathParams("m10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractID", func() {
					Expect(slist.GetZoneID()).To(Equal("m10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("m10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneID)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.HistoryList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "zone_histories")

			testtool.TestGetPathMethodForCountableList(&slist, "/zones/m1/zone_histories")
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
			testtool.TestGetMaxLimit(&slist, 100)
			testtool.TestClearItems(&slist)
			testtool.TestAddItem(&slist, s1)
		})
	})
	Context("HistoryListSearchKeywords", func() {
		testcase := []struct {
			keyword zones.HistoryListSearchKeywords
			values  url.Values
		}{
			{
				zones.HistoryListSearchKeywords{
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
				zones.HistoryListSearchKeywords{
					CommonSearchParams: api.CommonSearchParams{
						Type:   api.SearchTypeOR,
						Offset: int32(10),
						Limit:  int32(100),
					},
					FullText:    api.KeywordsString{"hogehoge", "🐰"},
					Description: api.KeywordsString{"🐇", "🍺"},
					Operator:    api.KeywordsString{"rabbit@example.jp", "SA0000000"},
				},
				/*
					_keywords_full_text[]
					_keywords_description[]
					_keywords_operator[]

				*/
				url.Values{
					"type":                    []string{"OR"},
					"offset":                  []string{"10"},
					"limit":                   []string{"100"},
					"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
					"_keywords_description[]": []string{"🐇", "🍺"},
					"_keywords_operator[]":    []string{"rabbit@example.jp", "SA0000000"},
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
	Describe("History", func() {
		var s, copy, nilSpec *zones.History
		BeforeEach(func() {
			s = &zones.History{}
		})
		Context("DeepCopy", func() {
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
