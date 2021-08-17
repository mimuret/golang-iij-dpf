package core_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ = Describe("jobs", func() {
	var (
		c          core.Job
		cl         *testtool.TestClient
		err        error
		reqId      string
		s1, s2, s3 core.Job
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = core.Job{
			RequestId: "9BCFE2E9C10D4D9A8444CB0B48C72830",
			Status:    core.JobStatusRunning,
		}
		s2 = core.Job{
			RequestId:   "B52CBA3FBA8D4A5C951C4EBD9EB48076",
			Status:      core.JobStatusSuccessful,
			ResourceUrl: "http://localhost/contracts/f1",
		}
		s3 = core.Job{
			RequestId:    "9F16F8E85D104D5C9C6BC58676B5D0BD",
			Status:       core.JobStatusFailed,
			ErrorType:    "fail",
			ErrorMessage: "error message",
		}
	})
	Describe("Contract", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/jobs/9BCFE2E9C10D4D9A8444CB0B48C72830", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "9BCFE2E9C10D4D9A8444CB0B48C72830",
					"status": "RUNNING"
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/jobs/B52CBA3FBA8D4A5C951C4EBD9EB48076", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "B52CBA3FBA8D4A5C951C4EBD9EB48076",
					"status": "SUCCESSFUL",
					"resources_url": "http://localhost/contracts/f1"
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/jobs/9F16F8E85D104D5C9C6BC58676B5D0BD", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "9F16F8E85D104D5C9C6BC58676B5D0BD",
					"status": "FAILED",
					"error_type": "fail",
					"error_message": "error message"
				}`)))

			})
			AfterEach(func() {
				httpmock.Reset()
			})
			When("returns Job s1", func() {
				BeforeEach(func() {
					c = core.Job{
						RequestId: "9BCFE2E9C10D4D9A8444CB0B48C72830",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("9BCFE2E9C10D4D9A8444CB0B48C72830"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns Job s2", func() {
				BeforeEach(func() {
					c = core.Job{
						RequestId: "B52CBA3FBA8D4A5C951C4EBD9EB48076",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("B52CBA3FBA8D4A5C951C4EBD9EB48076"))
					Expect(c).To(Equal(s2))
				})
			})
			When("returns Job s3", func() {
				BeforeEach(func() {
					c = core.Job{
						RequestId: "9F16F8E85D104D5C9C6BC58676B5D0BD",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("9F16F8E85D104D5C9C6BC58676B5D0BD"))
					Expect(c).To(Equal(s3))
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
					err = s1.SetPathParams("f10")
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
				It("can set ContractId", func() {
					Expect(s1.RequestId).To(Equal("f10"))
				})
			})
			When("arguments has extra value", func() {
				BeforeEach(func() {
					err = s1.SetPathParams("f10", 2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("arguments type missmatch (RequestId)", func() {
				BeforeEach(func() {
					err = s1.SetPathParams(2)
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *core.Job
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "jobs")
			testtool.TestGetGroup(&s1, "core")
			Context("GetPathMethod", func() {
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/jobs/9BCFE2E9C10D4D9A8444CB0B48C72830")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, "", "")
				})
			})
		})
		Context("GetError", func() {
			When("running", func() {
				It("returns nil", func() {
					Expect(s1.Status).To(Equal(core.JobStatusRunning))
					Expect(s1.GetError()).To(Succeed())
				})
			})
			When("successful", func() {
				It("returns nil", func() {
					Expect(s2.Status).To(Equal(core.JobStatusSuccessful))
					Expect(s2.GetError()).To(Succeed())
				})
			})
			When("failed", func() {
				It("returns err", func() {
					Expect(s3.Status).To(Equal(core.JobStatusFailed))
					Expect(s3.GetError()).To(HaveOccurred())
				})
			})
		})
	})
	Context("JobStatus", func() {
		Context("String", func() {
			When("JobStatusRunning", func() {
				It("returns RUNNING", func() {
					Expect(core.JobStatusRunning.String()).To(Equal("RUNNING"))
				})
			})
			When("JobStatusRunning", func() {
				It("returns SUCCESSFUL", func() {
					Expect(core.JobStatusSuccessful.String()).To(Equal("SUCCESSFUL"))
				})
			})
			When("JobStatusRunning", func() {
				It("returns FAILED", func() {
					Expect(core.JobStatusFailed.String()).To(Equal("FAILED"))
				})
			})
		})
	})
})
