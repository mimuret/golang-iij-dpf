package lb_domains_test

import (
	"context"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("Site", func() {
	var (
		c      lb_domains.Site
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 lb_domains.Site
		slist  lb_domains.SiteList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = lb_domains.Site{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			ResourceName: "site-1",
			Name:         "site-a",
			Description:  "site A comment",
			RRType:       lb_domains.SiteRRTypeA,
			LiveStatus:   lb_domains.StatusDown,
			Endpoints:    []lb_domains.Endpoint{},
		}
		s2 = lb_domains.Site{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			ResourceName: "rule-2",
			Name:         "site-aaaa",
			Description:  "site AAAA comment",
			RRType:       lb_domains.SiteRRTypeAAAA,
			LiveStatus:   lb_domains.StatusUp,
			Endpoints:    []lb_domains.Endpoint{},
		}
		slist = lb_domains.SiteList{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			Items: []lb_domains.Site{s1, s2},
		}
	})
	Describe("Site", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites/site-1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
						"resource_name": "site-1",
						"name": "site-a",
						"description": "site A comment",
						"rrtype": "A",
						"live_status": "down",
						"endpoints": []
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites/rule-2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
						"resource_name": "rule-2",
						"name": "site-aaaa",
						"description": "site AAAA comment",
						"rrtype": "AAAA",
						"live_status": "up",
						"endpoints": []
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns LBDomain m1", func() {
				BeforeEach(func() {
					c = lb_domains.Site{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
						ResourceName: "site-1",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("76D7462F50E3417C96898A827E0D1EC1"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns LBDomain m2", func() {
				BeforeEach(func() {
					c = lb_domains.Site{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
						ResourceName: "rule-2",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("B10F0042282A46C18BFB440B5B58BF8C"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/sites", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Create(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites"]).To(MatchJSON(`{
					"resource_name": "site-1",
					"name": "site-a",
					"rrtype": "A",
					"description": "site A comment"
				}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/lb_domains/b0000000000001/sites/site-1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Update(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites/site-1"]).To(MatchJSON(`{
					"name": "site-a",
					"description": "site A comment"
			}`))
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
					err = s1.SetPathParams("b0000000000010", "id1")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
					Expect(s1.ResourceName).To(Equal("id1"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "id1", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2, "id1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *lb_domains.Site
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "sites")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/lb_domains/b0000000000001/sites/site-1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/lb_domains/b0000000000001/sites/site-1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/lb_domains/b0000000000001/sites/site-1")
				})
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/lb_domains/b0000000000001/sites")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("SiteList", func() {
		var c lb_domains.SiteList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
					"results": [
						{
							"resource_name": "site-1",
							"name": "site-a",
							"description": "site A comment",
							"rrtype": "A",
							"live_status": "down",
							"endpoints": []
						},
						{
							"resource_name": "rule-2",
							"name": "site-aaaa",
							"description": "site AAAA comment",
							"rrtype": "AAAA",
							"live_status": "up",
							"endpoints": []
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = lb_domains.SiteList{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("2C18922432DC48D485613F5383A7ED8E"))
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
					err = slist.SetPathParams("b0000000000010")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(slist.LBDomainID).To(Equal("b0000000000010"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("b0000000000010", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *lb_domains.SiteList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "sites")
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
})
