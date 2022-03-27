package zones_test

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("contracts", func() {
	var (
		c     zones.Contract
		cl    *testtool.TestClient
		err   error
		reqId string
		s1    zones.Contract
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = zones.Contract{
			AttributeMeta: zones.AttributeMeta{
				ZoneID: "m1",
			},
			Contract: core.Contract{
				ID:          "f1",
				ServiceCode: "dpf0000001",
				State:       types.StateBeforeStart,
				Favorite:    types.FavoriteHighPriority,
				Plan:        core.PlanBasic,
				Description: "contract 1",
			},
		}
	})
	Describe("Contract", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/contract", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "A4FBAF142456476BBBA919737A4C7663",
					"result": {
						"id": "f1",
						"service_code": "dpf0000001",
						"state": 1,
						"favorite": 1,
						"plan": 1,
						"description": "contract 1"
						}
				}`)))
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns CommonConfig id1", func() {
				BeforeEach(func() {
					c = zones.Contract{
						AttributeMeta: zones.AttributeMeta{
							ZoneID: "m1",
						},
					}
					reqId, err = cl.Read(context.Background(), &c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("A4FBAF142456476BBBA919737A4C7663"))
					Expect(c).To(Equal(s1))
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
					err = s1.SetPathParams("m10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractID", func() {
					Expect(s1.GetZoneID()).To(Equal("m10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("m10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (ZoneID)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *zones.Contract
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "contract")

			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/zones/m1/contract")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
	})
})
