package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("common_configs/default", func() {
	var (
		cl    *testtool.TestClient
		err   error
		reqId string
		s1    contracts.CommonConfigDefault
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.CommonConfigDefault{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id: 10,
		}
	})
	Describe("CommonConfigDefault", func() {
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/default", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Apply(context.Background(), &s1, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("sends json", func() {
				Expect(cl.RequestBody["/contracts/f1/common_configs/default"]).To(MatchJSON(`{
							"common_config_id": 10
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
					err = s1.SetPathParams("f10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.GetContractId()).To(Equal("f10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("f10", 1)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ContractId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *contracts.CommonConfigDefault
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "common_configs/default")
			Context("GetPathMethod", func() {
				testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/contracts/f1/common_configs/default")
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/contracts/f1/common_configs/default")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, "", "")
				})
			})
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
})
