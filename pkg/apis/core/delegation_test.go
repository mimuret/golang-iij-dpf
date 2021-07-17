package core_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestDelegation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/delegations", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "51367A80A2B447EFA490AF379F12DD48",
		"results": [
			{
				"id": "m1",
				"service_code": "dpm000001",
				"name": "example.jp.",
				"network": "",
				"description": "zone 1",
				"delegation_requested_at": "2021-06-20T13:21:12.881Z"
			},
			{
				"id": "m2",
				"service_code": "dpm000002",
				"name": "168.192.in-addr.arpa.",
				"network": "192.168.0.0/16",
				"description": "zone 2",
				"delegation_requested_at": ""
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/delegations/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D371175B1BB248BE812A845C351FDE82",
		"result": {
			"count": 2
		}
	}`)))

	httpmock.RegisterResponder(http.MethodPost, "http://localhost/delegations", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0342A4B9A2EC4F1A86F8065B016E5BB8",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))
	atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T13:21:12.881Z")
	if err != nil {
		t.Fatal(err)
	}

	client := api.NewClient("", "http://localhost", nil)
	testcase := []core.Delegation{
		{
			Id:                    "m1",
			ServiceCode:           "dpm000001",
			Name:                  "example.jp.",
			Network:               "",
			Description:           "zone 1",
			DelegationRequestedAt: atTime,
		},
		{
			Id:                    "m2",
			ServiceCode:           "dpm000002",
			Name:                  "168.192.in-addr.arpa.",
			Network:               "192.168.0.0/16",
			Description:           "zone 2",
			DelegationRequestedAt: types.Time{},
		},
	}
	// GET /delegations
	list := &core.DelegationList{}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "51367A80A2B447EFA490AF379F12DD48") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	// GET /delegations/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "D371175B1BB248BE812A845C351FDE82")
		assert.Equal(t, list.GetCount(), int32(2))
	}

}

func TestDelegationApply(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/delegations/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D371175B1BB248BE812A845C351FDE82",
		"result": {
			"count": 2
		}
	}`)))

	httpmock.RegisterResponder(http.MethodPost, "http://localhost/delegations", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0342A4B9A2EC4F1A86F8065B016E5BB8",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	// POST /delegations
	apply := &core.DelegationApply{
		ZoneIds: []string{"m1", "m2"},
	}
	reqId, err := client.Apply(apply, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "0342A4B9A2EC4F1A86F8065B016E5BB8")
	}
	applyTestCase := map[string]interface{}{
		"zone_ids": []string{"m1", "m2"},
	}
	bs, err := testtool.MarshalMap(applyTestCase)
	if assert.NoError(t, err) {
		applyBody, err := api.MarshalApply(apply)
		if assert.NoError(t, err) {
			assert.Equal(t, test.UnmarshalToMapString(applyBody), test.UnmarshalToMapString(bs))
		}
	}
}

func TestDelegationListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword core.DelegationListSearchKeywords
		values  url.Values
	}{
		{
			core.DelegationListSearchKeywords{
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
			core.DelegationListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:    api.KeywordsString{"hogehoge", "🐰"},
				ServiceCode: api.KeywordsString{"mxxxxxx", "mxxxxxx1"},
				Name:        api.KeywordsString{"example.jp", "example.net"},
				Network:     api.KeywordsString{"192.168.0.0/24"},
				Favorite:    api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
				Description: api.KeywordsString{"あああ", "🍺"},
				Requested:   api.KeywordsBoolean{types.Disabled, types.Enabled},
			},
			url.Values{
				"type":                     []string{"OR"},
				"offset":                   []string{"10"},
				"limit":                    []string{"100"},
				"_keywords_full_text[]":    []string{"hogehoge", "🐰"},
				"_keywords_service_code[]": []string{"mxxxxxx", "mxxxxxx1"},
				"_keywords_name[]":         []string{"example.jp", "example.net"},
				"_keywords_network[]":      []string{"192.168.0.0/24"},
				"_keywords_favorite[]":     []string{"1", "2"},
				"_keywords_description[]":  []string{"あああ", "🍺"},
				"_keywords_requested[]":    []string{"0", "1"},
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
