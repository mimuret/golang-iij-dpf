package core_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/stretchr/testify/assert"
)

func TestJob(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
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

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		ReqID string
		Spec  core.Job
	}{
		{
			Spec: core.Job{
				RequestId: "9BCFE2E9C10D4D9A8444CB0B48C72830",
				Status:    core.JobStatusRunning,
			},
		},
		{
			Spec: core.Job{
				RequestId:   "B52CBA3FBA8D4A5C951C4EBD9EB48076",
				Status:      core.JobStatusSuccessful,
				ResourceUrl: "http://localhost/contracts/f1",
			},
		},
		{
			Spec: core.Job{
				RequestId:    "9F16F8E85D104D5C9C6BC58676B5D0BD",
				Status:       core.JobStatusFailed,
				ErrorType:    "fail",
				ErrorMessage: "error message",
			},
		},
	}
	// GET /job/{RequestId}
	for _, tc := range testcase {
		s := core.Job{
			RequestId: tc.Spec.RequestId,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.Spec.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = core.Job{}
		if assert.NoError(t, s.SetParams(tc.Spec.RequestId)) {
			reqId, err := client.Read(&s)
			if assert.NoError(t, err) {
				assert.Equal(t, reqId, tc.Spec.RequestId)
				assert.Equal(t, s, tc.Spec)
			}
		}
	}

}
