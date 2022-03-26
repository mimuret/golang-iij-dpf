package core_test

import (
	"context"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("delegations", func() {
	var (
		cl    *testtool.TestClient
		err   error
		reqId string
		s     core.DelegationApply
	)
	BeforeEach(func() {
		Expect(err).To(Succeed())
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s = core.DelegationApply{
			ZoneIDs: []string{"m1", "m2"},
		}
	})
	Describe("DelegationApply", func() {
		Context("Apply", func() {
			id1, bs1 := testtool.CreateAsyncResponse()
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/delegations", httpmock.NewBytesResponder(202, bs1))
				reqId, err = cl.Apply(context.Background(), &s, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("returns job_id", func() {
				Expect(err).To(Succeed())
				Expect(reqId).To(Equal(id1))
			})
			It("post json", func() {
				Expect(cl.RequestBody["/delegations"]).To(MatchJSON(`{
							"zone_ids": ["m1","m2"]
						}`))
			})
		})
		Context("SetPathParams", func() {
			BeforeEach(func() {
				err = s.SetPathParams()
			})
			It("nothing todo", func() {
				Expect(err).To(Succeed())
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.DelegationApply
			testtool.TestDeepCopyObject(&s, nilSpec)
			testtool.TestGetName(&s, "delegations")

			Context("GetPathMethod", func() {
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s, api.ActionApply, http.MethodPost, "/delegations")
				})
				When("other", func() {
					testtool.TestGetPathMethod(&s, api.ActionRead, "", "")
				})
			})
		})
	})
})
