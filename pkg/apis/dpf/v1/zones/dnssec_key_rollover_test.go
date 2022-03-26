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
)

var _ = Describe("dnssec/ksk_rollover", func() {
	var (
		cl    *testtool.TestClient
		err   error
		reqId string
		s1    zones.DnssecKskRollover
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.DnssecKskRollover{
			AttributeMeta: zones.AttributeMeta{
				ZoneID: "m1",
			},
		}
	})
	Describe("DnssecKskRollover", func() {
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/dnssec/ksk_rollover", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Apply(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/zones/m1/dnssec/ksk_rollover"]).To(MatchJSON(`{}`))
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
				It("can set ContractID", func() {
					Expect(s1.GetZoneID()).To(Equal("m10"))
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
			When("arguments type missmatch (ZoneID)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.DnssecKskRollover
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "dnssec/ksk_rollover")

			Context("GetPathMethod", func() {
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/zones/m1/dnssec/ksk_rollover")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, "", "")
				})
			})
		})
	})
})
