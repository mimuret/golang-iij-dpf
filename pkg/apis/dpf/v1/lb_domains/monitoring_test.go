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

var _ = Describe("Monitoring", func() {
	var (
		c      lb_domains.Monitoring
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 lb_domains.Monitoring
		slist  lb_domains.MonitoringList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = lb_domains.Monitoring{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			MonitoringCommon: lb_domains.MonitoringCommon{
				ResourceName: "id1",
				Name:         "monitoring-1",
				MType:        lb_domains.MonitoringMtypePing,
				Description:  "comment 1",
			},
			Props: &lb_domains.MonitoringPorpsPING{
				MonitoringPorpsCommon: lb_domains.MonitoringPorpsCommon{
					Location: lb_domains.MonitoringPropsLocationAll,
					Interval: 30,
					Holdtime: 0,
					Timeout:  1,
				},
			},
		}
		s2 = lb_domains.Monitoring{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			MonitoringCommon: lb_domains.MonitoringCommon{
				ResourceName: "id2",
				Name:         "monitoring-2",
				MType:        lb_domains.MonitoringMtypeStatic,
				Description:  "comment 2",
			},
			Props: &lb_domains.MonitoringPorpsStatic{
				Result: lb_domains.MonitoringPorpsStaticStatusUp,
			},
		}
		slist = lb_domains.MonitoringList{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			Items: []lb_domains.Monitoring{s1, s2},
		}
	})
	Describe("Monitoring", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/monitorings/id1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
							"resource_name": "id1",
							"name": "monitoring-1",
							"mtype": "ping",
							"description": "comment 1",
							"props": {
								"location": "all",
								"interval": 30,
								"holdtime": 0,
								"timeout": 1
							}
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/monitorings/id2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
						"resource_name": "id2",
						"name": "monitoring-2",
						"mtype": "static",
						"description": "comment 2",
						"props": {
							"result": "up"
						}
				}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns LBDomain m1", func() {
				BeforeEach(func() {
					c = lb_domains.Monitoring{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
						MonitoringCommon: lb_domains.MonitoringCommon{
							ResourceName: "id1",
						},
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
					c = lb_domains.Monitoring{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
						MonitoringCommon: lb_domains.MonitoringCommon{
							ResourceName: "id2",
						},
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
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/monitorings", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/monitorings"]).To(MatchJSON(`{
					"resource_name": "id1",
					"name": "monitoring-1",
					"mtype": "ping",
					"description": "comment 1",
					"props": {
							"location": "all",
							"interval": 30,
							"holdtime": 0,
							"timeout": 1
						}
				}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/lb_domains/b0000000000001/monitorings/id1", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/monitorings/id1"]).To(MatchJSON(`{
						"name": "monitoring-1",
						"description": "comment 1",
						"props": {
							"location": "all",
							"interval": 30,
							"holdtime": 0,
							"timeout": 1
						}
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
			var nilSpec *lb_domains.Monitoring
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "monitorings")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/lb_domains/b0000000000001/monitorings/id1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/lb_domains/b0000000000001/monitorings/id1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/lb_domains/b0000000000001/monitorings/id1")
				})
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/lb_domains/b0000000000001/monitorings")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("MonitoringList", func() {
		var c lb_domains.MonitoringList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/monitorings", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
					"results": [
						{
							"resource_name": "id1",
							"name": "monitoring-1",
							"mtype": "ping",
							"description": "comment 1",
							"props": {
								"location": "all",
								"interval": 30,
								"holdtime": 0,
								"timeout": 1
							}
						},
						{
							"resource_name": "id2",
							"name": "monitoring-2",
							"mtype": "static",
							"description": "comment 2",
							"props": {
								"result": "up"
							}
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = lb_domains.MonitoringList{
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
			var nilSpec *lb_domains.MonitoringList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "monitorings")
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
