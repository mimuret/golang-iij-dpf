package apiutils_test

import (
	"context"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apiutils"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("client.go", func() {
	var (
		c     *testtool.TestClient
		s     *testtool.TestSpec
		reqId string
		job   *core.Job
		err   error
		id    string
	)
	BeforeEach(func() {
		c = testtool.NewTestClient("token", "http://localhost", nil)
		s = &testtool.TestSpec{
			ID:     "id1",
			Name:   "hoge",
			Number: 100,
		}
	})
	Context("SyncCreate", func() {
		When("successful creating", func() {
			BeforeEach(func() {
				id = testtool.GenReqID()
				c.CreateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
					return id, nil
				}
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestID = id
					job.Status = core.JobStatusSuccessful
					return id, nil
				}
				reqId, job, err = apiutils.SyncCreate(context.Background(), c, s, nil)
			})
			It("not returns err", func() {
				Expect(err).To(Succeed())
			})
			It("return job", func() {
				Expect(job).NotTo(BeNil())
				Expect(job.Status).To(Equal(core.JobStatusSuccessful))
			})
			It("returns last request id", func() {
				Expect(reqId).To(Equal(id))
			})
		})
	})
	When("successful registration of creating, but create failed", func() {
		BeforeEach(func() {
			id = testtool.GenReqID()
			c.CreateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
				return id, nil
			}
			c.ReadFunc = func(s api.Spec) (requestId string, err error) {
				job := s.(*core.Job)
				job.RequestID = id
				job.Status = core.JobStatusFailed
				return id, nil
			}
			reqId, job, err = apiutils.SyncCreate(context.Background(), c, s, nil)
		})
		It("not returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("return job", func() {
			Expect(job).NotTo(BeNil())
			Expect(job.Status).To(Equal(core.JobStatusFailed))
		})
		It("returns last request id", func() {
			Expect(reqId).To(Equal(id))
		})
	})
	When("create failed", func() {
		BeforeEach(func() {
			reqId, job, err = apiutils.SyncCreate(context.Background(), c, s, nil)
		})
		It("returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("not return job", func() {
			Expect(job).To(BeNil())
		})
	})
	Context("SyncUpdate", func() {
		When("successful updating", func() {
			BeforeEach(func() {
				id = testtool.GenReqID()
				c.UpdateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
					return id, nil
				}
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestID = id
					job.Status = core.JobStatusSuccessful
					return id, nil
				}
				reqId, job, err = apiutils.SyncUpdate(context.Background(), c, s, nil)
			})
			It("not returns err", func() {
				Expect(err).To(Succeed())
			})
			It("return job", func() {
				Expect(job).NotTo(BeNil())
				Expect(job.Status).To(Equal(core.JobStatusSuccessful))
			})
			It("returns last request id", func() {
				Expect(reqId).To(Equal(id))
			})
		})
	})
	When("successful registration of updating, but update failed", func() {
		BeforeEach(func() {
			id = testtool.GenReqID()
			c.UpdateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
				return id, nil
			}
			c.ReadFunc = func(s api.Spec) (requestId string, err error) {
				job := s.(*core.Job)
				job.RequestID = id
				job.Status = core.JobStatusFailed
				return id, nil
			}
			reqId, job, err = apiutils.SyncUpdate(context.Background(), c, s, nil)
		})
		It("not returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("return job", func() {
			Expect(job).NotTo(BeNil())
			Expect(job.Status).To(Equal(core.JobStatusFailed))
		})
		It("returns last request id", func() {
			Expect(reqId).To(Equal(id))
		})
	})
	When("update failed", func() {
		BeforeEach(func() {
			reqId, job, err = apiutils.SyncUpdate(context.Background(), c, s, nil)
		})
		It("returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("not return job", func() {
			Expect(job).To(BeNil())
		})
	})
	Context("SyncApply", func() {
		When("successful appling", func() {
			BeforeEach(func() {
				id = testtool.GenReqID()
				c.ApplyFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
					return id, nil
				}
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestID = id
					job.Status = core.JobStatusSuccessful
					return id, nil
				}
				reqId, job, err = apiutils.SyncApply(context.Background(), c, s, nil)
			})
			It("not returns err", func() {
				Expect(err).To(Succeed())
			})
			It("return job", func() {
				Expect(job).NotTo(BeNil())
				Expect(job.Status).To(Equal(core.JobStatusSuccessful))
			})
			It("returns last request id", func() {
				Expect(reqId).To(Equal(id))
			})
		})
	})
	When("successful registration of appling, but apply failed", func() {
		BeforeEach(func() {
			id = testtool.GenReqID()
			c.ApplyFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
				return id, nil
			}
			c.ReadFunc = func(s api.Spec) (requestId string, err error) {
				job := s.(*core.Job)
				job.RequestID = id
				job.Status = core.JobStatusFailed
				return id, nil
			}
			reqId, job, err = apiutils.SyncApply(context.Background(), c, s, nil)
		})
		It("not returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("return job", func() {
			Expect(job).NotTo(BeNil())
			Expect(job.Status).To(Equal(core.JobStatusFailed))
		})
		It("returns last request id", func() {
			Expect(reqId).To(Equal(id))
		})
	})
	When("apply failed", func() {
		BeforeEach(func() {
			reqId, job, err = apiutils.SyncApply(context.Background(), c, s, nil)
		})
		It("returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("not return job", func() {
			Expect(job).To(BeNil())
		})
	})
	Context("SyncDelete", func() {
		When("successful deleting", func() {
			BeforeEach(func() {
				id = testtool.GenReqID()
				c.DeleteFunc = func(s api.Spec) (requestId string, err error) {
					return id, nil
				}
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestID = id
					job.Status = core.JobStatusSuccessful
					return id, nil
				}
				reqId, job, err = apiutils.SyncDelete(context.Background(), c, s)
			})
			It("not returns err", func() {
				Expect(err).To(Succeed())
			})
			It("return job", func() {
				Expect(job).NotTo(BeNil())
				Expect(job.Status).To(Equal(core.JobStatusSuccessful))
			})
			It("returns last request id", func() {
				Expect(reqId).To(Equal(id))
			})
		})
	})
	When("successful registration of deleteting, but delete failed", func() {
		BeforeEach(func() {
			id = testtool.GenReqID()
			c.DeleteFunc = func(s api.Spec) (requestId string, err error) {
				return id, nil
			}
			c.ReadFunc = func(s api.Spec) (requestId string, err error) {
				job := s.(*core.Job)
				job.RequestID = id
				job.Status = core.JobStatusFailed
				return id, nil
			}
			reqId, job, err = apiutils.SyncDelete(context.Background(), c, s)
		})
		It("not returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("return job", func() {
			Expect(job).NotTo(BeNil())
			Expect(job.Status).To(Equal(core.JobStatusFailed))
		})
		It("returns last request id", func() {
			Expect(reqId).To(Equal(id))
		})
	})
	When("delete failed", func() {
		BeforeEach(func() {
			reqId, job, err = apiutils.SyncDelete(context.Background(), c, s)
		})
		It("returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("not return job", func() {
			Expect(job).To(BeNil())
		})
	})
	Context("SyncCancel", func() {
		When("successful canceling", func() {
			BeforeEach(func() {
				id = testtool.GenReqID()
				c.CancelFunc = func(s api.Spec) (requestId string, err error) {
					return id, nil
				}
				c.ReadFunc = func(s api.Spec) (requestId string, err error) {
					job := s.(*core.Job)
					job.RequestID = id
					job.Status = core.JobStatusSuccessful
					return id, nil
				}
				reqId, job, err = apiutils.SyncCancel(context.Background(), c, s)
			})
			It("not returns err", func() {
				Expect(err).To(Succeed())
			})
			It("return job", func() {
				Expect(job).NotTo(BeNil())
				Expect(job.Status).To(Equal(core.JobStatusSuccessful))
			})
			It("returns last request id", func() {
				Expect(reqId).To(Equal(id))
			})
		})
	})
	When("successful registration of canceling, but cancel failed", func() {
		BeforeEach(func() {
			id = testtool.GenReqID()
			c.CancelFunc = func(s api.Spec) (requestId string, err error) {
				return id, nil
			}
			c.ReadFunc = func(s api.Spec) (requestId string, err error) {
				job := s.(*core.Job)
				job.RequestID = id
				job.Status = core.JobStatusFailed
				return id, nil
			}
			reqId, job, err = apiutils.SyncCancel(context.Background(), c, s)
		})
		It("not returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("return job", func() {
			Expect(job).NotTo(BeNil())
			Expect(job.Status).To(Equal(core.JobStatusFailed))
		})
		It("returns last request id", func() {
			Expect(reqId).To(Equal(id))
		})
	})
	When("cancel failed", func() {
		BeforeEach(func() {
			reqId, job, err = apiutils.SyncCancel(context.Background(), c, s)
		})
		It("returns err", func() {
			Expect(err).To(HaveOccurred())
		})
		It("not return job", func() {
			Expect(job).To(BeNil())
		})
	})
})
