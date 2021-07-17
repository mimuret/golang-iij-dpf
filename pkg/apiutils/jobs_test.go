package apiutils_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apiutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "apiutils Suite")
}

var _ = Describe("jobs", func() {
	var (
		c *api.Client
	)
	BeforeSuite(func() {
		httpmock.Activate()
	})
	AfterSuite(func() {
		httpmock.DeactivateAndReset()
	})
	BeforeEach(func() {
		httpmock.Reset()
		c = api.NewClient("token", "http://localhost", nil)
	})
	Context("WaitJob", func() {
		When("failed to read", func() {
			It("return err", func() {
				_, err := apiutils.WaitJob(c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to read Job"))
			})
		})
		When("job successful", func() {
			BeforeEach(func() {
				responses := [][]byte{
					[]byte(`{
					"request_id": "9BCFE2E9C10D4D9A8444CB0B48C72830",
					"status": "RUNNING"
				}`),
					[]byte(`{
					"request_id": "9BCFE2E9C10D4D9A8444CB0B48C72830",
					"status": "SUCCESSFUL"
				}`),
				}
				i := -1
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/jobs/9BCFE2E9C10D4D9A8444CB0B48C72830", func(r *http.Request) (*http.Response, error) {
					if i < 1 {
						i++
					}
					return httpmock.NewBytesResponse(200, responses[i]), nil
				})
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("can add first time", func() {
				eq := &core.Job{
					RequestId: "9BCFE2E9C10D4D9A8444CB0B48C72830",
					Status:    core.JobStatusSuccessful,
				}
				Eventually(func() (*core.Job, error) {
					return apiutils.WaitJob(c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
				}, time.Second*10).Should(Equal(eq))
			})
		})
		When("job failed", func() {
			BeforeEach(func() {
				responses := [][]byte{
					[]byte(`{
					"request_id": "9BCFE2E9C10D4D9A8444CB0B48C72830",
					"status": "RUNNING"
				}`),
					[]byte(`{
					"request_id": "9BCFE2E9C10D4D9A8444CB0B48C72830",
					"status": "FAILED"
				}`),
				}
				i := -1
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/jobs/9BCFE2E9C10D4D9A8444CB0B48C72830", func(r *http.Request) (*http.Response, error) {
					if i < 1 {
						i++
					}
					return httpmock.NewBytesResponse(200, responses[i]), nil
				})
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("can add first time", func() {
				var err error
				Eventually(func() error {
					_, err = apiutils.WaitJob(c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
					return err
				}, time.Second*10).Should(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("JobId 9BCFE2E9C10D4D9A8444CB0B48C72830 job failed"))
			})
		})
	})
})
