package core

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type JobStatus string

var _ apis.Spec = &Job{}

const (
	JobStatusRunning    = "RUNNING"
	JobStatusSuccessful = "SUCCESSFUL"
	JobStatusFailed     = "FAILED"
)

func (c JobStatus) String() string { return string(c) }

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Job struct {
	AttributeMeta
	RequestId    string    `read:"request_id"`
	Status       JobStatus `read:"status"`
	ResourceUrl  string    `read:"resources_url"`
	ErrorType    string    `read:"error_type"`
	ErrorMessage string    `read:"error_message"`
}

func (c *Job) Error() string {
	return "ErrorType: " + c.ErrorType + " Message: " + c.ErrorMessage
}

func (c *Job) GetError() error {
	if c.Status == JobStatusFailed {
		return c
	}
	return nil
}

func (c *Job) GetName() string { return "jobs" }
func (c *Job) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead:
		return action.ToMethod(), fmt.Sprintf("/jobs/%s", c.RequestId)
	}
	return "", ""
}
func (c *Job) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.RequestId)
}

func init() {
	Register.Add(&Job{})
}
