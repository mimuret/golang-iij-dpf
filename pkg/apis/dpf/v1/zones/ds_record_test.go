package zones_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("ds_records", func() {
	var (
		c      zones.DsRecordList
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.DsRecord
		slist  zones.DsRecordList
	)
	BeforeEach(func() {
		atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:23:51.071Z")
		Expect(err).To(Succeed())

		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.DsRecord{
			RRSet:     "46369 8 2 39F054DCB3EC1E93D8AE6D8F1AAAD91794055EA36895045FAF6F65F0 2FEBC579",
			TransitAt: atTime,
		}
		s2 = zones.DsRecord{
			RRSet:     "57367 8 2 C069902A17BC8E19F1D5809272219B694C09EC5BCF4AAD4911FC765B 1C233A69",
			TransitAt: atTime,
		}
		slist = zones.DsRecordList{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Items: []zones.DsRecord{s1, s2},
		}
	})
	Describe("DsRecordList", func() {
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/ds_records", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "CD445FC65BBA4977B78B347A87633947",
					"results": [
						{
							"rrset": "46369 8 2 39F054DCB3EC1E93D8AE6D8F1AAAD91794055EA36895045FAF6F65F0 2FEBC579",
							"transited_at": "2021-06-20T10:23:51.071Z"
						},
						{
							"rrset": "57367 8 2 C069902A17BC8E19F1D5809272219B694C09EC5BCF4AAD4911FC765B 1C233A69",
							"transited_at": "2021-06-20T10:23:51.071Z"
						}
					]
				}`)))

			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns m1", func() {
				BeforeEach(func() {
					c = zones.DsRecordList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("CD445FC65BBA4977B78B347A87633947"))
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
				It("can set ContractId", func() {
					Expect(slist.GetZoneId()).To(Equal("m10"))
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
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.DsRecordList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "ds_records")
			testtool.TestGetPathMethodForList(&slist, "/zones/m1/ds_records")
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
	Describe("DsRecord", func() {
		var (
			s, copy, nilSpec *zones.DsRecord
		)
		BeforeEach(func() {
			s = &zones.DsRecord{}
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
