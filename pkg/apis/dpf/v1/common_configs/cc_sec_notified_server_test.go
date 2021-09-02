package common_configs_test

import (
	"net"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("cc_sec_notified_servers", func() {
	var (
		c      common_configs.CcSecNotifiedServer
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 common_configs.CcSecNotifiedServer
		slist  common_configs.CcSecNotifiedServerList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = common_configs.CcSecNotifiedServer{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id:      1,
			Address: net.ParseIP("192.168.0.1"),
			TsigId:  0,
		}
		s2 = common_configs.CcSecNotifiedServer{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id:      2,
			Address: net.ParseIP("2001:db8::1"),
			TsigId:  2,
		}
		slist = common_configs.CcSecNotifiedServerList{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Items: []common_configs.CcSecNotifiedServer{s1, s2},
		}
	})
	Describe("CcSecNotifiedServer", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "5D7FECB0A56C49C2B73BD45DD174A575",
					"result": {
						"id": 1,
						"address": "192.168.0.1",
						"tsig_id": 0
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "5ABBBEF090DF4F208AD9EDEAC7D702A5",
					"result": {
						"id": 2,
						"address": "2001:db8::1",
						"tsig_id": 2
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns IPv4 server", func() {
				BeforeEach(func() {
					c = common_configs.CcSecNotifiedServer{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
						Id: 1,
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("5D7FECB0A56C49C2B73BD45DD174A575"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns IPv6 server", func() {
				BeforeEach(func() {
					c = common_configs.CcSecNotifiedServer{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
						Id: 2,
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("5ABBBEF090DF4F208AD9EDEAC7D702A5"))
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
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/2/cc_sec_notified_servers", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/3/cc_sec_notified_servers", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("ipv4 host", func() {
				BeforeEach(func() {
					s := common_configs.CcSecNotifiedServer{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 2,
						},
						Address: net.ParseIP("192.168.10.1"),
						TsigId:  1,
					}
					reqId, err = cl.Create(&s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/2/cc_sec_notified_servers"]).To(MatchJSON(`{
							"address": "192.168.10.1",
							"tsig_id": 1
						}`))
				})
			})
			When("ipv6 host", func() {
				BeforeEach(func() {
					s := common_configs.CcSecNotifiedServer{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 3,
						},
						Address: net.ParseIP("2001:db8::1"),
					}
					reqId, err = cl.Create(&s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/3/cc_sec_notified_servers"]).To(MatchJSON(`{
							"address": "2001:db8::1"
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
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("ipv4 host", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_sec_notified_servers/1"]).To(MatchJSON(`{
							"address": "192.168.0.1"
						}`))
				})
			})
			When("ipv6 host", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_sec_notified_servers/2"]).To(MatchJSON(`{
							"address": "2001:db8::1",
							"tsig_id": 2
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
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns IPv4 server", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(&s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("returns IPv6 server", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(&s2)
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
				It("can set CommonConfigId", func() {
					Expect(s1.GetCommonConfigId()).To(Equal(int64(100)))
				})
				It("can set Id", func() {
					Expect(s1.GetId()).To(Equal(int64(200)))
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
			When("arguments type missmatch (CommonConfigId)", func() {
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
			var nilSpec *common_configs.CcSecNotifiedServer
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "cc_sec_notified_servers")
			testtool.TestGetPathMethodForSpec(&s1, "/common_configs/1/cc_sec_notified_servers", "/common_configs/1/cc_sec_notified_servers/1")
		})
		Context("contracts.ChildSpec common test", func() {
			Context("GetId", func() {
				It("returns Id", func() {
					Expect(s1.GetId()).To(Equal(s1.Id))
				})
			})
			Context("SetId", func() {
				BeforeEach(func() {
					s1.SetId(2)
				})
				It("can set Id", func() {
					Expect(s1.GetId()).To(Equal(int64(2)))
				})
			})
		})
	})
	Describe("CcSecNotifiedServerList", func() {
		var (
			c common_configs.CcSecNotifiedServerList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "450B94B486D34D548313707FC7590CE0",
					"results": [
						{
							"id": 1,
							"address": "192.168.0.1",
							"tsig_id": 0
						},
						{
							"id": 2,
							"address": "2001:db8::1",
							"tsig_id": 2
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = common_configs.CcSecNotifiedServerList{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
					}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("450B94B486D34D548313707FC7590CE0"))
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
				It("can set CommonConfigId", func() {
					Expect(slist.GetCommonConfigId()).To(Equal(int64(99)))
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
			When("arguments type missmatch (CommonConfigId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("2")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *common_configs.CcSecNotifiedServerList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "cc_sec_notified_servers")
			testtool.TestGetPathMethodForList(&slist, "/common_configs/1/cc_sec_notified_servers")
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
