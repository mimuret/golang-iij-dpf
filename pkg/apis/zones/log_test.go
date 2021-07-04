package zones_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/logs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "802D7DDDEB8E43519E8767FDFD416FB5",
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
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/logs/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "F215B828F8B24A66884C6C4DEE6A27DA",
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
	// GET /zones/{ZoneId}/logs
	list := &zones.LogList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "802D7DDDEB8E43519E8767FDFD416FB5") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &zones.LogList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "802D7DDDEB8E43519E8767FDFD416FB5") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}

	// GET /zones/{ZoneId}/logs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "F215B828F8B24A66884C6C4DEE6A27DA")
	}

}
