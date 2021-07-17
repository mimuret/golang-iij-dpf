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

func TestTsig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "96043C4DF2EE4EEF9094FF566F87505A",
		"result": {
      "id": 1,
      "name": "hogehoge1",
      "algorithm": 0,
      "secret": "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
      "description": "tsig 1"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "E16A52E5C5DD4EB7B4A60F99E8FACE95",
		"result": {
			"id": 2,
      "name": "hogehoge2",
      "algorithm": 0,
      "secret": "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
      "description": "tsig 2"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "062ED092113E4509BE4E2E4EF6C6EB2B",
		"results": [
			{
				"id": 1,
				"name": "hogehoge1",
				"algorithm": 0,
				"secret": "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
				"description": "tsig 1"
			},
			{
				"id": 2,
				"name": "hogehoge2",
				"algorithm": 0,
				"secret": "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
				"description": "tsig 2"
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "F60D8455D0C7456286565E4D28A35951",
		"result": {
			"count": 2
		}
	}`)))
	httpmock.RegisterResponder(http.MethodPost, "http://localhost/contracts/f1/tsigs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "7A8E2A3D66534171A35CD5BBE607C871",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/7A8E2A3D66534171A35CD5BBE607C871"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "B6B2A13EED5A4F4CBD685C6A4A614F0D",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/B6B2A13EED5A4F4CBD685C6A4A614F0D"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1/tsigs/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "B6B2A13EED5A4F4CBD685C6A4A614F0D",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/B6B2A13EED5A4F4CBD685C6A4A614F0D"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/tsigs/1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D7812109C489437E90218F114AA8C3B7",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/D7812109C489437E90218F114AA8C3B7"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/contracts/f1/tsigs/2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D7812109C489437E90218F114AA8C3B7",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/D7812109C489437E90218F114AA8C3B7"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      contracts.Tsig
	}{
		{
			RequestId: "96043C4DF2EE4EEF9094FF566F87505A",
			Spec: contracts.Tsig{
				AttributeMeta: contracts.AttributeMeta{
					ContractId: "f1",
				},
				Id:          1,
				Name:        "hogehoge1",
				Algorithm:   contracts.TsigAlgorithmHMACSHA256,
				Secret:      "PDeNDXwzjh0OKbQLaVIZIQe6QR8La1uvfPiCCm3HL8Y=",
				Description: "tsig 1",
			},
		},
		{
			RequestId: "E16A52E5C5DD4EB7B4A60F99E8FACE95",
			Spec: contracts.Tsig{
				AttributeMeta: contracts.AttributeMeta{
					ContractId: "f1",
				},
				Id:          2,
				Name:        "hogehoge2",
				Algorithm:   contracts.TsigAlgorithmHMACSHA256,
				Secret:      "dplF06cJHAiSdZoe0IldzfF5FYq7U5+INTmb/esvXSo=",
				Description: "tsig 2",
			},
		},
	}
	// GET /contracts/{ContractId}/tsigs
	list := &contracts.TsigList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "062ED092113E4509BE4E2E4EF6C6EB2B") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &contracts.TsigList{}
	if assert.NoError(t, list.SetParams("f1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "062ED092113E4509BE4E2E4EF6C6EB2B") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc.Spec)
					}
				}
			}
		}
	}

	// POST /contracts/{ContractId}/tsigs
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "7A8E2A3D66534171A35CD5BBE607C871")
		}
	}
	// GET /contracts/{ContractId}/tsigs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "F60D8455D0C7456286565E4D28A35951")
	}

	// GET 	/contracts/{ContractId}/tsigs/{TsigId}
	for _, tc := range testcase {
		s := contracts.Tsig{
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
		s = contracts.Tsig{}
		if assert.NoError(t, s.SetParams("f1", tc.Spec.Id)) {
			reqId, err := client.Read(&s)
			if assert.NoError(t, err) {
				assert.Equal(t, reqId, tc.RequestId)
				assert.Equal(t, s, tc.Spec)
			}
		}
	}
	// Patch 	/contracts/{ContractId}/tsigs/{TsigId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "B6B2A13EED5A4F4CBD685C6A4A614F0D")
		}
	}
	// Delete 	/contracts/{ContractId}/tsigs/{TsigId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "D7812109C489437E90218F114AA8C3B7")
		}
	}
	createTestCase := []map[string]interface{}{
		{
			"name":        "hogehoge1",
			"description": "tsig 1",
		},
		{
			"name":        "hogehoge2",
			"description": "tsig 2",
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
			"description": "tsig 1",
		},
		{
			"description": "tsig 2",
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

func TestTsigCommonConfig(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1/common_configs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C38FBA61C657432695BC798DD96B0D9A",
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
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/tsigs/1/common_configs/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "1073C4F38C0A4027B4776F0F3E8197BB",
		"result": {
			"count": 2
		}
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []contracts.CommonConfig{
		{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                1,
			Name:              "共通設定1",
			ManagedDNSEnabled: types.Enabled,
			Default:           types.Enabled,
			Description:       "common config 1",
		},
		{
			AttributeMeta: contracts.AttributeMeta{
				ContractId: "f1",
			},
			Id:                2,
			Name:              "共通設定2",
			ManagedDNSEnabled: types.Disabled,
			Default:           types.Disabled,
			Description:       "common config 2",
		},
	}
	// GET /contracts/{ContractId}/tsigs/{TsigId}/common_configs
	list := &contracts.TsigCommonConfigList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
		Id: 1,
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "C38FBA61C657432695BC798DD96B0D9A") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &contracts.TsigCommonConfigList{}
	if assert.NoError(t, list.SetParams("f1", 1)) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "C38FBA61C657432695BC798DD96B0D9A") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}

	// GET /contracts/{ContractId}/tsigs/{TsigId}/common_configs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "1073C4F38C0A4027B4776F0F3E8197BB")
	}

}

func TestTsigListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword contracts.TsigListSearchKeywords
		values  url.Values
	}{
		{
			contracts.TsigListSearchKeywords{
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
			contracts.TsigListSearchKeywords{
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
