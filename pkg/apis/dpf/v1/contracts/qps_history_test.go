package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("qps_histories", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.QpsHistory
		slist  contracts.QpsHistoryList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.QpsHistory{
			ServiceCode: "f00001",
			Name:        "hoge",
			Values: []contracts.QpsValue{
				{
					Month: "2021-01",
					Qps:   100,
				},
				{
					Month: "2021-02",
					Qps:   100,
				},
			},
		}
		s2 = contracts.QpsHistory{
			ServiceCode: "f00002",
			Name:        "hogehoge",
			Values: []contracts.QpsValue{
				{
					Month: "2021-07",
					Qps:   20,
				},
				{
					Month: "2021-09",
					Qps:   100,
				},
			},
		}

		slist = contracts.QpsHistoryList{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Items: []contracts.QpsHistory{s1, s2},
		}
	})
	Describe("QpsHistoryList", func() {
		var (
			c contracts.QpsHistoryList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/qps/histories", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "3649092EAD70427A854CE933973AFFF8",
					"results": [
						{
							"service_code": "f00001",
							"name": "hoge",
							"values": [
								{
									"month": "2021-01",
									"qps": 100
								},
								{
									"month": "2021-02",
									"qps": 100
								}
							]
						},
						{
							"service_code": "f00002",
							"name": "hogehoge",
							"values": [
								{
									"month": "2021-07",
									"qps": 20
								},
								{
									"month": "2021-09",
									"qps": 100
								}
							]
						}
					]
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = contracts.QpsHistoryList{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
					}
					reqId, err = cl.List(context.Background(), &c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("3649092EAD70427A854CE933973AFFF8"))
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
			var nilSpec *contracts.QpsHistoryList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "qps/histories")
			testtool.TestGetPathMethodForList(&slist, "/contracts/f1/qps/histories")
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
	Describe("QpsHistory", func() {
		var (
			s, copy, nilSpec *contracts.QpsHistory
		)
		BeforeEach(func() {
			s = &contracts.QpsHistory{}
		})
		Context("DeepCopyt", func() {
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
	Describe("QpsValue", func() {
		var (
			s, copy, nilSpec *contracts.QpsValue
		)
		BeforeEach(func() {
			s = &contracts.QpsValue{}
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
