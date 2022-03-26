package contracts_test

import (
	"context"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("common_configs", func() {
	var (
		c      contracts.CommonConfig
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.CommonConfig
		slist  contracts.CommonConfigList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.CommonConfig{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                1,
			Name:              "共通設定1",
			ManagedDNSEnabled: types.Enabled,
			Default:           types.Enabled,
			Description:       "common config 1",
		}
		s2 = contracts.CommonConfig{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                2,
			Name:              "共通設定2",
			ManagedDNSEnabled: types.Disabled,
			Default:           types.Disabled,
			Description:       "common config 2",
		}
		slist = contracts.CommonConfigList{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Items: []contracts.CommonConfig{s1, s2},
		}
	})
	Describe("CommonConfig", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "9901FF93F9BB430F921C4541BBF6882E",
					"result": {
						"id": 1,
						"name": "共通設定1",
						"managed_dns_enabled": 1,
						"default": 1,
						"description": "common config 1"
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "C8FEC177EF804512A40EF7FB358037C0",
					"result": {
						"id": 2,
						"name": "共通設定2",
						"managed_dns_enabled": 0,
						"default": 0,
						"description": "common config 2"
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns CommonConfig id1", func() {
				BeforeEach(func() {
					c = contracts.CommonConfig{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
						Id: 1,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("9901FF93F9BB430F921C4541BBF6882E"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns CommonConfig id2", func() {
				BeforeEach(func() {
					c = contracts.CommonConfig{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
						Id: 2,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("C8FEC177EF804512A40EF7FB358037C0"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/contracts/f2/common_configs", httpmock.NewBytesResponder(202, bs1))
				s := contracts.CommonConfig{
					AttributeMeta: contracts.AttributeMeta{
						ContractId: "f2",
					},
					Name:        "create 🍻",
					Description: "🍺🐇",
				}
				reqId, err = cl.Create(context.Background(), &s, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/contracts/f2/common_configs"]).To(MatchJSON(`{
							"name": "create 🍻",
							"description": "🍺🐇"
						}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(202, bs1))
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
				Expect(cl.RequestBody["/contracts/f1/common_configs/1"]).To(MatchJSON(`{
							"name": "共通設定1",
							"description": "common config 1"
						}`))
			})
		})
		Context("Delete", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Delete(context.Background(), &s1)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("can deletel", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
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
					err = s1.SetPathParams("f10", 200)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.GetContractId()).To(Equal("f10"))
				})
				It("can set Id", func() {
					Expect(s1.GetId()).To(Equal(int64(200)))
				})
			})
			When("not enough arguments", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("f10")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("f10", 1, 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2, 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (Id)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("id1", "1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.CommonConfig
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "common_configs")
			testtool.TestGetPathMethodForSpec(&s1, "/contracts/f1/common_configs", "/contracts/f1/common_configs/1")
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
	Describe("CommonConfigList", func() {
		var c contracts.CommonConfigList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "D1203ABF5C6E4A73AD1F184F5F2C1A6B",
					"results": [
						{
							"id": 1,
							"name": "共通設定1",
							"managed_dns_enabled": 1,
							"default": 1,
							"description": "common config 1"
						},
						{
							"id": 2,
							"name": "共通設定2",
							"managed_dns_enabled": 0,
							"default": 0,
							"description": "common config 2"
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.CommonConfigList{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("D1203ABF5C6E4A73AD1F184F5F2C1A6B"))
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
					err = slist.SetPathParams("id12")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(slist.GetContractId()).To(Equal("id12"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = slist.SetPathParams("f1", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractId)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.CommonConfigList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "common_configs")
			testtool.TestGetPathMethodForCountableList(&slist, "/contracts/f1/common_configs")
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
		Context("api.CountableListSpec common test", func() {
			testtool.TestGetMaxLimit(&slist, 10000)
			testtool.TestClearItems(&slist)
			testtool.TestAddItem(&slist, s1)
		})
	})
	Describe("CommonConfigListSearchKeywords", func() {
		Context("GetValues", func() {
			testcase := []struct {
				keyword contracts.CommonConfigListSearchKeywords
				values  url.Values
			}{
				{
					contracts.CommonConfigListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeAND,
							Offset: int32(10),
							Limit:  int32(100),
						},
					},
					url.Values{
						"type":   []string{"AND"},
						"offset": []string{"10"},
						"limit":  []string{"100"},
					},
				},
				{
					contracts.CommonConfigListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeOR,
							Offset: int32(10),
							Limit:  int32(100),
						},
						FullText:    api.KeywordsString{"hogehoge", "🐰"},
						Name:        api.KeywordsString{"hogehogeあ🍺"},
						Description: api.KeywordsString{"あああ", "🍺"},
					},
					url.Values{
						"type":                    []string{"OR"},
						"offset":                  []string{"10"},
						"limit":                   []string{"100"},
						"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
						"_keywords_name[]":        []string{"hogehogeあ🍺"},
						"_keywords_description[]": []string{"あああ", "🍺"},
					},
				},
			}
			It("can convert url.Value", func() {
				for _, tc := range testcase {
					s, err := tc.keyword.GetValues()
					Expect(err).To(Succeed())
					Expect(s).To(Equal(tc.values))
				}
			})
		})
	})
})
