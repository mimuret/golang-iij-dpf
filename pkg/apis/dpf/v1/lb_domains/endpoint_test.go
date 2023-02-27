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

var _ = Describe("Endpoint", func() {
	var (
		c      lb_domains.Endpoint
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 lb_domains.Endpoint
		slist  lb_domains.EndpointList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = lb_domains.Endpoint{
			SiteAttributeMeta: lb_domains.SiteAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				SiteResourceName: "site-1",
			},
			ResourceName:     "endpoint-1",
			Name:             "host-1",
			MonitoringTarget: "www.example.jp.",
			Description:      "endpoint#1",
			Weight:           1,
			ManualFaillback:  false,
			ManualFaillOver:  false,
			Enabled:          false,
			LiveStatus:       lb_domains.StatusUp,
			ReadyStatus:      lb_domains.StatusDown,
			Rdata:            []lb_domains.EndpointRdata{{Value: "192.168.0.1"}, {Value: "192.168.1.1"}},
			Monitorings: []lb_domains.MonitoringEndpoint{
				{
					MonitoringResourceName: "id1",
					Enabled:                true,
					Monitoring: &lb_domains.Monitoring{
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
					},
				},
			},
		}
		s2 = lb_domains.Endpoint{
			SiteAttributeMeta: lb_domains.SiteAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				SiteResourceName: "site-1",
			},
			ResourceName:     "endpoint-2",
			Name:             "host-2",
			MonitoringTarget: "2001:db8::1",
			Description:      "endpoint#2",
			Weight:           255,
			ManualFaillback:  true,
			ManualFaillOver:  true,
			Enabled:          true,
			LiveStatus:       lb_domains.StatusDown,
			ReadyStatus:      lb_domains.StatusDown,
			Rdata:            []lb_domains.EndpointRdata{{Value: "192.168.0.2"}, {Value: "192.168.1.2"}},
			Monitorings:      []lb_domains.MonitoringEndpoint{},
		}
		slist = lb_domains.EndpointList{
			SiteAttributeMeta: lb_domains.SiteAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				SiteResourceName: "site-1",
			},
			Items: []lb_domains.Endpoint{s1, s2},
		}
	})
	Describe("Endpoint", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
						"resource_name": "endpoint-1",
						"name": "host-1",
						"monitoring_target": "www.example.jp.",
						"description": "endpoint#1",
						"weight": 1,
						"manual_failback": false,
						"manual_failover": false,
						"enabled": false,
						"live_status": "up",
						"ready_status": "down",
						"rdata": [
							{
								"value": "192.168.0.1"
							},
							{
								"value": "192.168.1.1"
							}
						],
						"monitorings": [
							{
								"resource_name": "id1",
								"enabled": true,
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
						]
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
						"resource_name": "endpoint-2",
						"name": "host-2",
						"monitoring_target": "2001:db8::1",
						"description": "endpoint#2",
						"weight": 255,
						"manual_failback": true,
						"manual_failover": true,
						"enabled": true,
						"live_status": "down",
						"ready_status": "down",
						"rdata": [
							{
								"value": "192.168.0.2"
							},
							{
								"value": "192.168.1.2"
							}
						],
						"monitorings": []
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns endpoint-1", func() {
				BeforeEach(func() {
					c = lb_domains.Endpoint{
						SiteAttributeMeta: lb_domains.SiteAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							SiteResourceName: "site-1",
						},
						ResourceName: "endpoint-1",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("76D7462F50E3417C96898A827E0D1EC1"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns endpoint-2", func() {
				BeforeEach(func() {
					c = lb_domains.Endpoint{
						SiteAttributeMeta: lb_domains.SiteAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							SiteResourceName: "site-1",
						},
						ResourceName: "endpoint-2",
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
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites/site-1/endpoints"]).To(MatchJSON(`{
					"resource_name": "endpoint-1",
					"name": "host-1",
					"monitoring_target": "www.example.jp.",
					"description": "endpoint#1",
					"weight": 1,
					"manual_failback": false,
					"manual_failover": false,
					"enabled": false,
					"rdata": [
						{
							"value": "192.168.0.1"
						},
						{
							"value": "192.168.1.1"
						}
					],
					"monitorings": [
						{"resource_name": "id1", "enabled": true}
					]
				}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1"]).To(MatchJSON(`{
					"name": "host-1",
					"monitoring_target": "www.example.jp.",
					"description": "endpoint#1",
					"weight": 1,
					"manual_failback": false,
					"manual_failover": false,
					"enabled": false,
					"rdata": [
						{
							"value": "192.168.0.1"
						},
						{
							"value": "192.168.1.1"
						}
					],
					"monitorings": [
						{"resource_name": "id1", "enabled": true}
					]
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
					err = s1.SetPathParams("b0000000000010", "site-10", "endpoint-10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
					Expect(s1.SiteResourceName).To(Equal("site-10"))
					Expect(s1.ResourceName).To(Equal("endpoint-10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-10", "endpoint-10", 0)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "site-10", 0)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *lb_domains.Endpoint
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "endpoints")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/lb_domains/b0000000000001/sites/site-1/endpoints/endpoint-1")
				})
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/lb_domains/b0000000000001/sites/site-1/endpoints")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("EndpointList", func() {
		var c lb_domains.EndpointList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/sites/site-1/endpoints", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
					"results": [
						{
							"resource_name": "endpoint-1",
							"name": "host-1",
							"monitoring_target": "www.example.jp.",
							"description": "endpoint#1",
							"weight": 1,
							"manual_failback": false,
							"manual_failover": false,
							"enabled": false,
							"live_status": "up",
							"ready_status": "down",
							"enabled": false,
							"rdata": [
								{
									"value": "192.168.0.1"
								},
								{
									"value": "192.168.1.1"
								}
							],
							"monitorings": [
								{
									"resource_name": "id1",
									"enabled": true,
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
							]
						},
						{
							"resource_name": "endpoint-2",
							"name": "host-2",
							"monitoring_target": "2001:db8::1",
							"description": "endpoint#2",
							"weight": 255,
							"manual_failback": true,
							"manual_failover": true,
							"enabled": true,
							"live_status": "down",
							"ready_status": "down",
							"rdata": [
								{
									"value": "192.168.0.2"
								},
								{
									"value": "192.168.1.2"
								}
							],
							"monitorings": []
							}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = lb_domains.EndpointList{
						SiteAttributeMeta: lb_domains.SiteAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							SiteResourceName: "site-1",
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
					err = slist.SetPathParams("b0000000000010", "site-10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(slist.LBDomainID).To(Equal("b0000000000010"))
					Expect(slist.SiteResourceName).To(Equal("site-10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("b0000000000010", "site-10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("b0000000000010", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *lb_domains.EndpointList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "endpoints")
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
