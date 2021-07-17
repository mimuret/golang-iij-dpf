package zones_test

import (
	"net"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestZoneProxy(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_proxy", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0607103482984C3B9EC782480423CE63",
		"result": {
			"enabled": 1
		}
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/zone_proxy", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "24CF745B54A24E7BA9D613FBDA8AC849",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	tc := &zones.ZoneProxy{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Enabled: types.Enabled,
	}
	// GET /zones/{ZoneId}/zone_proxy
	spec := &zones.ZoneProxy{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Enabled: types.Enabled,
	}
	reqId, err := client.Read(spec)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "0607103482984C3B9EC782480423CE63")
		assert.Equal(t, spec, tc)
	}
	spec = &zones.ZoneProxy{}
	if assert.NoError(t, spec.SetParams("m1")) {
		reqId, err := client.Read(spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "0607103482984C3B9EC782480423CE63")
			assert.Equal(t, spec, tc)
		}
	}
	// PATCH /zones/{ZoneId}/zone_proxy
	reqId, err = client.Update(spec, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "24CF745B54A24E7BA9D613FBDA8AC849")
	}
	updateTestCase := map[string]interface{}{
		"enabled": 1,
	}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(spec)
		if assert.NoError(t, err) {
			assert.Equal(t, testtool.UnmarshalToMapString(updateBody), testtool.UnmarshalToMapString(bs))
		}
	}
}
func TestZoneProxyHealthCheck(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_proxy/health_check", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "751E0AB9585E48059090FA1B46E5B7A3",
		"results": [
			{
				"address": "192.168.0.1",
				"status": "success",
				"tsig_name": "",
				"enabled": 0
			},
			{
				"address": "2001:db8::1",
				"status": "fail",
				"tsig_name": "tsig2",
				"enabled": 1
			}
		]
	}`)))
	testcase := []zones.ZoneProxyHealthCheck{
		{
			Address:  net.ParseIP("192.168.0.1"),
			Status:   zones.ZoneProxyStatusSuccess,
			TsigName: "",
			Enabled:  types.Disabled,
		},
		{
			Address:  net.ParseIP("2001:db8::1"),
			Status:   zones.ZoneProxyStatusFail,
			TsigName: "tsig2",
			Enabled:  types.Enabled,
		},
	}
	// GET /zones/{ZoneId}/zone_proxy/health_check
	list := &zones.ZoneProxyHealthCheckList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	client := api.NewClient("", "http://localhost", nil)
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "751E0AB9585E48059090FA1B46E5B7A3") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
}
