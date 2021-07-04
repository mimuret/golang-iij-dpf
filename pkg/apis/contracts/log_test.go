package contracts_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

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
	atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T07:55:17.753Z")
	if err != nil {
		t.Fatal(err)
	}
	client := api.NewClient("", "http://localhost", nil)
	testcase := []core.Log{
		{
			Time:      atTime,
			LogType:   "service",
			Operator:  "user1",
			Operation: "add_cc_primary",
			Target:    "1",
			Status:    core.LogStatusStart,
		},
		{
			Time:      atTime,
			LogType:   "common_config",
			Operator:  "user2",
			Operation: "create_tsig",
			Target:    "2",
			Status:    core.LogStatusSuccess,
		},
	}
	// GET /contracts/{ContractId}/logs
	list := &contracts.LogList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "8026628BA5AD4ECA93F8506972DD50A7") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &contracts.LogList{}
	if assert.NoError(t, list.SetParams("f1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "8026628BA5AD4ECA93F8506972DD50A7") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}

	}
	// GET /contracts/{ContractId}/logs/count
	_, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
	}

}
