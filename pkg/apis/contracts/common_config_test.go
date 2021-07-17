package contracts_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestCommonConfig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "9901FF93F9BB430F921C4541BBF6882E",
		"result": {
			"id": 1,
			"name": "共通設定1",
			"managed_dns_enabled": 1,
			"default": 1,
			"description": "common config 1"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C8FEC177EF804512A40EF7FB358037C0",
		"result": {
			"id": 2,
			"name": "共通設定2",
			"managed_dns_enabled": 0,
			"default": 0,
			"description": "common config 2"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D1203ABF5C6E4A73AD1F184F5F2C1A6B",
		"results": [
			{
				"id": 1,
				"name": "共通設定1",
				"managed_dns_enabled": 1,
				"default": 1,
				"description": "common config 1"
			},
			{
				"id": 2,
				"name": "共通設定2",
				"managed_dns_enabled": 0,
				"default": 0,
				"description": "common config 2"
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodPost, "http://localhost/contracts/f1/common_configs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D0D5476A5DE2464181719F9D204872A9",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/D0D5476A5DE2464181719F9D204872A9"
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/common_configs/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "7B6D76C947054539911197484EC6C931",
		"result": {
			"count": 2
		}
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "C08F5FB10D744C5597B6001BE39CB1B9",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/C08F5FB10D744C5597B6001BE39CB1B9"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "C08F5FB10D744C5597B6001BE39CB1B9",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/C08F5FB10D744C5597B6001BE39CB1B9"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/common_configs/1", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "A8FDB4B8E3A54065AB0CF43B5B225F45",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/A8FDB4B8E3A54065AB0CF43B5B225F45"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/common_configs/2", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "A8FDB4B8E3A54065AB0CF43B5B225F45",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/A8FDB4B8E3A54065AB0CF43B5B225F45"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      contracts.CommonConfig
	}{
		{
			RequestId: "9901FF93F9BB430F921C4541BBF6882E",
			Spec: contracts.CommonConfig{
				AttributeMeta: contracts.AttributeMeta{
					ContractId: "f1",
				},
				Id:                1,
				Name:              "共通設定1",
				ManagedDNSEnabled: types.Enabled,
				Default:           types.Enabled,
				Description:       "common config 1",
			},
		},
		{
			RequestId: "C8FEC177EF804512A40EF7FB358037C0",
			Spec: contracts.CommonConfig{
				AttributeMeta: contracts.AttributeMeta{
					ContractId: "f1",
				},
				Id:                2,
				Name:              "共通設定2",
				ManagedDNSEnabled: types.Disabled,
				Default:           types.Disabled,
				Description:       "common config 2",
			},
		},
	}
	// GET /contracts/{ContractId}/common_configs
	list := &contracts.CommonConfigList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "D1203ABF5C6E4A73AD1F184F5F2C1A6B") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &contracts.CommonConfigList{}
	assert.NoError(t, list.SetParams("f1"))
	reqId, err = client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "D1203ABF5C6E4A73AD1F184F5F2C1A6B") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}

	// POST /contracts/{ContractId}/common_configs
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "D0D5476A5DE2464181719F9D204872A9")
		}
	}
	// GET /contracts/{ContractId}/common_configs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "7B6D76C947054539911197484EC6C931")
	}

	// GET /contracts/{ContractId}/common_configs/{CommonConfigId}
	for _, tc := range testcase {
		s := contracts.CommonConfig{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = contracts.CommonConfig{}
		assert.NoError(t, s.SetParams("f1", tc.Spec.Id))
		reqId, err = client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
	}
	// Patch /contracts/{ContractId}/common_configs/{CommonConfigId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "C08F5FB10D744C5597B6001BE39CB1B9")
		}
	}
	// Delete /contracts/{ContractId}/common_configs/{CommonConfigId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "A8FDB4B8E3A54065AB0CF43B5B225F45")
		}
	}

	createTestCase := []map[string]interface{}{
		{
			"name":        "共通設定1",
			"description": "common config 1",
		},
		{
			"name":        "共通設定2",
			"description": "common config 2",
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
			"name":        "共通設定1",
			"description": "common config 1",
		},
		{
			"name":        "共通設定2",
			"description": "common config 2",
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

func TestCommonConfigDefault(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/default", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "02BF0C2901E84DA08D4290691C13F06F",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/02BF0C2901E84DA08D4290691C13F06F"
	}`)))
	// PATCH /contracts/{ContractId}/common_configs/default
	tcDefault := &contracts.CommonConfigDefault{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
		CommonConfigId: 1,
	}
	client := api.NewClient("", "http://localhost", nil)
	reqId, err := client.Apply(tcDefault, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "02BF0C2901E84DA08D4290691C13F06F")
	}
	tcDefault = &contracts.CommonConfigDefault{
		CommonConfigId: 1,
	}
	if assert.NoError(t, tcDefault.SetParams("f1")) {
		reqId, err := client.Apply(tcDefault, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "02BF0C2901E84DA08D4290691C13F06F")
		}
	}
	updateTestCase := map[string]interface{}{
		"common_config_id": 1,
	}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tcDefault)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
		}
	}
}

func TestCommonConfigManagedDns(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/common_configs/1/managed_dns", httpmock.NewBytesResponder(202, []byte(`{
		"request_id": "10A13F6BDF6F476CBECDEF9F0B6FF585",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/10A13F6BDF6F476CBECDEF9F0B6FF585"
	}`)))
	// PATCH /contracts/{ContractId}/common_configs/managed_dns
	tcManagedDns := &contracts.CommonConfigManagedDns{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
		Id:                1,
		ManagedDnsEnabled: types.Enabled,
	}
	client := api.NewClient("", "http://localhost", nil)
	reqId, err := client.Apply(tcManagedDns, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "10A13F6BDF6F476CBECDEF9F0B6FF585")
	}
	tcManagedDns = &contracts.CommonConfigManagedDns{
		ManagedDnsEnabled: types.Enabled,
	}
	if assert.NoError(t, tcManagedDns.SetParams("f1", 1)) {
		client := api.NewClient("", "http://localhost", nil)
		reqId, err := client.Apply(tcManagedDns, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "10A13F6BDF6F476CBECDEF9F0B6FF585")
		}
	}
	updateTestCase := map[string]interface{}{
		"managed_dns_enabled": 1,
	}
	bs, err := testtool.MarshalMap(updateTestCase)
	if assert.NoError(t, err) {
		updateBody, err := api.MarshalUpdate(tcManagedDns)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(updateBody), test.UnmarshalToMapString(bs))
		}
	}
}

func TestCommonConfigListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword contracts.CommonConfigListSearchKeywords
		values  url.Values
	}{
		{
			contracts.CommonConfigListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeAND,
					Offset: int32(10),
					Limit:  int32(100),
				},
			},
			url.Values{
				"type":   []string{"AND"},
				"offset": []string{"10"},
				"limit":  []string{"100"},
			},
		},
		{
			contracts.CommonConfigListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:    api.KeywordsString{"hogehoge", "🐰"},
				Name:        api.KeywordsString{"hogehogeあ🍺"},
				Description: api.KeywordsString{"あああ", "🍺"},
			},
			url.Values{
				"type":                    []string{"OR"},
				"offset":                  []string{"10"},
				"limit":                   []string{"100"},
				"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
				"_keywords_name[]":        []string{"hogehogeあ🍺"},
				"_keywords_description[]": []string{"あああ", "🍺"},
			},
		},
	}
	for _, tc := range testcase {
		s, err := query.Values(tc.keyword)
		if assert.NoError(t, err) {
			assert.Equal(t, s, tc.values)
		}
	}
}
