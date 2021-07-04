package zones_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestContract(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/contract", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "A4FBAF142456476BBBA919737A4C7663",
		"result": {
			"id": "f1",
			"service_code": "dpf0000001",
			"state": 1,
			"favorite": 1,
			"plan": 1,
			"description": "contract 1"
			}
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := zones.Contract{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Contract: core.Contract{
			Id:          "f1",
			ServiceCode: "dpf0000001",
			State:       types.StateBeforeStart,
			Favorite:    types.FavoriteHighPriority,
			Plan:        core.PlanBasic,
			Description: "contract 1",
		},
	}
	s := zones.Contract{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	// GET /zones/{ZoneId}/contract
	reqId, err := client.Read(&s)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "A4FBAF142456476BBBA919737A4C7663")
		assert.Equal(t, s, testcase)
	}
	s = zones.Contract{}
	if assert.NoError(t, s.SetParams("m1")) {
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "A4FBAF142456476BBBA919737A4C7663")
			assert.Equal(t, s, testcase)
		}
	}
}
