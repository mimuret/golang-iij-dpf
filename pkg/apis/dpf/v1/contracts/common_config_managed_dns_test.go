package contracts_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("common_configs/managed_dns", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 contracts.CommonConfigManagedDns
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = contracts.CommonConfigManagedDns{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			ID:                1,
			ManagedDnsEnabled: types.Enabled,
		}
		s2 = contracts.CommonConfigManagedDns{
			AttributeMeta: contracts.AttributeMeta{
				ContractID: "f1",
			},
			ID:                2,
			ManagedDnsEnabled: types.Disabled,
		}
	})
	Describe("CommonConfigManagedDns", func() {
		Context("Apply", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/1/managed_dns", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/2/managed_dns", httpmock.NewBytesResponder(202, bs2))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("enable", func() {
				BeforeEach(func() {
					reqId, err = cl.Apply(context.Background(), &s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("apply json", func() {
					Expect(cl.RequestBody["/contracts/f1/common_configs/1/managed_dns"]).To(MatchJSON(`{
							"managed_dns_enabled": 1
						}`))
				})
			})
			When("disable", func() {
				BeforeEach(func() {
					reqId, err = cl.Apply(context.Background(), &s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("apply json", func() {
					Expect(cl.RequestBody["/contracts/f1/common_configs/2/managed_dns"]).To(MatchJSON(`{
							"managed_dns_enabled": 0
						}`))
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
					err = s1.SetPathParams("f10", 200)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
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
			var nilSpec *contracts.CommonConfigManagedDns
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "common_configs/1/managed_dns")
			Context("GetPathMethod", func() {
				testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/contracts/f1/common_configs/1/managed_dns")
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/contracts/f1/common_configs/1/managed_dns")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, "", "")
				})
			})
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
})
