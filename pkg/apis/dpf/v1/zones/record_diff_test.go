package zones_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("records/diffs", func() {
	var (
		c          zones.RecordDiffList
		cl         *testtool.TestClient
		err        error
		reqId      string
		s1, s2, s3 zones.RecordDiff
		slist      zones.RecordDiffList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.RecordDiff{
			New: &zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneID: "m1",
				},
				Id:     "r1",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeA,
				RData: []zones.RecordRDATA{
					{Value: "192.168.1.1"},
					{Value: "192.168.1.2"},
				},
				State:       zones.RecordStateBeforeUpdate,
				Description: "SERVER Site 1(IPv4)",
				Operator:    "user1",
			},
		}
		s2 = zones.RecordDiff{
			Old: &zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneID: "m1",
				},
				Id:     "r1",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeA,
				RData: []zones.RecordRDATA{
					{Value: "192.168.2.1"},
					{Value: "192.168.2.2"},
				},
				State:       zones.RecordStateApplied,
				Description: "SERVER Site 2(IPv4)",
				Operator:    "user1",
			},
		}
		s3 = zones.RecordDiff{
			New: &zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneID: "m1",
				},
				Id:     "r3",
				Name:   "www.example.jp.",
				TTL:    300,
				RRType: zones.TypeTXT,
				RData: []zones.RecordRDATA{
					{Value: "new"},
				},
				State:       zones.RecordStateToBeAdded,
				Description: "SERVER(TXT)",
				Operator:    "user1",
			},
			Old: &zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneID: "m1",
				},
				Id:     "r3",
				Name:   "www.example.jp.",
				TTL:    300,
				RRType: zones.TypeTXT,
				RData: []zones.RecordRDATA{
					{Value: "old"},
				},
				State:       zones.RecordStateToBeAdded,
				Description: "SERVER(TXT)",
				Operator:    "user1",
			},
		}
		slist = zones.RecordDiffList{
			AttributeMeta: zones.AttributeMeta{
				ZoneID: "m1",
			},
			Items: []zones.RecordDiff{s1, s2, s3},
		}
	})
	Context("RecordDiffList", func() {
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/diffs", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "C36898E9087C4E2F85C94D146BC6FEDD",
					"results": [
						{
							"new": {
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
								"state": 5,
								"description": "SERVER Site 1(IPv4)",
								"operator": "user1"	
							}
						},
						{
							"old": {
								"id": "r1",
								"name": "www.example.jp.",
								"ttl": 30,
								"rrtype": "A",
								"rdata": [
									{
										"value": "192.168.2.1"
									},
									{
										"value": "192.168.2.2"
									}
								],
								"state": 0,
								"description": "SERVER Site 2(IPv4)",
								"operator": "user1"
							}
						},
						{
							"new": {
								"id": "r3",
								"name": "www.example.jp.",
								"ttl": 300,
								"rrtype": "TXT",
								"rdata": [
									{
										"value": "new"
									}
								],
								"state": 1,
								"description": "SERVER(TXT)",
								"operator": "user1"
							},
							"old": {
								"id": "r3",
								"name": "www.example.jp.",
								"ttl": 300,
								"rrtype": "TXT",
								"rdata": [
									{
										"value": "old"
									}
								],
								"state": 1,
								"description": "SERVER(TXT)",
								"operator": "user1"
							}
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/diffs/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EE547488A8B74999AD630664802B4E25",
					"result": {
						"count": 3
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = zones.RecordDiffList{
						AttributeMeta: zones.AttributeMeta{
							ZoneID: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("C36898E9087C4E2F85C94D146BC6FEDD"))
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
				It("can set ZoneID", func() {
					Expect(slist.GetZoneID()).To(Equal("m1"))
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
			When("arguments type missmatch (ZoneID)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.RecordDiffList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "records/diffs")

			testtool.TestGetPathMethodForCountableList(&slist, "/zones/m1/records/diffs")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 3)
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
	Describe("RecordDiff", func() {
		var s, copy, nilSpec *zones.RecordDiff
		BeforeEach(func() {
			s = &zones.RecordDiff{}
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
