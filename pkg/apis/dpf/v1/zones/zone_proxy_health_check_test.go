package zones_test

import (
	net "net"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("zone_proxy", func() {
	var (
		c      zones.ZoneProxyHealthCheckList
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.ZoneProxyHealthCheck
		slist  zones.ZoneProxyHealthCheckList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.ZoneProxyHealthCheck{
			Address:  net.ParseIP("192.168.0.1"),
			Status:   zones.ZoneProxyStatusSuccess,
			TsigName: "",
			Enabled:  types.Disabled,
		}
		s2 = zones.ZoneProxyHealthCheck{
			Address:  net.ParseIP("2001:db8::1"),
			Status:   zones.ZoneProxyStatusFail,
			TsigName: "tsig2",
			Enabled:  types.Enabled,
		}
		slist = zones.ZoneProxyHealthCheckList{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Items: []zones.ZoneProxyHealthCheck{s1, s2},
		}
	})
	Describe("ZoneProxy", func() {
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_proxy/health_check", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "751E0AB9585E48059090FA1B46E5B7A3",
					"results": [
						{
							"address": "192.168.0.1",
							"status": "success",
							"tsig_name": "",
							"enabled": 0
						},
						{
							"address": "2001:db8::1",
							"status": "fail",
							"tsig_name": "tsig2",
							"enabled": 1
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = zones.ZoneProxyHealthCheckList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("751E0AB9585E48059090FA1B46E5B7A3"))
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
					err = slist.SetPathParams("m100")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ZonetId", func() {
					Expect(slist.ZoneId).To(Equal("m100"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("m100", 2)
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
			var nilSpec *zones.ZoneProxyHealthCheckList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "zone_proxy/health_check")

			testtool.TestGetPathMethodForList(&slist, "/zones/m1/zone_proxy/health_check")
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
	Describe("ZoneProxyHealthCheck", func() {
		var (
			s, copy, nilSpec *zones.ZoneProxyHealthCheck
		)
		BeforeEach(func() {
			s = &zones.ZoneProxyHealthCheck{}
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
