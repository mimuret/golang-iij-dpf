package common_configs_test

import (
	"context"
	"net"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("cc_primaries", func() {
	var (
		c      common_configs.CcPrimary
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 common_configs.CcPrimary
		slist  common_configs.CcPrimaryList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = common_configs.CcPrimary{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			ID:      1,
			Address: net.ParseIP("192.168.0.1"),
			TsigID:  0,
			Enabled: types.Disabled,
		}
		s2 = common_configs.CcPrimary{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			ID:      2,
			Address: net.ParseIP("2001:db8::1"),
			TsigID:  2,
			Enabled: types.Enabled,
		}
		slist = common_configs.CcPrimaryList{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			Items: []common_configs.CcPrimary{s1, s2},
		}
	})
	Describe("CCPrimary", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "73A87F978AFD42D8B0AC810218AEC7CC",
					"result": {
						"id": 1,
						"address": "192.168.0.1",
						"tsig_id": null,
						"enabled": 0
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EBCA51F1CCD8457FB242766B59B5786E",
					"result": {
						"id": 2,
						"address": "2001:db8::1",
						"tsig_id": 2,
						"enabled": 1
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns IPv4 server", func() {
				BeforeEach(func() {
					c = common_configs.CcPrimary{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 1,
						},
						ID: 1,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("73A87F978AFD42D8B0AC810218AEC7CC"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns IPv6 server", func() {
				BeforeEach(func() {
					c = common_configs.CcPrimary{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 1,
						},
						ID: 2,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("EBCA51F1CCD8457FB242766B59B5786E"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/2/cc_primaries", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/3/cc_primaries", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("ipv4 host", func() {
				BeforeEach(func() {
					s := common_configs.CcPrimary{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 2,
						},
						Address: net.ParseIP("192.168.10.1"),
						TsigID:  1,
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/2/cc_primaries"]).To(MatchJSON(`{
							"address": "192.168.10.1",
							"tsig_id": 1
						}`))
				})
			})
			When("ipv6 host", func() {
				BeforeEach(func() {
					s := common_configs.CcPrimary{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 3,
						},
						Address: net.ParseIP("2001:db8::1"),
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/3/cc_primaries"]).To(MatchJSON(`{
							"address": "2001:db8::1",
							"tsig_id": null
						}`))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("ipv4 host", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_primaries/1"]).To(MatchJSON(`{
							"address": "192.168.0.1",
							"enabled": 0,
							"tsig_id": null
						}`))
				})
			})
			When("ipv6 host", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_primaries/2"]).To(MatchJSON(`{
							"address": "2001:db8::1",
							"tsig_id": 2,
							"enabled": 1
						}`))
				})
			})
		})
		Context("Delete", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns IPv4 server", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(context.Background(), &s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("returns IPv6 server", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(context.Background(), &s2)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
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
					err = s1.SetPathParams(100, 200)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set CommonConfigID", func() {
					Expect(s1.GetCommonConfigID()).To(Equal(int64(100)))
				})
				It("can set Id", func() {
					Expect(s1.GetID()).To(Equal(int64(200)))
				})
			})
			When("not enough arguments", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(100)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(100, 1, 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (CommonConfigID)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("2", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (Id)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(100, "1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *common_configs.CcPrimary
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "cc_primaries")
			testtool.TestGetPathMethodForSpec(&s1, "/common_configs/1/cc_primaries", "/common_configs/1/cc_primaries/1")
		})
		Context("contracts.ChildSpec common test", func() {
			Context("GetID", func() {
				It("returns Id", func() {
					Expect(s1.GetID()).To(Equal(s1.ID))
				})
			})
			Context("SetID", func() {
				BeforeEach(func() {
					s1.SetID(2)
				})
				It("can set Id", func() {
					Expect(s1.GetID()).To(Equal(int64(2)))
				})
			})
		})
	})
	Describe("CCPrimaryList", func() {
		var c common_configs.CcPrimaryList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "1194A148B9B34C5F81977EF17A39CD4C",
					"results": [
						{
							"id": 1,
							"address": "192.168.0.1",
							"tsig_id": null,
							"enabled": 0
						},
						{
							"id": 2,
							"address": "2001:db8::1",
							"tsig_id": 2,
							"enabled": 1
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = common_configs.CcPrimaryList{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 1,
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("1194A148B9B34C5F81977EF17A39CD4C"))
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
					err = slist.SetPathParams(99)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set CommonConfigID", func() {
					Expect(slist.GetCommonConfigID()).To(Equal(int64(99)))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(100, 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (CommonConfigID)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("2")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *common_configs.CcPrimaryList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "cc_primaries")
			testtool.TestGetPathMethodForList(&slist, "/common_configs/1/cc_primaries")
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
