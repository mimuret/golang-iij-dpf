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

var _ = Describe("RuleMethod", func() {
	var (
		c          lb_domains.RuleMethod
		cl         *testtool.TestClient
		err        error
		reqId      string
		s1, s2     lb_domains.RuleMethod
		priorityS1 uint
		slist      lb_domains.RuleMethodList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		priorityS1 = 10
		s1 = lb_domains.RuleMethod{
			RuleAttributeMeta: lb_domains.RuleAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				RuleResourceName: "rule-1",
			},
			Priority: &priorityS1,
			Method: &lb_domains.RuleMethodEntryA{
				lb_domains.RuleMethodPropsCommon{
					ResourceName: "method-1",
					MType:        lb_domains.RuleMethodMTypeEntryA,
					Enabled:      true,
					LiveStatus:   lb_domains.StatusUp,
					ReadyStatus:  lb_domains.StatusDown,
				},
			},
		}
		s2 = lb_domains.RuleMethod{
			RuleAttributeMeta: lb_domains.RuleAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				RuleResourceName: "rule-1",
			},
			Method: &lb_domains.RuleMethodEntryA{
				lb_domains.RuleMethodPropsCommon{
					ResourceName: "method-2",
					MType:        lb_domains.RuleMethodMTypeEntryA,
					Enabled:      true,
					LiveStatus:   lb_domains.StatusUp,
					ReadyStatus:  lb_domains.StatusUp,
				},
			},
		}
		slist = lb_domains.RuleMethodList{
			RuleAttributeMeta: lb_domains.RuleAttributeMeta{
				AttributeMeta: lb_domains.AttributeMeta{
					LBDomainID: "b0000000000001",
				},
				RuleResourceName: "rule-1",
			},
			Items: []lb_domains.RuleMethod{s1, s2},
		}
	})
	Describe("RuleMethod", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "76D7462F50E3417C96898A827E0D1EC1",
					"result": {
							"priority": 10,
							"method": {
								"resource_name": "method-1",
								"mtype": "entry_a",
								"enabled": true,
								"live_status": "up",
								"ready_status": "down"
							}
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
					"result": {
						"method": {
							"resource_name": "method-2",
							"mtype": "entry_a",
							"enabled": true,
							"live_status": "up",
							"ready_status": "up"
						}
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns LBDomain m1", func() {
				BeforeEach(func() {
					c = lb_domains.RuleMethod{
						RuleAttributeMeta: lb_domains.RuleAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							RuleResourceName: "rule-1",
						},
						Method: &lb_domains.RuleMethodPropsCommon{
							ResourceName: "method-1",
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
					c = lb_domains.RuleMethod{
						RuleAttributeMeta: lb_domains.RuleAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							RuleResourceName: "rule-1",
						},
						Method: &lb_domains.RuleMethodPropsCommon{
							ResourceName: "method-2",
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
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/lb_domains/b0000000000001/rules/rule-1/rule_methods", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Create(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("patch json", func() {
				Expect(cl.RequestBody["/lb_domains/b0000000000001/rules/rule-1/rule_methods"]).To(MatchJSON(`{
					"priority": 10,
					"method": {
						"resource_name": "method-1",
						"mtype": "entry_a",
						"enabled": true
					}
				}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Update(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("patch json", func() {
				Expect(cl.RequestBody["/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1"]).To(MatchJSON(`{
					"priority": 10,
					"method": {
						"enabled": true
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
					err = s1.SetPathParams("b0000000000010", "rule-10", "method-10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(s1.LBDomainID).To(Equal("b0000000000010"))
					Expect(s1.RuleResourceName).To(Equal("rule-10"))
					Expect(s1.GetMethodResourceName()).To(Equal("method-10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "rule-10", "method-10", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("b0000000000010", "rule-10", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *lb_domains.RuleMethod
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "rule_methods")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/lb_domains/b0000000000001/rules/rule-1/rule_methods/method-1")
				})
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/lb_domains/b0000000000001/rules/rule-1/rule_methods")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("RuleMethodList", func() {
		var c lb_domains.RuleMethodList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/lb_domains/b0000000000001/rules/rule-1/rule_methods", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "2C18922432DC48D485613F5383A7ED8E",
					"results": [
						{
							"priority": 10,
							"method": {
								"resource_name": "method-1",
								"mtype": "entry_a",
								"enabled": true,
								"live_status": "up",
								"ready_status": "down"
							}
						},
						{
							"method": {
								"resource_name": "method-2",
								"mtype": "entry_a",
								"enabled": true,
								"live_status": "up",
								"ready_status": "up"
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
					c = lb_domains.RuleMethodList{
						RuleAttributeMeta: lb_domains.RuleAttributeMeta{
							AttributeMeta: lb_domains.AttributeMeta{
								LBDomainID: "b0000000000001",
							},
							RuleResourceName: "rule-1",
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
					err = slist.SetPathParams("b0000000000010", "rule-10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set LBDomainID", func() {
					Expect(slist.LBDomainID).To(Equal("b0000000000010"))
					Expect(slist.RuleResourceName).To(Equal("rule-10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("b0000000000010", "rule-1", 2)
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
			var nilSpec *lb_domains.RuleMethodList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "rule_methods")
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
