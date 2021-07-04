package common_configs_test

import (
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestCcSecNotifiedServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "5D7FECB0A56C49C2B73BD45DD174A575",
		"result": {
			"id": 1,
			"address": "192.168.0.1",
			"tsig_id": 0
		}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "5ABBBEF090DF4F208AD9EDEAC7D702A5",
		"result": {
			"id": 2,
			"address": "2001:db8::1",
			"tsig_id": 2
		}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_notified_servers", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "450B94B486D34D548313707FC7590CE0",
		"results": [
			{
				"id": 1,
				"address": "192.168.0.1",
				"tsig_id": 0
			},
			{
				"id": 2,
				"address": "2001:db8::1",
				"tsig_id": 2
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/1/cc_sec_notified_servers", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "0C03D9AD2BBE488CB7E5C29241C44AC2",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0C03D9AD2BBE488CB7E5C29241C44AC2"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "AD6699A67620481F8ACD93848970FD57",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/AD6699A67620481F8ACD93848970FD57"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "AD6699A67620481F8ACD93848970FD57",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/AD6699A67620481F8ACD93848970FD57"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_notified_servers/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "5605D84FD5D04A7FB42FD299AD1EBF89",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/5605D84FD5D04A7FB42FD299AD1EBF89"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_notified_servers/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "5605D84FD5D04A7FB42FD299AD1EBF89",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/5605D84FD5D04A7FB42FD299AD1EBF89"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      common_configs.CcSecNotifiedServer
	}{
		{
			RequestId: "5D7FECB0A56C49C2B73BD45DD174A575",
			Spec: common_configs.CcSecNotifiedServer{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      1,
				Address: net.ParseIP("192.168.0.1"),
				TsigId:  0,
			},
		},
		{
			RequestId: "5ABBBEF090DF4F208AD9EDEAC7D702A5",
			Spec: common_configs.CcSecNotifiedServer{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      2,
				Address: net.ParseIP("2001:db8::1"),
				TsigId:  2,
			},
		},
	}
	// GET /common_configs/{CommonConfigId}/cc_sec_notified_servers
	list := &common_configs.CcSecNotifiedServerList{
		AttributeMeta: common_configs.AttributeMeta{
			CommonConfigId: 1,
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "450B94B486D34D548313707FC7590CE0") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &common_configs.CcSecNotifiedServerList{}
	list.SetParams(1)
	reqId, err = client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "450B94B486D34D548313707FC7590CE0") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	// POST /common_configs/{CommonConfigId}/cc_sec_notified_servers
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		assert.Equal(t, reqId, "0C03D9AD2BBE488CB7E5C29241C44AC2")
		assert.NoError(t, err)
	}
	// GET /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId}
	for _, tc := range testcase {
		s := common_configs.CcSecNotifiedServer{
			AttributeMeta: common_configs.AttributeMeta{
				CommonConfigId: 1,
			},
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		fmt.Printf("%v, %v", tc, err)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = common_configs.CcSecNotifiedServer{}
		assert.NoError(t, s.SetParams(1, tc.Spec.Id))
		reqId, err = client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
	}
	// Patch /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "AD6699A67620481F8ACD93848970FD57")
		}
	}
	// Delete /common_configs/{CommonConfigId}/cc_sec_notified_servers/{CcSecNotifiedServerId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "5605D84FD5D04A7FB42FD299AD1EBF89")
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
		bs, err := api.MarshalMap(tc)
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
		},
		{
			"address": "2001:db8::1",
			"tsig_id": 2,
		},
	}
	for i, tc := range updateTestCase {
		bs, err := api.MarshalMap(tc)
		if assert.NoError(t, err) {
			updateBody, err := api.MarshalUpdate(testcase[i].Spec)
			if assert.NoError(t, err) {
				assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
			}
		}
	}
}
