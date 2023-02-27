package lb_domains_test

import (
	"context"
	_ "embed"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

//go:embed testdata/config_get.json
var configReadJson []byte

//go:embed testdata/config_put.json
var configUpdateJson []byte

var _ = Describe("Config", func() {
	var (
		c     lb_domains.Config
		cl    *testtool.TestClient
		err   error
		reqId string
		s1    lb_domains.Config
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = lb_domains.Config{
			AttributeMeta: lb_domains.AttributeMeta{
				LBDomainID: "b0000000000001",
			},
			Sites: []lb_domains.Site{
				{
					AttributeMeta: lb_domains.AttributeMeta{
						LBDomainID: "b0000000000001",
					},
					ResourceName: "site-a",
					Name:         "site(A)",
					Description:  "site(A) comment",
					RRType:       lb_domains.SiteRRTypeA,
					LiveStatus:   lb_domains.StatusDown,
					Endpoints: []lb_domains.Endpoint{
						{
							SiteAttributeMeta: lb_domains.SiteAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								SiteResourceName: "site-a",
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
						},
					},
				},
				{
					AttributeMeta: lb_domains.AttributeMeta{
						LBDomainID: "b0000000000001",
					},
					ResourceName: "site-aaaa",
					Name:         "site(AAAA)",
					Description:  "site(AAAA) comment",
					RRType:       lb_domains.SiteRRTypeAAAA,
					LiveStatus:   lb_domains.StatusUp,
					Endpoints: []lb_domains.Endpoint{
						{
							SiteAttributeMeta: lb_domains.SiteAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								SiteResourceName: "site-aaaa",
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
							Rdata:            []lb_domains.EndpointRdata{{Value: "2001:db8::1"}, {Value: "2001:db8::2"}},
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
						},
					},
				},
			},
			Monitorings: []lb_domains.Monitoring{
				{
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
			Rules: []lb_domains.Rule{
				{
					AttributeMeta: lb_domains.AttributeMeta{
						LBDomainID: "b0000000000001",
					},
					ResourceName: "rule-a",
					Name:         "default",
					Description:  "default rule",
					Methods: []lb_domains.RuleMethod{
						{
							RuleAttributeMeta: lb_domains.RuleAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								RuleResourceName: "rule-a",
							},
							Method: &lb_domains.RuleMethodEntryA{
								lb_domains.RuleMethodPropsCommon{
									ResourceName: "entry-a-1",
									MType:        lb_domains.RuleMethodMTypeEntryA,
									Enabled:      true,
									LiveStatus:   lb_domains.StatusUp,
									ReadyStatus:  lb_domains.StatusDown,
								},
							},
						},
						{
							RuleAttributeMeta: lb_domains.RuleAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								RuleResourceName: "rule-a",
							},
							Method: &lb_domains.RuleMethodEntryAAAA{
								lb_domains.RuleMethodPropsCommon{
									ResourceName: "entry-aaaa-1",
									MType:        lb_domains.RuleMethodMTypeEntryAAAA,
									Enabled:      true,
									LiveStatus:   lb_domains.StatusUp,
									ReadyStatus:  lb_domains.StatusDown,
								},
							},
						},
						{
							RuleAttributeMeta: lb_domains.RuleAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								RuleResourceName: "rule-a",
							},
							Method: &lb_domains.RuleMethodExitSite{
								RuleMethodPropsCommon: lb_domains.RuleMethodPropsCommon{
									ResourceName: "exit-a-1",
									MType:        lb_domains.RuleMethodMTypeExitSite,
									Enabled:      true,
									LiveStatus:   lb_domains.StatusUp,
									ReadyStatus:  lb_domains.StatusDown,
								},
								ParentResourceName: "entry-a-1",
								SiteResourceName:   "site-a",
							},
						},
						{
							RuleAttributeMeta: lb_domains.RuleAttributeMeta{
								AttributeMeta: lb_domains.AttributeMeta{
									LBDomainID: "b0000000000001",
								},
								RuleResourceName: "rule-a",
							},
							Method: &lb_domains.RuleMethodExitSite{
								RuleMethodPropsCommon: lb_domains.RuleMethodPropsCommon{
									ResourceName: "exit-aaaa-1",
									MType:        lb_domains.RuleMethodMTypeExitSite,
									Enabled:      true,
									LiveStatus:   lb_domains.StatusUp,
									ReadyStatus:  lb_domains.StatusDown,
								},
								ParentResourceName: "entry-aaaa-1",
								SiteResourceName:   "site-aaaa",
							},
						},
					},
				},
			},
		}
	})
	Describe("Config", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/config", httpmock.NewBytesResponder(200, configReadJson))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns LBDomain m1", func() {
				BeforeEach(func() {
					c = lb_domains.Config{
						AttributeMeta: lb_domains.AttributeMeta{
							LBDomainID: "b0000000000001",
						},
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("76D7462F50E3417C96898A827E0D1EC1"))
					Expect(c.Sites).To(Equal(s1.Sites))
					Expect(c.Monitorings).To(Equal(s1.Monitorings))
					Expect(c.Rules).To(Equal(s1.Rules))
				})
			})
		})
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPut, "http://localhost/lb_domains/b0000000000001/config", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/lb_domains/b0000000000001/config"]).To(MatchJSON(configUpdateJson))
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
					err = s1.SetPathParams("b0000000000010")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
