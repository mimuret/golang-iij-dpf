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
)

var _ = Describe("tsigs", func() {
	var (
		c      contracts.Tsig
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.Tsig
		slist  contracts.TsigList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.Tsig{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			Id:          1,
			Name:        "hogehoge1",
			Algorithm:   contracts.TsigAlgorithmHMACSHA256,
			Secret:      "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
			Description: "tsig 1",
		}
		s2 = contracts.Tsig{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			Id:          2,
			Name:        "hogehoge2",
			Algorithm:   contracts.TsigAlgorithmHMACSHA256,
			Secret:      "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
			Description: "tsig 2",
		}
		slist = contracts.TsigList{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			Items: []contracts.Tsig{s1, s2},
		}
	})
	Describe("Tsig", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "96043C4DF2EE4EEF9094FF566F87505A",
					"result": {
						"id": 1,
						"name": "hogehoge1",
						"algorithm": 0,
						"secret": "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
						"description": "tsig 1"
						}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "E16A52E5C5DD4EB7B4A60F99E8FACE95",
					"result": {
						"id": 2,
						"name": "hogehoge2",
						"algorithm": 0,
						"secret": "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
						"description": "tsig 2"
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns Tsig id1", func() {
				BeforeEach(func() {
					c = contracts.Tsig{
						AttributeMeta: contracts.AttributeMeta{
							ContractID: "f1",
						},
						Id: 1,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("96043C4DF2EE4EEF9094FF566F87505A"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns Tsig id2", func() {
				BeforeEach(func() {
					c = contracts.Tsig{
						AttributeMeta: contracts.AttributeMeta{
							ContractID: "f1",
						},
						Id: 2,
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("E16A52E5C5DD4EB7B4A60F99E8FACE95"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/contracts/f2/tsigs", httpmock.NewBytesResponder(202, bs1))
				s := contracts.Tsig{
					AttributeMeta: contracts.AttributeMeta{
						ContractID: "f2",
					},
					Name:        "hoge-huga",
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
				Expect(cl.RequestBody["/contracts/f2/tsigs"]).To(MatchJSON(`{
						"name": "hoge-huga",
						"description": "🍺🐇"
					}`))
			})
		})
		Context("Update", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Update(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("sends json", func() {
				Expect(cl.RequestBody["/contracts/f1/tsigs/1"]).To(MatchJSON(`{
							"description": "tsig 1"
						}`))
			})
		})
		Context("Delete", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Delete(context.Background(), &s1)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("can delete", func() {
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
				It("can set ContractID", func() {
					Expect(s1.GetContractID()).To(Equal("f10"))
				})
				It("can set Id", func() {
					Expect(s1.GetID()).To(Equal(int64(200)))
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
			When("arguments type missmatch (ContractID)", func() {
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
			var nilSpec *contracts.Tsig
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "tsigs")

			testtool.TestGetPathMethodForSpec(&s1, "/contracts/f1/tsigs", "/contracts/f1/tsigs/1")
		})
		Context("contracts.ChildSpec common test", func() {
			Context("GetID", func() {
				It("returns Id", func() {
					Expect(s1.GetID()).To(Equal(s1.Id))
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
	Describe("TsigList", func() {
		var c contracts.TsigList
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "062ED092113E4509BE4E2E4EF6C6EB2B",
					"results": [
						{
							"id": 1,
							"name": "hogehoge1",
							"algorithm": 0,
							"secret": "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
							"description": "tsig 1"
						},
						{
							"id": 2,
							"name": "hogehoge2",
							"algorithm": 0,
							"secret": "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
							"description": "tsig 2"
						}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "F60D8455D0C7456286565E4D28A35951",
					"result": {
						"count": 2
					}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.TsigList{
						AttributeMeta: contracts.AttributeMeta{
							ContractID: "f1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("062ED092113E4509BE4E2E4EF6C6EB2B"))
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
				It("can set ContractID", func() {
					Expect(slist.GetContractID()).To(Equal("id12"))
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
			When("arguments type missmatch (ContractID)", func() {
				BeforeEach(func() {
					err = slist.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.TsigList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "tsigs")

			testtool.TestGetPathMethodForCountableList(&slist, "/contracts/f1/tsigs")
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
	Describe("TsigListSearchKeywords", func() {
		Context("GetValues", func() {
			testcase := []struct {
				keyword contracts.TsigListSearchKeywords
				values  url.Values
			}{
				{
					contracts.TsigListSearchKeywords{
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
					contracts.TsigListSearchKeywords{
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
	Describe("TsigAlgorithm", func() {
		var a contracts.TsigAlgorithm
		Context("String", func() {
			When("TsigAlgorithmHMACSHA256", func() {
				BeforeEach(func() {
					a = contracts.TsigAlgorithmHMACSHA256
				})
				It("returns HMAC-SHA256", func() {
					Expect(a.String()).To(Equal("HMAC-SHA256"))
				})
			})
		})
	})
})
