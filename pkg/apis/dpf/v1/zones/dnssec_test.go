package zones_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("dnssec", func() {
	var (
		c      zones.Dnssec
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.Dnssec
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.Dnssec{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Enabled: types.Enabled,
			State:   zones.DnssecStateEnable,
			DsState: zones.DSStateDisclose,
		}
		s2 = zones.Dnssec{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m2",
			},
			Enabled: types.Disabled,
			State:   zones.DnssecStateEnable,
			DsState: zones.DSStateDisclose,
		}
	})
	Describe("Dnssec", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/dnssec", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "533F010C0DEC4B71891D1491C5308A10",
					"result": {
						"enabled": 1,
						"state": 2,
						"ds_state": 3
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns m1", func() {
				BeforeEach(func() {
					c = zones.Dnssec{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("533F010C0DEC4B71891D1491C5308A10"))
					Expect(c).To(Equal(s1))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/dnssec", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m2/dnssec", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("enable", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m1/dnssec"]).To(MatchJSON(`{
								"enabled": 1
					}`))
				})
			})
			When("disable", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m2/dnssec"]).To(MatchJSON(`{
								"enabled": 0
					}`))
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
					err = s1.SetPathParams("m10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.GetZoneId()).To(Equal("m10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m10", 2)
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
			var nilSpec *zones.Dnssec
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "dnssec")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/dnssec")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/zones/m1/dnssec")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Context("DnssecStateToString", func() {
		Context("String", func() {
			Expect(zones.DnssecStateEnable.String()).To(Equal("Enable"))
		})
	})
	Context("DSStateToSString", func() {
		Context("String", func() {
			Expect(zones.DSStateClose.String()).To(Equal("Close"))
		})
	})
})
