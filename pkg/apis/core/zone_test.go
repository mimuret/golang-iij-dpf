package core_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestZone(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "76D7462F50E3417C96898A827E0D1EC1",
		"result": {
				"id": "m1",
				"common_config_id": 1,
				"service_code": "dpm0000001",
				"state": 1,
				"favorite": 1,
				"name": "example.jp.",
				"network": "",
				"description": "zone 1",
				"zone_proxy_enabled": 0
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "B10F0042282A46C18BFB440B5B58BF8C",
		"result": {
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
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "2C18922432DC48D485613F5383A7ED8E",
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
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "9C518C729E5541D999389C686FE8987D",
		"result": {
			"count": 2
		}
	}`)))

	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "884BE65D7F5E4D60A0E7F71398552A21",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/884BE65D7F5E4D60A0E7F71398552A21"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "884BE65D7F5E4D60A0E7F71398552A21",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/884BE65D7F5E4D60A0E7F71398552A21"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C6739DFF6CBD4F428AA90BF19F0F2738",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/C6739DFF6CBD4F428AA90BF19F0F2738"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m2/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C6739DFF6CBD4F428AA90BF19F0F2738",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/C6739DFF6CBD4F428AA90BF19F0F2738"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      core.Zone
	}{
		{
			RequestId: "76D7462F50E3417C96898A827E0D1EC1",
			Spec: core.Zone{
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
		},
		{
			RequestId: "B10F0042282A46C18BFB440B5B58BF8C",
			Spec: core.Zone{
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
		},
	}
	// GET /zones
	list := &core.ZoneList{}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "2C18922432DC48D485613F5383A7ED8E") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	// GET /zones/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "9C518C729E5541D999389C686FE8987D")
	}

	// GET /zones/{ContractId}
	for _, tc := range testcase {
		s := core.Zone{
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = core.Zone{}
		if assert.NoError(t, s.SetParams(tc.Spec.Id)) {
			reqId, err := client.Read(&s)
			if assert.NoError(t, err) {
				assert.Equal(t, reqId, tc.RequestId)
				assert.Equal(t, s, tc.Spec)
			}
		}
	}
	// PATCH /zones/{ContractId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "884BE65D7F5E4D60A0E7F71398552A21")
		}
	}

	// DELETE /zones/{ContractId}/changed
	for _, tc := range testcase {
		reqId, err := client.Cancel(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "C6739DFF6CBD4F428AA90BF19F0F2738")
		}
	}
	updateTestCase := []map[string]interface{}{
		{
			"favorite":    1,
			"description": "zone 1",
		},
		{
			"favorite":    2,
			"description": "zone 2",
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

func TestZoneApply(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "F25AEB9139EE487AAF5AF650B2F83D09",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/C6739DFF6CBD4F428AA90BF19F0F2738"
	}`)))
	client := api.NewClient("", "http://localhost", nil)

	tc := &core.ZoneApply{
		Id:          "m1",
		Description: "hogehoge",
	}
	// PATCH /zones/{ZoneId}/changes
	reqId, err := client.Apply(tc, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "F25AEB9139EE487AAF5AF650B2F83D09")
	}
	tc = &core.ZoneApply{
		Description: "hogehoge",
	}
	if assert.NoError(t, tc.SetParams("m1")) {
		reqId, err := client.Apply(tc, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "F25AEB9139EE487AAF5AF650B2F83D09")
		}
	}
	applyTestCase := map[string]interface{}{
		"description": "hogehoge",
	}
	bs, err := api.MarshalMap(applyTestCase)
	if assert.NoError(t, err) {
		applyBody, err := api.MarshalApply(tc)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(applyBody), test.UnmarshalToMapString(bs))
		}
	}

}

func TeetZoneListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword core.ZoneListSearchKeywords
		values  url.Values
	}{
		{
			core.ZoneListSearchKeywords{
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
			core.ZoneListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:         api.KeywordsString{"hogehoge", "🐰"},
				ServiceCode:      api.KeywordsString{"mxxxxxx", "mxxxxxx1"},
				Name:             api.KeywordsString{"example.jp", "example.net"},
				Network:          api.KeywordsString{"192.168.0.0/24"},
				State:            api.KeywordsState{types.FavoriteHighPriority, types.FavoriteHighPriority},
				Favorite:         api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
				Description:      api.KeywordsString{"あああ", "🍺"},
				CommonConfigId:   api.KeywordsId{100, 200},
				ZoneProxyEnabled: api.KeywordsBoolean{types.FavoriteHighPriority, types.FavoriteLowPriority},
			},
			/*
				_keywords_full_text[]
				_keywords_service_code[]
				_keywords_name[]
				_keywords_network[]
				_keywords_state[]
				_keywords_favorite[]
				_keywords_description[]
				_keywords_common_config_id[]
				_keywords_zone_proxy_enabled[]
			*/
			url.Values{
				"type":                           []string{"OR"},
				"offset":                         []string{"10"},
				"limit":                          []string{"100"},
				"_keywords_full_text[]":          []string{"hogehoge", "🐰"},
				"_keywords_service_code[]":       []string{"mxxxxxx", "mxxxxxx1"},
				"_keywords_name[]":               []string{"example.jp", "example.net"},
				"_keywords_network[]":            []string{"192.168.0.0/24"},
				"_keywords_state[]":              []string{"1", "2"},
				"_keywords_favorite[]":           []string{"1", "2"},
				"_keywords_description[]":        []string{"あああ", "🍺"},
				"_keywords_common_config_id[]":   []string{"100", "200"},
				"_keywords_zone_proxy_enabled[]": []string{"0", "1"},
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
