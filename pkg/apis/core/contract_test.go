package core_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestContract(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "D7155DA183ED47C48437183873DAC765",
		"result": {
			"id": "f1",
			"service_code": "dpf0000001",
			"state": 1,
			"favorite": 1,
			"plan": 1,
			"description": "contract 1"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "EECF5BC0FF834BFFA4F0253224667D10",
		"result": {
			"id": "f2",
			"service_code": "dpf0000002",
			"state": 2,
			"favorite": 2,
			"plan": 2,
			"description": "contract 2"
			}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "1773D7B5F415464C944837053CE8BBF2",
		"results": [
			{
				"id": "f1",
				"service_code": "dpf0000001",
				"state": 1,
				"favorite": 1,
				"plan": 1,
				"description": "contract 1"
				},
			{
				"id": "f2",
				"service_code": "dpf0000002",
				"state": 2,
				"favorite": 2,
				"plan": 2,
				"description": "contract 2"
				}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "BD88B5E38397477989002302583C7FCC",
		"result": {
			"count": 2
		}
	}`)))

	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "A22AC3E5BBAE4F0395828D0D1E50B954",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/contracts/f2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "A22AC3E5BBAE4F0395828D0D1E50B954",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/CC84F36FC71749E4A50BD208DD20E535"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      core.Contract
	}{
		{
			RequestId: "D7155DA183ED47C48437183873DAC765",
			Spec: core.Contract{
				Id:          "f1",
				ServiceCode: "dpf0000001",
				State:       types.StateBeforeStart,
				Favorite:    types.FavoriteHighPriority,
				Plan:        core.PlanBasic,
				Description: "contract 1",
			},
		},
		{
			RequestId: "EECF5BC0FF834BFFA4F0253224667D10",
			Spec: core.Contract{
				Id:          "f2",
				ServiceCode: "dpf0000002",
				State:       types.StateRunning,
				Favorite:    types.FavoriteLowPriority,
				Plan:        core.PlanPremium,
				Description: "contract 2",
			},
		},
	}
	// GET /contracts
	list := &core.ContractList{}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "1773D7B5F415464C944837053CE8BBF2") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	// GET /contracts/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "BD88B5E38397477989002302583C7FCC")
	}

	// GET /contracts/{ContractId}
	for _, tc := range testcase {
		s := core.Contract{
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
		s = core.Contract{}
		if assert.NoError(t, s.SetParams(tc.Spec.Id)) {
			reqId, err := client.Read(&s)
			if assert.NoError(t, err) {
				assert.Equal(t, reqId, tc.RequestId)
				assert.Equal(t, s, tc.Spec)
			}
		}
	}
	// Patch /contracts/{ContractId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "A22AC3E5BBAE4F0395828D0D1E50B954")
		}
	}
	updateTestCase := []map[string]interface{}{
		{
			"favorite":    1,
			"description": "contract 1",
		},
		{
			"favorite":    2,
			"description": "contract 2",
		},
	}
	for i, tc := range updateTestCase {
		bs, err := testtool.MarshalMap(tc)
		if assert.NoError(t, err) {
			updateBody, err := api.MarshalUpdate(testcase[i].Spec)
			if assert.NoError(t, err) {
				assert.Equal(t, testtool.UnmarshalToMapString(updateBody), testtool.UnmarshalToMapString(bs))
			}
		}
	}
}

func TestTsigListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword core.ContractListSearchKeywords
		values  url.Values
	}{
		{
			core.ContractListSearchKeywords{
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
			core.ContractListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:    api.KeywordsString{"hogehoge", "🐰"},
				ServiceCode: api.KeywordsString{"fxxxxxx", "fxxxxxx1"},
				Plan:        core.KeywordsPlan{core.PlanBasic, core.PlanPremium},
				State:       api.KeywordsState{types.StateBeforeStart, types.StateRunning},
				Favorite:    api.KeywordsFavorite{types.FavoriteHighPriority, types.FavoriteLowPriority},
				Description: api.KeywordsString{"あああ", "🍺"},
			},
			url.Values{
				"type":                     []string{"OR"},
				"offset":                   []string{"10"},
				"limit":                    []string{"100"},
				"_keywords_full_text[]":    []string{"hogehoge", "🐰"},
				"_keywords_service_code[]": []string{"fxxxxxx", "fxxxxxx1"},
				"_keywords_plan[]":         []string{"1", "2"},
				"_keywords_state[]":        []string{"1", "2"},
				"_keywords_favorite[]":     []string{"1", "2"},
				"_keywords_description[]":  []string{"あああ", "🍺"},
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
