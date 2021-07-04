package zones_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestHistory(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "71B4D2B268744DBB8C3220C3A38C29E5",
		"results": [
			{
				"id": 1,
				"committed_at": "2021-06-20T10:50:37.396Z",
				"description": "commit 1",
				"operator": "user1"
				},
			{
				"id": 2,
				"committed_at": "2021-06-20T10:50:37.396Z",
				"description": "commit 2",
				"operator": "user2"
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "8BAB2071CF2C4509A93E485D7F368A71",
		"result": {
			"count": 2
		}
	}`)))
	atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:50:37.396Z")
	if err != nil {
		t.Fatal(err)
	}
	client := api.NewClient("", "http://localhost", nil)
	testcase := []zones.History{
		{
			Id:          1,
			CommittedAt: atTime,
			Description: "commit 1",
			Operator:    "user1",
		},
		{
			Id:          2,
			CommittedAt: atTime,
			Description: "commit 2",
			Operator:    "user2",
		},
	}
	// GET /zones/{ZoneId}/logs
	list := &zones.HistoryList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "71B4D2B268744DBB8C3220C3A38C29E5") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &zones.HistoryList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "71B4D2B268744DBB8C3220C3A38C29E5") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}

	// GET /zones/{ZoneId}/logs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "8BAB2071CF2C4509A93E485D7F368A71")
	}
}

func TestHistoryText(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/zone_histories/1/text", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "CC2BBC44169940CE96D7E8A99FA98958",
		"result": {
			"id": 1,
			"committed_at": "2021-06-20T10:50:37.396Z",
			"description": "commit 1",
			"operator": "user1",
			"text": "zone data"
		}
	}`)))
	atTime, err := types.ParseTime(time.RFC3339Nano, "2021-06-20T10:50:37.396Z")
	if err != nil {
		t.Fatal(err)
	}
	client := api.NewClient("", "http://localhost", nil)

	// GET /zones/{ZoneId}/zone_histories/{ZoneHistoryId}/text
	textTestCase := &zones.HistoryText{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		History: zones.History{
			Id:          1,
			CommittedAt: atTime,
			Description: "commit 1",
			Operator:    "user1",
		},
		Text: "zone data",
	}
	text := &zones.HistoryText{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		History: zones.History{
			Id: 1,
		},
	}
	reqId, err := client.Read(text)
	if assert.NoError(t, err) {
		assert.Equal(t, text, textTestCase)
		assert.Equal(t, reqId, "CC2BBC44169940CE96D7E8A99FA98958")
	}
	text = &zones.HistoryText{}
	if assert.NoError(t, text.SetParams("m1", 1)) {
		reqId, err := client.Read(text)
		if assert.NoError(t, err) {
			assert.Equal(t, text, textTestCase)
			assert.Equal(t, reqId, "CC2BBC44169940CE96D7E8A99FA98958")
		}
	}
}

func TestHistoryListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword zones.HistoryListSearchKeywords
		values  url.Values
	}{
		{
			zones.HistoryListSearchKeywords{
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
			zones.HistoryListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:    api.KeywordsString{"hogehoge", "🐰"},
				Description: api.KeywordsString{"🐇", "🍺"},
				Operator:    api.KeywordsString{"rabbit@example.jp", "SA0000000"},
			},
			/*
				_keywords_full_text[]
				_keywords_description[]
				_keywords_operator[]

			*/
			url.Values{
				"type":                    []string{"OR"},
				"offset":                  []string{"10"},
				"limit":                   []string{"100"},
				"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
				"_keywords_description[]": []string{"🐇", "🍺"},
				"_keywords_operator[]":    []string{"rabbit@example.jp", "SA0000000"},
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
