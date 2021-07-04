package common_configs_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/common_configs"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCcSecTrabsferAcl(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C5E6406DCBFC4931A0EE480ABA9FA505",
		"result": {
			"id": 1,
			"network": "192.168.1.0/24",
			"tsig_id": 0
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "3D82ADCA5479408094DEF876A679BFD3",
		"result": {
			"id": 2,
			"network": "2001:db8::/64",
			"tsig_id": 2
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/common_configs/1/cc_sec_transfer_acls", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "13032F0785DA47768B3E36A1A854A808",
		"results": [
			{
				"id": 1,
				"network": "192.168.1.0/24",
				"tsig_id": 0
				},
			{
				"id": 2,
				"network": "2001:db8::/64",
				"tsig_id": 2
				}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodPost, "http://localhost/common_configs/1/cc_sec_transfer_acls", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "4DF3754227CD48B68FFEDEEF0FD5FF1A",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/4DF3754227CD48B68FFEDEEF0FD5FF1A"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "694969D627E7409DA5680D05829E1BA7",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/694969D627E7409DA5680D05829E1BA7"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "694969D627E7409DA5680D05829E1BA7",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/694969D627E7409DA5680D05829E1BA7"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_transfer_acls/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "1BEC54E5B26A4ADD9748AC8265389840",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/1BEC54E5B26A4ADD9748AC8265389840"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/common_configs/1/cc_sec_transfer_acls/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "1BEC54E5B26A4ADD9748AC8265389840",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/1BEC54E5B26A4ADD9748AC8265389840"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      common_configs.CcSecTransferAcl
	}{
		{
			RequestId: "C5E6406DCBFC4931A0EE480ABA9FA505",
			Spec: common_configs.CcSecTransferAcl{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      1,
				Network: types.MustParseIPNet("192.168.1.0/24"),
				TsigId:  0,
			},
		},
		{
			RequestId: "3D82ADCA5479408094DEF876A679BFD3",
			Spec: common_configs.CcSecTransferAcl{
				AttributeMeta: common_configs.AttributeMeta{
					CommonConfigId: 1,
				},
				Id:      2,
				Network: types.MustParseIPNet("2001:db8::/64"),
				TsigId:  2,
			},
		},
	}
	// GET /common_configs/{CommonConfigId}/cc_sec_transfer_acls
	list := &common_configs.CcSecTransferAclList{
		AttributeMeta: common_configs.AttributeMeta{
			CommonConfigId: 1,
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "13032F0785DA47768B3E36A1A854A808") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &common_configs.CcSecTransferAclList{}
	list.SetParams(1)
	reqId, err = client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "13032F0785DA47768B3E36A1A854A808") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}

	// POST /common_configs/{CommonConfigId}/cc_sec_transfer_acls
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		assert.Equal(t, reqId, "4DF3754227CD48B68FFEDEEF0FD5FF1A")
		assert.NoError(t, err)
	}
	// GET /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId}
	for _, tc := range testcase {
		s := common_configs.CcSecTransferAcl{
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
		s = common_configs.CcSecTransferAcl{}
		assert.NoError(t, s.SetParams(1, tc.Spec.Id))
		reqId, err = client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
	}
	// Patch /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "694969D627E7409DA5680D05829E1BA7")
		}
	}
	// Delete /common_configs/{CommonConfigId}/cc_sec_transfer_acls/{CcSecTransferAclId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "1BEC54E5B26A4ADD9748AC8265389840")
		}
	}
	createTestCase := []map[string]interface{}{
		{
			"network": "192.168.1.0/24",
		},
		{
			"network": "2001:db8::/64",
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
			"network": "192.168.1.0/24",
		},
		{
			"network": "2001:db8::/64",
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
