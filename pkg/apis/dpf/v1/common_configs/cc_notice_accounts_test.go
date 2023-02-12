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

var _ = Describe("cc_notice_accounts", func() {
	var (
		c      common_configs.CcNoticeAccount
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 common_configs.CcNoticeAccount
		slist  common_configs.CcNoticeAccountList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = common_configs.CcNoticeAccount{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			ResourceName: "user1",
			Name:         "user1",
			Lang:         common_configs.CcNoticeLangJA,
			Props: common_configs.CcNoticeProps{
				Mail: "mail@example.jp",
				Phone: &common_configs.CcNoticePhone{
					CountryCode: "81",
					Number:      "300000000",
				},
			},
		}
		s2 = common_configs.CcNoticeAccount{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			ResourceName: "user2",
			Name:         "user2",
			Lang:         common_configs.CcNoticeLangENUS,
			Props: common_configs.CcNoticeProps{
				Mail: "mail2@example.com",
			},
		}
		slist = common_configs.CcNoticeAccountList{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigID: 1,
			},
			Items: []common_configs.CcNoticeAccount{s1, s2},
		}
	})
	Describe("CcNoticeAccount", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_notice_accounts/user1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "73A87F978AFD42D8B0AC810218AEC7CC",
					"result": {
						"resource_name": "user1",
						"name": "user1",
						"lang": "ja",
						"props": {
							"mail": "mail@example.jp",
							"phone": {
								"country_code": "81",
								"number": "300000000"
							}
						}
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_notice_accounts/user2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EBCA51F1CCD8457FB242766B59B5786E",
					"result": {
						"resource_name": "user2",
						"name": "user2",
						"lang": "en_US",
						"props": {
							"mail": "mail2@example.com"
						}
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns phone and mail", func() {
				BeforeEach(func() {
					c = common_configs.CcNoticeAccount{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 1,
						},
						ResourceName: "user1",
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("73A87F978AFD42D8B0AC810218AEC7CC"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns mail only", func() {
				BeforeEach(func() {
					c = common_configs.CcNoticeAccount{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 1,
						},
						ResourceName: "user2",
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
				id3, bs3 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/10/cc_notice_accounts", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/20/cc_notice_accounts", httpmock.NewBytesResponder(202, bs2))
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/30/cc_notice_accounts", httpmock.NewBytesResponder(202, bs3))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("phone and mail", func() {
				BeforeEach(func() {
					s := common_configs.CcNoticeAccount{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 10,
						},
						ResourceName: "user10",
						Name:         "user10",
						Lang:         common_configs.CcNoticeLangJA,
						Props: common_configs.CcNoticeProps{
							Mail: "mail10@example.jp",
						},
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/10/cc_notice_accounts"]).To(MatchJSON(`{
						"resource_name": "user10",
						"name": "user10",
						"lang": "ja",
							"props": {
								"mail": "mail10@example.jp"
							}
						}`))
				})
			})
			When("phone only", func() {
				BeforeEach(func() {
					s := common_configs.CcNoticeAccount{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 20,
						},
						Name: "user20",
						Lang: common_configs.CcNoticeLangJA,
						Props: common_configs.CcNoticeProps{
							Phone: &common_configs.CcNoticePhone{
								CountryCode: "81",
								Number:      "300000000",
							},
						},
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/20/cc_notice_accounts"]).To(MatchJSON(`{
						"name": "user20",
						"lang": "ja",
						"props": {
							"phone": {
								"country_code": "81",
								"number": "300000000"
							}
						}
					}`))
				})
			})
			When("phone and mail", func() {
				BeforeEach(func() {
					s := common_configs.CcNoticeAccount{
						AttributeMeta: common_configs.AttributeMeta{
							CommonConfigID: 30,
						},
						Name: "user30",
						Lang: common_configs.CcNoticeLangJA,
						Props: common_configs.CcNoticeProps{
							Mail: "mail5@example.jp",
							Phone: &common_configs.CcNoticePhone{
								CountryCode: "81",
								Number:      "300000000",
							},
						},
					}
					reqId, err = cl.Create(context.Background(), &s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id3))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/common_configs/30/cc_notice_accounts"]).To(MatchJSON(`{
						"name": "user30",
						"lang": "ja",
						"props": {
							"mail": "mail5@example.jp",
							"phone": {
								"country_code": "81",
								"number": "300000000"
							}
						}
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
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_notice_accounts/user1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_notice_accounts/user2", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("update mail and phone", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("patch json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_notice_accounts/user1"]).To(MatchJSON(`{
						"name": "user1",
						"lang": "ja",
							"props": {
								"mail": "mail@example.jp",
								"phone": {
									"country_code": "81",
									"number": "300000000"
								}
							}
						}`))
				})
			})
			When("update mail only", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(context.Background(), &s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("patch json", func() {
					Expect(cl.RequestBody["/common_configs/1/cc_notice_accounts/user2"]).To(MatchJSON(`{
						"name": "user2",
						"lang": "en_US",
						"props": {
							"mail": "mail2@example.com"
						}
					}`))
				})
			})
		})
		Context("Delete", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_notice_accounts/user1", httpmock.NewBytesResponder(202, bs1))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("delete", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(context.Background(), &s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
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
					err = s1.SetPathParams(100, "user")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set CommonConfigID", func() {
					Expect(s1.GetCommonConfigID()).To(Equal(int64(100)))
				})
				It("can set ResourceName", func() {
					Expect(s1.ResourceName).To(Equal("user"))
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
					err = s1.SetPathParams(100, "user", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (CommonConfigID)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("2", "user")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (Id)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(100, 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *common_configs.CcPrimary
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "cc_notice_accounts")
			testtool.TestGetPathMethodForSpec(&s1, "/common_configs/1/cc_notice_accounts", "/common_configs/1/cc_notice_accounts/user1")
		})
	})
	Describe("CcNoticeAccountList", func() {
		var c common_configs.CcNoticeAccountList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_notice_accounts", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "1194A148B9B34C5F81977EF17A39CD4C",
					"results": [
						{
							"resource_name": "user1",
							"name": "user1",
							"lang": "ja",
							"props": {
								"mail": "mail@example.jp",
								"phone": {
									"country_code": "81",
									"number": "300000000"
								}
							}
						},
						{
							"resource_name": "user2",
							"name": "user2",
							"lang": "en_US",
							"props": {
								"mail": "mail2@example.com"
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
					c = common_configs.CcNoticeAccountList{
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
			testtool.TestGetName(&slist, "cc_notice_accounts")
			testtool.TestGetPathMethodForList(&slist, "/common_configs/1/cc_notice_accounts")
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
