package core_test

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
	"github.com/stretchr/testify/assert"
)

func TestLogListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword core.LogListSearchKeywords
		values  url.Values
	}{
		{
			core.LogListSearchKeywords{
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
			core.LogListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:  api.KeywordsString{"hogehoge", "🐰"},
				LogType:   api.KeywordsString{"service", "record", "dnssec"},
				Operator:  api.KeywordsString{"rabbit@example.jp", "SA0000000"},
				Operation: api.KeywordsString{"updating_default_ttl", "dismiss_zone_edits"},
				Target:    api.KeywordsString{"hoge", "fuga"},
				Detail:    api.KeywordsString{"🐇", "🍺"},
				RequestId: api.KeywordsString{"f02fe1e1404140cab93c8f7af26081b7", "f098f808cde249d48a885490f7622df9"},
				Status:    core.KeywordsLogStatus{core.LogStatusStart, core.LogStatusSuccess, core.LogStatusFailure, core.LogStatusRetry},
			},
			/*
				_keywords_full_text[]
				_keywords_log_type[]
				_keywords_operator[]
				_keywords_operation[]
				_keywords_target[]
				_keywords_detail[]
				_keywords_request_id[]
				_keywords_status[]

			*/
			url.Values{
				"type":                   []string{"OR"},
				"offset":                 []string{"10"},
				"limit":                  []string{"100"},
				"_keywords_full_text[]":  []string{"hogehoge", "🐰"},
				"_keywords_log_type[]":   []string{"service", "record", "dnssec"},
				"_keywords_operator[]":   []string{"rabbit@example.jp", "SA0000000"},
				"_keywords_operation[]":  []string{"updating_default_ttl", "dismiss_zone_edits"},
				"_keywords_target[]":     []string{"hoge", "fuga"},
				"_keywords_detail[]":     []string{"🐇", "🍺"},
				"_keywords_request_id[]": []string{"f02fe1e1404140cab93c8f7af26081b7", "f098f808cde249d48a885490f7622df9"},
				"_keywords_status[]":     []string{"start", "success", "failure", "retry"},
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
