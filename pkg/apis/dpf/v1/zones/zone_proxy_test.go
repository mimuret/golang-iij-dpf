package zones_test

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("zone_proxy", func() {
	var (
		c      zones.ZoneProxy
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.ZoneProxy
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.ZoneProxy{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Enabled: types.Enabled,
		}
		s2 = zones.ZoneProxy{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m2",
			},
			Enabled: types.Disabled,
		}
	})
	Describe("ZoneProxy", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_proxy", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0607103482984C3B9EC782480423CE63",
					"result": {
						"enabled": 1
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m2/zone_proxy", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "63896F8889174DF59FF79F5C5F4B173A",
					"result": {
						"enabled": 0
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns id=r1", func() {
				BeforeEach(func() {
					c = zones.ZoneProxy{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0607103482984C3B9EC782480423CE63"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns id=r2", func() {
				BeforeEach(func() {
					c = zones.ZoneProxy{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m2",
						},
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("63896F8889174DF59FF79F5C5F4B173A"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/zone_proxy", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m2/zone_proxy", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("enable", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m1/zone_proxy"]).To(MatchJSON(`{
								"enabled": 1
							}`))
				})
			})
			When("disabled", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/zones/m2/zone_proxy"]).To(MatchJSON(`{
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
					err = s1.SetPathParams("m100")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ZonetId", func() {
					Expect(s1.ZoneId).To(Equal("m100"))
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
			var nilSpec *zones.ZoneProxy
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "zone_proxy")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/zone_proxy")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/zones/m1/zone_proxy")
				})
				When("other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
})
