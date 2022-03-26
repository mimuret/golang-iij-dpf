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

var _ = Describe("default_ttl", func() {
	var (
		c      zones.DefaultTTL
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 zones.DefaultTTL
		zdiff  zones.DefaultTTLDiff
		slist  zones.DefaultTTLDiffList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.DefaultTTL{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Value:    3600,
			State:    zones.DefaultTTLStateApplied,
			Operator: "user1",
		}
		s2 = zones.DefaultTTL{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Value:    300,
			State:    zones.DefaultTTLStateApplied,
			Operator: "user2",
		}
		zdiff = zones.DefaultTTLDiff{
			New: &s1,
			Old: &s2,
		}
		slist = zones.DefaultTTLDiffList{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Items: []zones.DefaultTTLDiff{zdiff},
		}
	})
	Describe("DefaultTTL", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/default_ttl", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "19AB6F91095C4E4E83659AD75ADFE0E3",
					"result": {
						"value": 3600,
						"state": 0,
						"operator": "user1"
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns CommonConfig id1", func() {
				BeforeEach(func() {
					c = zones.DefaultTTL{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("19AB6F91095C4E4E83659AD75ADFE0E3"))
					Expect(c).To(Equal(s1))
				})
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/default_ttl", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/zones/m1/default_ttl"]).To(MatchJSON(`{
							"value": 3600
						}`))
			})
		})
		Context("Cancel", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/default_ttl/changes", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Cancel(context.Background(), &s1)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("body is empty", func() {
				Expect(cl.RequestBody["/zones/m1/default_ttl"]).To(Equal(""))
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
				It("can set ContractId", func() {
					Expect(s1.GetZoneId()).To(Equal("m10"))
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
			When("arguments type missmatch (ZoneId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.DefaultTTL
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "default_ttl")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/default_ttl")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/zones/m1/default_ttl")
				})
				When("action is ActionCancel", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCancel, http.MethodDelete, "/zones/m1/default_ttl/changes")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
	Describe("DefaultTTLDiffList", func() {
		var c zones.DefaultTTLDiffList
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/default_ttl/diffs", httpmock.NewBytesResponder(200, []byte(`{
						"request_id": "6FB71A568C66435BB1B7DCBC6911D881",
						"results": [
							{
								"new": {
									"value": 3600,
									"state": 0,
									"operator": "user1"		
								},
								"old": {
									"value": 300,
									"state": 0,
									"operator": "user2"
								}
						}]
					}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns diff", func() {
				BeforeEach(func() {
					c = zones.DefaultTTLDiffList{
						AttributeMeta: zones.AttributeMeta{
							ZoneId: "m1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("6FB71A568C66435BB1B7DCBC6911D881"))
					Expect(c).To(Equal(slist))
				})
			})
		})
		Context("SetPathParams", func() {
			BeforeEach(func() {
				err = slist.SetPathParams()
			})
			It("nothing todo", func() {
				Expect(err).To(Succeed())
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.DefaultTTLDiffList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "default_ttl/diffs")

			testtool.TestGetPathMethodForList(&slist, "/zones/m1/default_ttl/diffs")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 1)
			Context("Index", func() {
				When("index exist", func() {
					It("returns item", func() {
						Expect(slist.Index(0)).To(Equal(zdiff))
					})
				})
			})
		})
	})
	Context("DefaultTTLState", func() {
		Context("String", func() {
			Expect(zones.DefaultTTLStateApplied.String()).To(Equal("Applied"))
			Expect(zones.DefaultTTLStateToBeUpdate.String()).To(Equal("ToBeUpdate"))
			Expect(zones.DefaultTTLStateBeforeUpdate.String()).To(Equal("BeforeUpdate"))
		})
	})
	Context("DefaultTTLDiff", func() {
		var s, copy, nilSpec *zones.DefaultTTLDiff
		BeforeEach(func() {
			s = &zones.DefaultTTLDiff{}
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
