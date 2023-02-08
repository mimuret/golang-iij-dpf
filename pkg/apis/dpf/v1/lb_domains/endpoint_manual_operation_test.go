package lb_domains_test

import (
	"context"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("EndpointManualFailover", func() {
	var (
		cl    *testtool.TestClient
		err   error
		reqId string
	)
	Describe("EndpointManualFailover", func() {
		var s1 lb_domains.EndpointManualFailover
		BeforeEach(func() {
			cl = testtool.NewTestClient("", "http://localhost", nil)
			s1 = lb_domains.EndpointManualFailover{
				SiteAttributeMeta: lb_domains.SiteAttributeMeta{
					AttributeMeta: lb_domains.AttributeMeta{
						LBDomainID: "b0000000000001",
					},
					SiteResourceName: "site-1",
				},
				EndpointResourceName: "endpoint-1",
			}
		})
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1/failover", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1/failover"]).To(MatchJSON(`{}`))
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
					err = s1.SetPathParams("b0000000000010", "site-1", "endpoint-1")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
					Expect(s1.SiteResourceName).To(Equal("site-1"))
					Expect(s1.EndpointResourceName).To(Equal("endpoint-1"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-1", "endpoint-1", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-1", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
	Describe("EndpointManualFailback", func() {
		var s1 lb_domains.EndpointManualFailback
		BeforeEach(func() {
			cl = testtool.NewTestClient("", "http://localhost", nil)
			s1 = lb_domains.EndpointManualFailback{
				SiteAttributeMeta: lb_domains.SiteAttributeMeta{
					AttributeMeta: lb_domains.AttributeMeta{
						LBDomainID: "b0000000000001",
					},
					SiteResourceName: "site-1",
				},
				EndpointResourceName: "endpoint-1",
			}
		})
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1/failback", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1/failback"]).To(MatchJSON(`{}`))
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
					err = s1.SetPathParams("b0000000000010", "site-1", "endpoint-1")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
					Expect(s1.SiteResourceName).To(Equal("site-1"))
					Expect(s1.EndpointResourceName).To(Equal("endpoint-1"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-1", "endpoint-1", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-1", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
