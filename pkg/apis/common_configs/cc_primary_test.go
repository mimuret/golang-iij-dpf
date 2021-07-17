package common_configs_test

import (
	"net"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCcPrimaryG(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "73A87F978AFD42D8B0AC810218AEC7CC",
		"result": {
			"id": 1,
			"address": "192.168.0.1",
			"tsig_id": 0,
			"enabled": 0
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "EBCA51F1CCD8457FB242766B59B5786E",
		"result": {
			"id": 2,
			"address": "2001:db8::1",
			"tsig_id": 2,
			"enabled": 1
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_primaries", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "1194A148B9B34C5F81977EF17A39CD4C",
		"results": [
			{
				"id": 1,
				"address": "192.168.0.1",
				"tsig_id": 0,
				"enabled": 0
			},
			{
				"id": 2,
				"address": "2001:db8::1",
				"tsig_id": 2,
				"enabled": 1
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/1/cc_primaries", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "5FE53B81E09B44B49BAD2A013BF60F0E",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/5FE53B81E09B44B49BAD2A013BF60F0E"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "EFF4DFBA20A44B0E95B8C634B884D12E",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/EFF4DFBA20A44B0E95B8C634B884D12E"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "EFF4DFBA20A44B0E95B8C634B884D12E",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/EFF4DFBA20A44B0E95B8C634B884D12E"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_primaries/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "225BC653AB1C421181EBDBBF51E4ED64",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/225BC653AB1C421181EBDBBF51E4ED64"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_primaries/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "225BC653AB1C421181EBDBBF51E4ED64",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/225BC653AB1C421181EBDBBF51E4ED64"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      common_configs.CcPrimary
	}{
		{
			RequestId: "73A87F978AFD42D8B0AC810218AEC7CC",
			Spec: common_configs.CcPrimary{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      1,
				Address: net.ParseIP("192.168.0.1"),
				TsigId:  0,
				Enabled: types.Disabled,
			},
		},
		{
			RequestId: "EBCA51F1CCD8457FB242766B59B5786E",
			Spec: common_configs.CcPrimary{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      2,
				Address: net.ParseIP("2001:db8::1"),
				TsigId:  2,
				Enabled: types.Enabled,
			},
		},
	}
	// GET /common_configs/{CommonConfigId}/cc_primaries
	list := &common_configs.CcPrimaryList{
		AttributeMeta: common_configs.AttributeMeta{
			CommonConfigId: 1,
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "1194A148B9B34C5F81977EF17A39CD4C") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &common_configs.CcPrimaryList{}
	list.SetParams(1)
	reqId, err = client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "1194A148B9B34C5F81977EF17A39CD4C") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}

	// POST /common_configs/{CommonConfigId}/cc_primaries
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		assert.Equal(t, reqId, "5FE53B81E09B44B49BAD2A013BF60F0E")
		assert.NoError(t, err)
	}
	// GET /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId}
	for _, tc := range testcase {
		s := common_configs.CcPrimary{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = common_configs.CcPrimary{}
		assert.NoError(t, s.SetParams(1, tc.Spec.Id))
		reqId, err = client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
	}
	// Patch /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "EFF4DFBA20A44B0E95B8C634B884D12E")
		}
	}
	// Delete /common_configs/{CommonConfigId}/cc_primaries/{CcPrimaryId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "225BC653AB1C421181EBDBBF51E4ED64")
		}
	}
	createTestCase := []map[string]interface{}{
		{
			"address": "192.168.0.1",
		},
		{
			"address": "2001:db8::1",
			"tsig_id": 2,
		},
	}
	for i, tc := range createTestCase {
		bs, err := testtool.MarshalMap(tc)
		if assert.NoError(t, err) {
			createBody, err := api.MarshalCreate(testcase[i].Spec)
			if assert.NoError(t, err) {
				assert.Equal(t, test.UnmarshalToMapString(createBody), test.UnmarshalToMapString(bs))
			}
		}
	}
	updateTestCase := []map[string]interface{}{
		{
			"address": "192.168.0.1",
			"enabled": 0,
		},
		{
			"address": "2001:db8::1",
			"tsig_id": 2,
			"enabled": 1,
		},
	}
	for i, tc := range updateTestCase {
		bs, err := testtool.MarshalMap(tc)
		if assert.NoError(t, err) {
			updateBody, err := api.MarshalUpdate(testcase[i].Spec)
			if assert.NoError(t, err) {
				assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
			}
		}
	}
}
