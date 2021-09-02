package zones_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("history", func() {
	var (
		c      zones.HistoryText
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.HistoryText
	)
	BeforeEach(func() {
		atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:50:37.396Z")
		Expect(err).To(Succeed())

		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.HistoryText{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			History: zones.History{
				Id:          1,
				CommittedAt: atTime,
				Description: "commit 1",
				Operator:    "user1",
			},
			Text: "zone data 1",
		}
		s2 = zones.HistoryText{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			History: zones.History{
				Id:          2,
				CommittedAt: atTime,
				Description: "commit 2",
				Operator:    "user2",
			},
			Text: "zone data 2",
		}
	})
	Describe("DsRecordList", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories/1/text", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "CC2BBC44169940CE96D7E8A99FA98958",
					"result": {
						"id": 1,
						"committed_at": "2021-06-20T10:50:37.396Z",
						"description": "commit 1",
						"operator": "user1",
						"text": "zone data 1"
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories/2/text", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "CFBD5CBBE7C147C9BD2549551F018D11",
					"result": {
						"id": 2,
						"committed_at": "2021-06-20T10:50:37.396Z",
						"description": "commit 2",
						"operator": "user2",
						"text": "zone data 2"
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns id1", func() {
				BeforeEach(func() {
					c = zones.HistoryText{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
						History: zones.History{
							Id: 1,
						},
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("CC2BBC44169940CE96D7E8A99FA98958"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns id2", func() {
				BeforeEach(func() {
					c = zones.HistoryText{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
						History: zones.History{
							Id: 2,
						},
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("CFBD5CBBE7C147C9BD2549551F018D11"))
					Expect(c).To(Equal(s2))
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
					err = s1.SetPathParams("m10", 10)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.GetZoneId()).To(Equal("m10"))
				})
				It("can set HistoryId", func() {
					Expect(s1.Id).To(Equal(int64(10)))
				})
			})
			When("not enough arguments", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m10")
				})
				It("not returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m10", 2, 1)
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
			When("arguments type missmatch (HistoryId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m1", "h1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.HistoryText
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "zone_histories/1/text")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/zone_histories/1/text")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
})
