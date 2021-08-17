package zones_test

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("zones", func() {
	var (
		cl    *testtool.TestClient
		err   error
		reqId string
		s     zones.ZoneApply
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s = zones.ZoneApply{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Description: "commit records",
		}
	})
	Describe("ZoneApply", func() {
		Context("Apply", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/changes", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Apply(&s, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/zones/m1/changes"]).To(MatchJSON(`{
							"description": "commit records"
						}`))
			})
		})
		Context("Cancel", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/changes", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Cancel(&s)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/zones/m1/changes"]).To(Equal(""))
			})
		})
		Context("SetPathParams", func() {
			When("no arguments, nothing to do", func() {
				BeforeEach(func() {
					err = s.SetPathParams()
				})
				It("returns error", func() {
					Expect(err).To(Succeed())
				})
			})
			When("enough arguments", func() {
				BeforeEach(func() {
					err = s.SetPathParams("m100")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ZonetId", func() {
					Expect(s.ZoneId).To(Equal("m100"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s.SetPathParams("m100", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = s.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.ZoneApply
			testtool.TestDeepCopyObject(&s, nilSpec)
			testtool.TestGetName(&s, "apply")
			testtool.TestGetGroup(&s, "zones")
			Context("GetPathMethod", func() {
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s, api.ActionApply, http.MethodPatch, "/zones/m1/changes")
				})
				When("other", func() {
					testtool.TestGetPathMethod(&s, api.ActionRead, "", "")
				})
			})
		})
	})
})
