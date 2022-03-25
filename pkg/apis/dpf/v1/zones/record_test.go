package zones_test

import (
	"context"
	"net/http"
	"net/url"

	"github.com/miekg/dns"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("records", func() {
	var (
		c      zones.Record
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.Record
		slist  zones.RecordList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.Record{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Id:     "r1",
			Name:   "www.example.jp.",
			TTL:    30,
			RRType: zones.TypeA,
			RData: []zones.RecordRDATA{
				{Value: "192.168.1.1"},
				{Value: "192.168.1.2"},
			},
			State:       zones.RecordStateApplied,
			Description: "SERVER(IPv4)",
			Operator:    "user1",
		}
		s2 = zones.Record{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Id:     "r2",
			Name:   "www.example.jp.",
			TTL:    30,
			RRType: zones.TypeAAAA,
			RData: []zones.RecordRDATA{
				{Value: "2001:db8::1"},
				{Value: "2001:db8::2"},
			},
			State:       zones.RecordStateToBeAdded,
			Description: "SERVER(IPv6)",
			Operator:    "user1",
		}
		slist = zones.RecordList{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Items: []zones.Record{s1, s2},
		}
	})
	Describe("Record", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "743C1E2428264E368D37233049AD49B5",
					"result": {
						"id": "r1",
						"name": "www.example.jp.",
						"ttl": 30,
						"rrtype": "A",
						"rdata": [
							{
								"value": "192.168.1.1"
							},
							{
								"value": "192.168.1.2"
							}
						],
						"state": 0,
						"description": "SERVER(IPv4)",
						"operator": "user1"
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "6B69B707901B479AAF83477B31340FEE",
					"result": {
						"id": "r2",
						"name": "www.example.jp.",
						"ttl": 30,
						"rrtype": "AAAA",
						"rdata": [
							{
								"value": "2001:db8::1"
							},
							{
								"value": "2001:db8::2"
							}
						],
						"state": 1,
						"description": "SERVER(IPv6)",
						"operator": "user1"
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns id=r1", func() {
				BeforeEach(func() {
					c = zones.Record{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
						Id: "r1",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("743C1E2428264E368D37233049AD49B5"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns id=r2", func() {
				BeforeEach(func() {
					c = zones.Record{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
						Id: "r2",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("6B69B707901B479AAF83477B31340FEE"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/zones/m2/records", httpmock.NewBytesResponder(202, bs1))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("create MX", func() {
				BeforeEach(func() {
					s := zones.Record{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m2",
						},
						Name:   "s1.example.jp.",
						TTL:    30,
						RRType: zones.TypeMX,
						RData: []zones.RecordRDATA{
							{Value: "10 mx1.example.jp."},
							{Value: "20 mx2.example.jp."},
						},
						Description: "MX SERVER",
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m2/records"]).To(MatchJSON(`{
							"name": "s1.example.jp.",
							"ttl": 30,
							"rrtype": "MX",
							"rdata": [
								{"value": "10 mx1.example.jp."},
								{"value": "20 mx2.example.jp."}
							],
							"description": "MX SERVER"
						}`))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("A record", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m1/records/r1"]).To(MatchJSON(`{
						"ttl": 30,
						"rdata": [
							{
								"value": "192.168.1.1"
							},
							{
								"value": "192.168.1.2"
							}
						],
						"description": "SERVER(IPv4)"
					}`))
				})
			})
			When("AAAA record", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m1/records/r2"]).To(MatchJSON(`{
						"ttl": 30,
						"rdata": [
							{
								"value": "2001:db8::1"
							},
							{
								"value": "2001:db8::2"
							}
						],
						"description": "SERVER(IPv6)"
						}`))
				})
			})
		})
		Context("Delete", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("remove A", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(context.Background(), &s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("remove AAAA", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(context.Background(), &s2)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
			})
		})
		Context("Cancel", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r1/changes", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r2/changes", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("cancel edit(A)", func() {
				BeforeEach(func() {
					reqId, err = cl.Cancel(context.Background(), &s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("cancel edit(AAAA)", func() {
				BeforeEach(func() {
					reqId, err = cl.Cancel(context.Background(), &s2)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
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
					err = s1.SetPathParams("m10", "r10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set CommonConfigId", func() {
					Expect(s1.GetZoneId()).To(Equal("m10"))
				})
				It("can set RecordId", func() {
					Expect(s1.Id).To(Equal("r10"))
				})
			})
			When("not enough arguments", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m10")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m19", "r10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2, "r1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (RecordId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m1", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.Record
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "records")

			Context("", func() {
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/zones/m1/records")
				})
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/records/r1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/zones/m1/records/r1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/zones/m1/records/r1")
				})
				When("action is ActionCancel", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCancel, http.MethodDelete, "/zones/m1/records/r1/changes")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})

			})
		})
	})
	Describe("RecordList", func() {
		var (
			c zones.RecordList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0987808B8D38403C8A1FBE66D72D68F7",
					"results": [
						{
							"id": "r1",
							"name": "www.example.jp.",
							"ttl": 30,
							"rrtype": "A",
							"rdata": [
								{
									"value": "192.168.1.1"
								},
								{
									"value": "192.168.1.2"
								}
							],
							"state": 0,
							"description": "SERVER(IPv4)",
							"operator": "user1"
						},
						{
							"id": "r2",
							"name": "www.example.jp.",
							"ttl": 30,
							"rrtype": "AAAA",
							"rdata": [
								{
									"value": "2001:db8::1"
								},
								{
									"value": "2001:db8::2"
								}
							],
							"state": 1,
							"description": "SERVER(IPv6)",
							"operator": "user1"
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = zones.RecordList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0987808B8D38403C8A1FBE66D72D68F7"))
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
					err = slist.SetPathParams("m1")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ZoneId", func() {
					Expect(slist.GetZoneId()).To(Equal("m1"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("m1", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.RecordList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "records")

			testtool.TestGetPathMethodForCountableList(&slist, "/zones/m1/records")
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
	Describe("CurrentRecordList", func() {
		var (
			c     zones.CurrentRecordList
			slist zones.CurrentRecordList
		)
		BeforeEach(func() {
			slist = zones.CurrentRecordList{
				AttributeMeta: zones.AttributeMeta{
					ZoneId: "m1",
				},
				Items: []zones.Record{s1, s2},
			}
		})
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/currents", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0987808B8D38403C8A1FBE66D72D68F7",
					"results": [
						{
							"id": "r1",
							"name": "www.example.jp.",
							"ttl": 30,
							"rrtype": "A",
							"rdata": [
								{
									"value": "192.168.1.1"
								},
								{
									"value": "192.168.1.2"
								}
							],
							"state": 0,
							"description": "SERVER(IPv4)",
							"operator": "user1"
						},
						{
							"id": "r2",
							"name": "www.example.jp.",
							"ttl": 30,
							"rrtype": "AAAA",
							"rdata": [
								{
									"value": "2001:db8::1"
								},
								{
									"value": "2001:db8::2"
								}
							],
							"state": 1,
							"description": "SERVER(IPv6)",
							"operator": "user1"
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = zones.CurrentRecordList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0987808B8D38403C8A1FBE66D72D68F7"))
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
					err = slist.SetPathParams("m1")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ZoneId", func() {
					Expect(slist.GetZoneId()).To(Equal("m1"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("m1", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.CurrentRecordList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "records/currents")

			testtool.TestGetPathMethodForCountableList(&slist, "/zones/m1/records/currents")
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
	Describe("RecordListSearchKeywords", func() {
		Context("GetValues", func() {
			var (
				testcase = []struct {
					keyword zones.RecordListSearchKeywords
					values  url.Values
				}{
					{
						zones.RecordListSearchKeywords{
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
						zones.RecordListSearchKeywords{
							CommonSearchParams: api.CommonSearchParams{
								Type:   api.SearchTypeOR,
								Offset: int32(10),
								Limit:  int32(100),
							},
							FullText:    api.KeywordsString{"hogehoge", "🐰"},
							Name:        api.KeywordsString{"example.jp.", "128/0.0.168.192.in-addr.arpa."},
							TTL:         []int32{300, 2147483647},
							RRType:      zones.KeywordsType{"A", "CNAME", "HTTPS"},
							RData:       api.KeywordsString{"192.168.0.1", "\"hogehoge\""},
							Description: api.KeywordsString{"🐇", "🍺"},
							Operator:    api.KeywordsString{"rabbit@example.jp", "SA0000000"},
						},
						/*
							_keywords_full_text[]
							_keywords_name[]
							_keywords_ttl[]
							_keywords_rrtype[]
							_keywords_rdata[]
							_keywords_description[]
							_keywords_operator[]

						*/
						url.Values{
							"type":                    []string{"OR"},
							"offset":                  []string{"10"},
							"limit":                   []string{"100"},
							"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
							"_keywords_name[]":        []string{"example.jp.", "128/0.0.168.192.in-addr.arpa."},
							"_keywords_ttl[]":         []string{"300", "2147483647"},
							"_keywords_rrtype[]":      []string{"A", "CNAME", "HTTPS"},
							"_keywords_rdata[]":       []string{"192.168.0.1", "\"hogehoge\""},
							"_keywords_description[]": []string{"🐇", "🍺"},
							"_keywords_operator[]":    []string{"rabbit@example.jp", "SA0000000"},
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
	Context("Type", func() {
		var testcase = []struct {
			Type zones.Type
			Str  string
			Code uint16
		}{
			{zones.TypeSOA, "SOA", dns.TypeSOA},
			{zones.TypeA, "A", dns.TypeA},
			{zones.TypeAAAA, "AAAA", dns.TypeAAAA},
			{zones.TypeCAA, "CAA", dns.TypeCAA},
			{zones.TypeCNAME, "CNAME", dns.TypeCNAME},
			{zones.TypeDS, "DS", dns.TypeDS},
			{zones.TypeNS, "NS", dns.TypeNS},
			{zones.TypeMX, "MX", dns.TypeMX},
			{zones.TypeNAPTR, "NAPTR", dns.TypeNAPTR},
			{zones.TypeSRV, "SRV", dns.TypeSRV},
			{zones.TypeTXT, "TXT", dns.TypeTXT},
			{zones.TypeTLSA, "TLSA", dns.TypeTLSA},
			{zones.TypePTR, "PTR", dns.TypePTR},
			{zones.TypeSVCB, "SVCB", dns.TypeSVCB},
			{zones.TypeHTTPS, "HTTPS", dns.TypeHTTPS},
			{zones.TypeANAME, "ANAME", 65280},
		}
		Context("String", func() {
			It("returns string", func() {
				for _, tc := range testcase {
					Expect(tc.Type.String()).To(Equal(tc.Str))
				}
			})
		})
		Context("Uint16", func() {
			It("returns uint16", func() {
				for _, tc := range testcase {
					Expect(tc.Type.Uint16()).To(Equal(tc.Code))
				}
			})
		})
		Context("Uint16ToType", func() {
			It("returns Type", func() {
				for _, tc := range testcase {
					Expect(zones.Uint16ToType(tc.Code)).To(Equal(tc.Type))
				}
			})
		})
	})
	Context("RecordStateToString", func() {
		Context("String", func() {
			It("returns string", func() {
				Expect(zones.RecordStateApplied.String()).To(Equal("Applied"))
			})
		})
	})
	Context("RecordRDATA", func() {
		var s, copy, nilSpec *zones.RecordRDATA
		BeforeEach(func() {
			s = &zones.RecordRDATA{
				Value: "hoge",
			}
		})
		Context("String", func() {
			It("returns Value", func() {
				Expect(s.String()).To(Equal("hoge"))
			})
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
	Context("RecordRDATASlice", func() {
		var s, copy, nilSpec zones.RecordRDATASlice
		BeforeEach(func() {
			s = zones.RecordRDATASlice{
				zones.RecordRDATA{
					Value: "192.168.0.1",
				},
				zones.RecordRDATA{
					Value: "192.168.0.2",
				},
			}
		})
		Context("String", func() {
			It("returns comma separated string", func() {
				Expect(s.String()).To(Equal("192.168.0.1,192.168.0.2"))
			})
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
