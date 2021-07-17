package apiutils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

func WaitJob(c api.ClientInterface, jobId string, interval time.Duration) (*core.Job, error) {
	job := &core.Job{
		RequestId: jobId,
	}
	if _, err := c.Read(job); err != nil {
		return nil, err
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for job.Status == core.JobStatusRunning {
		job.RequestId = jobId
		if err := c.WatchRead(ctx, interval, job); err != nil {
			return nil, err
		}
	}
	if job.Status == core.JobStatusFailed {
		return nil, fmt.Errorf("JobId %s job failed: type: %s msg: %s", jobId, job.ErrorType, job.ErrorMessage)
	}
	return job, nil
}
