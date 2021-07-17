package zones_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTTL(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/default_ttl", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "19AB6F91095C4E4E83659AD75ADFE0E3",
		"result": {
			"value": 3600,
			"state": 0,
			"operator": "user1"
		}
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/default_ttl", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "BAE4B17422364C5F84E2020CF2FDC28A",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/BAE4B17422364C5F84E2020CF2FDC28A"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/default_ttl/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "8CD3CF7558864069B32D894FACDB7109",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/8CD3CF7558864069B32D894FACDB7109"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	tc := &zones.DefaultTTL{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Value:    3600,
		State:    zones.DefaultTTLStateApplied,
		Operator: "user1",
	}
	// GET /zones/{ZoneId}/default_ttl
	spec := &zones.DefaultTTL{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.Read(spec)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "19AB6F91095C4E4E83659AD75ADFE0E3")
		assert.Equal(t, spec, tc)
	}
	spec = &zones.DefaultTTL{}
	if assert.NoError(t, spec.SetParams("m1")) {
		reqId, err := client.Read(spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "19AB6F91095C4E4E83659AD75ADFE0E3")
			assert.Equal(t, spec, tc)
		}
	}

	// PATCH /zones{ZoneId}/default_ttl
	reqId, err = client.Update(spec, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "BAE4B17422364C5F84E2020CF2FDC28A")
	}
	// Delete /zones{ZoneId}/default_ttl/changes
	reqId, err = client.Cancel(spec)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "8CD3CF7558864069B32D894FACDB7109")
	}
	updateTestCase := map[string]interface{}{
		"value": 3600,
	}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tc)
		if assert.NoError(t, err) {
			assert.Equal(t, testtool.UnmarshalToMapString(updateBody), testtool.UnmarshalToMapString(bs))
		}
	}

}

func TestDefaultTTLDiffList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/default_ttl/diffs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "6FB71A568C66435BB1B7DCBC6911D881",
		"results": [
			{
				"new": {
					"value": 3600,
					"state": 0,
					"operator": "user1"		
				},
				"old": {
					"value": 300,
					"state": 0,
					"operator": "user2"
				}
		}]
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	tc := &zones.DefaultTTLDiffList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Items: []zones.DefaultTTLDiff{
			{
				New: &zones.DefaultTTL{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Value:    3600,
					State:    zones.DefaultTTLStateApplied,
					Operator: "user1",
				},
				Old: &zones.DefaultTTL{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Value:    300,
					State:    zones.DefaultTTLStateApplied,
					Operator: "user2",
				},
			},
		},
	}
	list := &zones.DefaultTTLDiffList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	// GET /zones/{ZoneId}/default_ttl/diffs
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "6FB71A568C66435BB1B7DCBC6911D881")
		assert.Equal(t, list, tc)
	}
	list = &zones.DefaultTTLDiffList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "6FB71A568C66435BB1B7DCBC6911D881")
			assert.Equal(t, list, tc)
		}
	}
}
