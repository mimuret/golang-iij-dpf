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

var _ = Describe("managed_dns_servers", func() {
	var (
		c     zones.ManagedDnsList
		cl    *testtool.TestClient
		err   error
		reqId string
		slist zones.ManagedDnsList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		slist = zones.ManagedDnsList{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Items: []string{
				"ns1.example.jp.",
				"ns1.example.net.",
				"ns1.example.com.",
			},
		}
	})
	Describe("ManagedDnsList", func() {
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/managed_dns_servers", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "F4F8591B01B84EFBBA0AD17029DB7DA4",
					"results": [
						"ns1.example.jp.",
						"ns1.example.net.",
						"ns1.example.com."
					]
				}`)))

			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns managed_dns servers", func() {
				BeforeEach(func() {
					c = zones.ManagedDnsList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("F4F8591B01B84EFBBA0AD17029DB7DA4"))
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
				It("can set ZoneId", func() {
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
			var nilSpec *zones.ManagedDnsList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "managed_dns_servers")

			testtool.TestGetPathMethodForList(&slist, "/zones/m1/managed_dns_servers")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 3)
			Context("Index", func() {
				When("index exist", func() {
					It("returns item", func() {
						Expect(slist.Index(0)).To(Equal("ns1.example.jp."))
					})
				})
			})
		})
	})
})
