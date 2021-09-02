package apiutils_test

import (
	"context"
	"fmt"
	"time"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apiutils"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("jobs", func() {
	var (
		c   *testtool.TestClient
		err error
	)
	BeforeEach(func() {
		c = testtool.NewTestClient("token", "http://localhost", nil)
	})
	Context("WaitJob", func() {
		When("failed to read first", func() {
			It("return err", func() {
				_, err := apiutils.WaitJob(context.Background(), c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to read Job"))
			})
		})
		When("failed to watch read", func() {
			BeforeEach(func() {
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestId = "33CF62F20EFC468E84F9779BF6FF1B4D"
					job.Status = core.JobStatusRunning
					return "ok", nil
				}
				c.WatchReadFunc = func(ctx context.Context, interval time.Duration, s api.Spec) error {
					return fmt.Errorf("err")
				}
				_, err = apiutils.WaitJob(context.Background(), c, "33CF62F20EFC468E84F9779BF6FF1B4D", time.Second)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("job successful", func() {
			BeforeEach(func() {
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestId = "33CF62F20EFC468E84F9779BF6FF1B4D"
					job.Status = core.JobStatusRunning
					return "ok", nil
				}
				c.WatchReadFunc = func(ctx context.Context, interval time.Duration, s api.Spec) error {
					job := s.(*core.Job)
					job.RequestId = "9BCFE2E9C10D4D9A8444CB0B48C72830"
					job.Status = core.JobStatusSuccessful
					return nil
				}
			})
			It("returns last job", func() {
				eq := &core.Job{
					RequestId: "9BCFE2E9C10D4D9A8444CB0B48C72830",
					Status:    core.JobStatusSuccessful,
				}
				Eventually(func() (*core.Job, error) {
					return apiutils.WaitJob(context.Background(), c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
				}, time.Second*10).Should(Equal(eq))
			})
		})
		When("job failed", func() {
			BeforeEach(func() {
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestId = "33CF62F20EFC468E84F9779BF6FF1B4D"
					job.Status = core.JobStatusFailed
					return "ok", nil
				}
			})
			It("can add first time", func() {
				var err error
				Eventually(func() error {
					_, err = apiutils.WaitJob(context.Background(), c, "9BCFE2E9C10D4D9A8444CB0B48C72830", time.Second)
					return err
				}, time.Second*10).Should(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("JobId 9BCFE2E9C10D4D9A8444CB0B48C72830 job failed"))
			})
		})
	})
	Context("ParseeResourceSystemId", func() {
		var (
			job *core.Job
			id  string
			err error
		)
		When("normal", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: "https://api.dns-platform.jp/dpf/v1/zones/a"}
				id, err = apiutils.ParseeResourceSystemId(job)
			})
			It("return err", func() {
				Expect(err).To(Succeed())
			})
			It("return id", func() {
				Expect(id).To(Equal("a"))
			})
		})
		When("resourceUrl is invalid url", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: "%1"}
				id, err = apiutils.ParseeResourceSystemId(job)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse resource-url"))
			})
		})
		When("resourceUrl is empty", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: ""}
				id, err = apiutils.ParseeResourceSystemId(job)
			})
			It("return empty", func() {
				Expect(err).To(Succeed())
				Expect(id).To(Equal(""))
			})
		})
	})
	Context("ParseeResourceSystemId", func() {
		var (
			job *core.Job
			id  int64
			err error
		)
		When("normal", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: "https://api.dns-platform.jp/dpf/v1/common_configs/100"}
				id, err = apiutils.ParseeResourceId(job)
			})
			It("return err", func() {
				Expect(err).To(Succeed())
			})
			It("return id", func() {
				Expect(id).To(Equal(int64(100)))
			})
		})
		When("resourceUrl is invalid url", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: "%1"}
				id, err = apiutils.ParseeResourceId(job)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse resource-url"))
			})
		})
		When("resourceUrl is empty", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: ""}
				id, err = apiutils.ParseeResourceId(job)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to convert to int64"))
			})
		})
		When("id is not int64", func() {
			BeforeEach(func() {
				job = &core.Job{ResourceUrl: "https://api.dns-platform.jp/dpf/v1/zones/m1"}
				id, err = apiutils.ParseeResourceId(job)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to convert to int64"))
			})
		})
	})
})
