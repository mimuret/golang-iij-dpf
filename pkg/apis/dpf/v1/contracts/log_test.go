package contracts_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/contracts"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("logs", func() {
	var (
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 core.Log
		atTime types.Time
		slist  contracts.LogList
	)
	BeforeEach(func() {
		atTime, err = types.ParseTime(time.RFC3339Nano, "2021-06-20T07:55:17.753Z")
		Expect(err).To((Succeed()))
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Log{
			Time:      atTime,
			LogType:   "service",
			Operator:  "user1",
			Operation: "add_cc_primary",
			Target:    "1",
			Status:    core.LogStatusStart,
		}
		s2 = core.Log{
			Time:      atTime,
			LogType:   "common_config",
			Operator:  "user2",
			Operation: "create_tsig",
			Target:    "2",
			Status:    core.LogStatusSuccess,
		}

		slist = contracts.LogList{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Items: []core.Log{s1, s2},
		}
	})
	Describe("LogList", func() {
		var (
			c contracts.LogList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/logs", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "8026628BA5AD4ECA93F8506972DD50A7",
					"results": [
						{
							"time": "2021-06-20T07:55:17.753Z",
							"log_type": "service",
							"operator": "user1",
							"operation": "add_cc_primary",
							"target": "1",
							"status": "start"
							},
						{
							"time": "2021-06-20T07:55:17.753Z",
							"log_type": "common_config",
							"operator": "user2",
							"operation": "create_tsig",
							"target": "2",
							"status": "success"
							}
					]
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/logs/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "8026628BA5AD4ECA93F8506972DD50A7",
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
					c = contracts.LogList{
						AttributeMeta: contracts.AttributeMeta{
							ContractId: "f1",
						},
					}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("8026628BA5AD4ECA93F8506972DD50A7"))
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
			var nilSpec *contracts.LogList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "logs")
			testtool.TestGetPathMethodForCountableList(&slist, "/contracts/f1/logs")
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
			testtool.TestGetMaxLimit(&slist, 100)
			testtool.TestClearItems(&slist)
			testtool.TestAddItem(&slist, s1)
		})
	})
})
