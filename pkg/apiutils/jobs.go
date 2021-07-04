package apiutils

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

func WaitJob(c *api.Client, jobId string) (*core.Job, error) {
	job := &core.Job{
		RequestId: jobId,
	}
	if _, err := c.Read(job); err != nil {
		return nil, err
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for job.Status == core.JobStatusRunning {
		if err := c.WatchRead(ctx, job); err != nil {
			return nil, err
		}
	}
	if job.Status == core.JobStatusFailed {
		return nil, fmt.Errorf("JobId %s job failed: type: %s msg: %s", jobId, job.ErrorType, job.ErrorMessage)
	}
	return job, nil
}
