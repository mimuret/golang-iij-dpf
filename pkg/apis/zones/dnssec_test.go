package zones_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestDNSSEC(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/dnssec", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "533F010C0DEC4B71891D1491C5308A10",
		"result": {
			"enabled": 1,
			"state": 2,
			"ds_state": 3
		}
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/dnssec", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "DDEACA36C111402C984ECE43A8E5C55C",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/DDEACA36C111402C984ECE43A8E5C55C"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	tc := &zones.Dnssec{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Enabled: types.Enabled,
		State:   zones.DnssecStateEnable,
		DsState: zones.DSStateDisclose,
	}
	// GET /zones/{ZoneId}/dnssec
	spec := &zones.Dnssec{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.Read(spec)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "533F010C0DEC4B71891D1491C5308A10")
		assert.Equal(t, spec, tc)
	}
	spec = &zones.Dnssec{}
	if assert.NoError(t, spec.SetParams("m1")) {
		reqId, err := client.Read(spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "533F010C0DEC4B71891D1491C5308A10")
			assert.Equal(t, spec, tc)
		}
	}

	// PATCH /zones/{ZoneId}/dnssec
	reqId, err = client.Update(spec, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "DDEACA36C111402C984ECE43A8E5C55C")
	}
	updateTestCase := map[string]interface{}{
		"enabled": 1,
	}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tc)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
		}
	}

}

func TestDnssecKskRollover(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/dnssec/ksk_rollover", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "81E9C891FD3440DBAB0E038798140A49",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/81E9C891FD3440DBAB0E038798140A49"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	tc := &zones.DnssecKskRollover{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	// PATCH /zones/{ZoneId}/dnssec/ksk_rollover
	reqId, err := client.Apply(tc, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "81E9C891FD3440DBAB0E038798140A49")
	}
	tc = &zones.DnssecKskRollover{}
	if assert.NoError(t, tc.SetParams("m1")) {
		reqId, err := client.Apply(tc, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "81E9C891FD3440DBAB0E038798140A49")
		}
	}

	updateTestCase := map[string]interface{}{}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tc)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
		}
	}

}
