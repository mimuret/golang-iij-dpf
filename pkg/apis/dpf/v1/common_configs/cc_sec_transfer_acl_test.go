package common_configs_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("cc_sec_transfer_acls", func() {
	var (
		c      common_configs.CcSecTransferAcl
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 common_configs.CcSecTransferAcl
		slist  common_configs.CcSecTransferAclList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = common_configs.CcSecTransferAcl{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id:      1,
			Network: testtool.MustParseIPNet("192.168.1.0/24"),
			TsigId:  0,
		}
		s2 = common_configs.CcSecTransferAcl{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id:      2,
			Network: testtool.MustParseIPNet("2001:db8::/64"),
			TsigId:  2,
		}
		slist = common_configs.CcSecTransferAclList{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Items: []common_configs.CcSecTransferAcl{s1, s2},
		}
	})
	Describe("CcSecTransferAcl", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "C5E6406DCBFC4931A0EE480ABA9FA505",
					"result": {
						"id": 1,
						"network": "192.168.1.0/24",
						"tsig_id": 0
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "3D82ADCA5479408094DEF876A679BFD3",
					"result": {
						"id": 2,
						"network": "2001:db8::/64",
						"tsig_id": 2
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns IPv4 server", func() {
				BeforeEach(func() {
					c = common_configs.CcSecTransferAcl{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
						Id: 1,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("C5E6406DCBFC4931A0EE480ABA9FA505"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns IPv6 server", func() {
				BeforeEach(func() {
					c = common_configs.CcSecTransferAcl{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
						Id: 2,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("3D82ADCA5479408094DEF876A679BFD3"))
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
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/2/cc_sec_transfer_acls", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/3/cc_sec_transfer_acls", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("ipv4 host", func() {
				BeforeEach(func() {
					s := common_configs.CcSecTransferAcl{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 2,
						},
						Network: testtool.MustParseIPNet("10.0.0.0/8"),
						TsigId:  1,
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/2/cc_sec_transfer_acls"]).To(MatchJSON(`{
							"network": "10.0.0.0/8",
							"tsig_id": 1
						}`))
				})
			})
			When("ipv6 host", func() {
				BeforeEach(func() {
					s := common_configs.CcSecTransferAcl{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 3,
						},
						Network: testtool.MustParseIPNet("2001:db8::/48"),
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/3/cc_sec_transfer_acls"]).To(MatchJSON(`{
							"network": "2001:db8::/48"
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
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(202, bs2))
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
					Expect(cl.RequestBody["/common_configs/1/cc_sec_transfer_acls/1"]).To(MatchJSON(`{
							"network": "192.168.1.0/24"
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
					Expect(cl.RequestBody["/common_configs/1/cc_sec_transfer_acls/2"]).To(MatchJSON(`{
							"network": "2001:db8::/64",
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
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(202, bs2))
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
			var nilSpec *common_configs.CcSecTransferAcl
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "cc_sec_transfer_acls")
			testtool.TestGetPathMethodForSpec(&s1, "/common_configs/1/cc_sec_transfer_acls", "/common_configs/1/cc_sec_transfer_acls/1")
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
	Describe("CcSecTransferAclList", func() {
		var (
			c common_configs.CcSecTransferAclList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "13032F0785DA47768B3E36A1A854A808",
					"results": [
						{
							"id": 1,
							"network": "192.168.1.0/24",
							"tsig_id": 0
							},
						{
							"id": 2,
							"network": "2001:db8::/64",
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
					c = common_configs.CcSecTransferAclList{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigId: 1,
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("13032F0785DA47768B3E36A1A854A808"))
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
			var nilSpec *common_configs.CcSecTransferAclList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "cc_sec_transfer_acls")
			testtool.TestGetPathMethodForList(&slist, "/common_configs/1/cc_sec_transfer_acls")
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
