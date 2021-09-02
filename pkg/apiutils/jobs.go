package apiutils

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"path"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

func WaitJob(ctx context.Context, c api.ClientInterface, jobId string, interval time.Duration) (*core.Job, error) {
	job := &core.Job{
		RequestId: jobId,
	}
	if _, err := c.Read(job); err != nil {
		return nil, fmt.Errorf("failed to read Job: %w", err)
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	for job.Status == core.JobStatusRunning {
		job.RequestId = jobId
		if err := c.WatchRead(ctx, interval, job); err != nil {
			return nil, err
		}
	}
	if job.Status == core.JobStatusFailed {
		return job, fmt.Errorf("JobId %s job failed: type: %s msg: %s", jobId, job.ErrorType, job.ErrorMessage)
	}
	return job, nil
}

func ParseeResourceSystemId(job *core.Job) (string, error) {
	u, err := url.Parse(job.ResourceUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse resource-url: %s , %w", job.ResourceUrl, err)
	}
	_, id := path.Split(u.Path)
	return id, nil
}

func ParseeResourceId(job *core.Job) (int64, error) {
	idStr, err := ParseeResourceSystemId(job)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert to int64 resource-url: %s id: %s, %w", job.ResourceUrl, idStr, err)
	}
	return id, nil
}
