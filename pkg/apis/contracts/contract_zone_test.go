package contracts_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestContractZone(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/zones", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "E8FCA53D138A4AAD8DCED5514A5AF7CB",
		"results": [
			{
				"id": "m1",
				"common_config_id": 1,
				"service_code": "dpm0000001",
				"state": 1,
				"favorite": 1,
				"name": "example.jp.",
				"network": "",
				"description": "zone 1",
				"zone_proxy_enabled": 0
			},
			{
				"id": "m2",
				"common_config_id": 2,
				"service_code": "dpm0000002",
				"state": 2,
				"favorite":2,
				"name": "168.192.in-addr.arpa.",
				"network": "192.168.0.0/16",
				"description": "zone 2",
				"zone_proxy_enabled": 1
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/zones/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "11AE53EAC6D44124AE9116980BC31021",
		"result": {
			"count": 2
		}
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []core.Zone{
		{
			Id:               "m1",
			CommonConfigId:   1,
			ServiceCode:      "dpm0000001",
			State:            types.StateBeforeStart,
			Favorite:         types.FavoriteHighPriority,
			Name:             "example.jp.",
			Network:          "",
			Description:      "zone 1",
			ZoneProxyEnabled: types.Disabled,
		},
		{
			Id:               "m2",
			CommonConfigId:   2,
			ServiceCode:      "dpm0000002",
			State:            types.StateRunning,
			Favorite:         types.FavoriteLowPriority,
			Name:             "168.192.in-addr.arpa.",
			Network:          "192.168.0.0/16",
			Description:      "zone 2",
			ZoneProxyEnabled: types.Enabled,
		},
	}
	list := &contracts.ContractZoneList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	// GET /contracts/{ContractId}/zones
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "E8FCA53D138A4AAD8DCED5514A5AF7CB") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &contracts.ContractZoneList{}
	if assert.NoError(t, list.SetParams("f1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "E8FCA53D138A4AAD8DCED5514A5AF7CB") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}
	// GET /contracts/{ContractId}/zones/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "11AE53EAC6D44124AE9116980BC31021")
	}

}

func TestContractZoneCommonConfig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/zones/common_configs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "E1D3B778EB8440ED953B84241AC40C37",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/E1D3B778EB8440ED953B84241AC40C37"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	// PATCH /contracts/{ContractId}/zones/common_configs
	tc := &contracts.ContractZoneCommonConfig{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
		CommonConfigId: 1,
		ZoneIds: []string{
			"m1",
			"m2",
		},
	}
	reqId, err := client.Apply(tc, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "E1D3B778EB8440ED953B84241AC40C37")
	}
	tc = &contracts.ContractZoneCommonConfig{
		CommonConfigId: 1,
		ZoneIds: []string{
			"m1",
			"m2",
		},
	}
	if assert.NoError(t, tc.SetParams("f1")) {
		reqId, err := client.Apply(tc, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "E1D3B778EB8440ED953B84241AC40C37")
		}
	}
	updateTestCase := map[string]interface{}{
		"common_config_id": 1,
		"zone_ids":         []string{"m1", "m2"},
	}
	bs, err := api.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tc)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
		}
	}

}
