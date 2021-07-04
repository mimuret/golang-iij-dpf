package zones_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestDsRecord(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/ds_records", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "CD445FC65BBA4977B78B347A87633947",
		"results": [
			{
				"rrset": "46369 8 2 39F054DCB3EC1E93D8AE6D8F1AAAD91794055EA36895045FAF6F65F0 2FEBC579",
				"transited_at": "2021-06-20T10:23:51.071Z"
			},
			{
				"rrset": "57367 8 2 C069902A17BC8E19F1D5809272219B694C09EC5BCF4AAD4911FC765B 1C233A69",
				"transited_at": "2021-06-20T10:23:51.071Z"
			}
		]
	}`)))
	atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:23:51.071Z")
	if err != nil {
		t.Fatal(err)
	}
	client := api.NewClient("", "http://localhost", nil)
	testcase := []zones.DsRecord{
		{
			RRSet:     "46369 8 2 39F054DCB3EC1E93D8AE6D8F1AAAD91794055EA36895045FAF6F65F0 2FEBC579",
			TransitAt: atTime,
		},
		{
			RRSet:     "57367 8 2 C069902A17BC8E19F1D5809272219B694C09EC5BCF4AAD4911FC765B 1C233A69",
			TransitAt: atTime,
		},
	}
	// GET /zones/{ZoneId}/ds_records
	list := &zones.DsRecordList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "CD445FC65BBA4977B78B347A87633947") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &zones.DsRecordList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err = client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "CD445FC65BBA4977B78B347A87633947") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}
}
