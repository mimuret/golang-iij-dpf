package zones_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestReord(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "743C1E2428264E368D37233049AD49B5",
		"result": {
      "id": "r1",
      "name": "www.example.jp.",
      "ttl": 30,
      "rrtype": "A",
      "rdata": [
        {
          "value": "192.168.1.1"
        },
        {
          "value": "192.168.1.2"
        }
      ],
      "state": 0,
      "description": "SERVER(IPv4)",
      "operator": "user1"
		}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "6B69B707901B479AAF83477B31340FEE",
		"result": {
      "id": "r2",
      "name": "www.example.jp.",
      "ttl": 30,
      "rrtype": "AAAA",
      "rdata": [
        {
          "value": "2001:db8::1"
        },
        {
          "value": "2001:db8::2"
        }
      ],
      "state": 1,
      "description": "SERVER(IPv6)",
      "operator": "user1"
		}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0987808B8D38403C8A1FBE66D72D68F7",
		"results": [
			{
				"id": "r1",
				"name": "www.example.jp.",
				"ttl": 30,
				"rrtype": "A",
				"rdata": [
					{
						"value": "192.168.1.1"
					},
					{
						"value": "192.168.1.2"
					}
				],
				"state": 0,
				"description": "SERVER(IPv4)",
				"operator": "user1"
			},
			{
				"id": "r2",
				"name": "www.example.jp.",
				"ttl": 30,
				"rrtype": "AAAA",
				"rdata": [
					{
						"value": "2001:db8::1"
					},
					{
						"value": "2001:db8::2"
					}
				],
				"state": 1,
				"description": "SERVER(IPv6)",
				"operator": "user1"
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "1DD94C67E9DA4391AB028914F8ACAB90",
		"result": {
			"count": 2
		}
	}`)))

	httpmock.RegisterResponder(http.MethodPost, "http://localhost/zones/m1/records", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "CB62C4812BC14A669983FAF9F56066C5",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/CB62C4812BC14A669983FAF9F56066C5"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "2A2D7BB22EB443F1855DA8E498035434",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))
	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "2A2D7BB22EB443F1855DA8E498035434",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/CC84F36FC71749E4A50BD208DD20E535"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "8597D1E5B2604F3C853F0A4BAF84F901",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r2", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "8597D1E5B2604F3C853F0A4BAF84F901",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/CC84F36FC71749E4A50BD208DD20E535"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r1/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "314ACF75BD3C4D968C28ACB03AC17BEB",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0CAC3AEFB6334ECCA7B70AF76D73508B"
	}`)))
	httpmock.RegisterResponder(http.MethodDelete, "http://localhost/zones/m1/records/r2/changes", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "314ACF75BD3C4D968C28ACB03AC17BEB",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/CC84F36FC71749E4A50BD208DD20E535"
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Spec      zones.Record
	}{
		{
			RequestId: "743C1E2428264E368D37233049AD49B5",
			Spec: zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneId: "m1",
				},
				Id:     "r1",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeA,
				RData: []zones.RecordRDATA{
					{Value: "192.168.1.1"},
					{Value: "192.168.1.2"},
				},
				State:       zones.RecordStateApplied,
				Description: "SERVER(IPv4)",
				Operator:    "user1",
			},
		},
		{
			RequestId: "6B69B707901B479AAF83477B31340FEE",
			Spec: zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneId: "m1",
				},
				Id:     "r2",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeAAAA,
				RData: []zones.RecordRDATA{
					{Value: "2001:db8::1"},
					{Value: "2001:db8::2"},
				},
				State:       zones.RecordStateToBeAdded,
				Description: "SERVER(IPv6)",
				Operator:    "user1",
			},
		},
	}
	// GET /zones/{ZoneId}/records
	list := &zones.RecordList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "0987808B8D38403C8A1FBE66D72D68F7") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}
	list = &zones.RecordList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "0987808B8D38403C8A1FBE66D72D68F7") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc.Spec)
					}
				}
			}
		}
	}

	// POST /zones/{ZoneId}/records
	for _, tc := range testcase {
		reqId, err := client.Create(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "CB62C4812BC14A669983FAF9F56066C5")
		}
	}
	// GET /zones/{ZoneId}/records/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "1DD94C67E9DA4391AB028914F8ACAB90")
	}
	// GET /ones/{ZoneId}}/records/{RecordId}
	for _, tc := range testcase {
		s := zones.Record{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Id: tc.Spec.Id,
		}
		reqId, err := client.Read(&s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s, tc.Spec)
		}
	}
	// Patch /zones/{ZoneId}/records/{RecordId}
	for _, tc := range testcase {
		reqId, err := client.Update(&tc.Spec, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "2A2D7BB22EB443F1855DA8E498035434")
		}
	}
	// Delete /zones/{ZoneId}/records/{RecordId}
	for _, tc := range testcase {
		reqId, err := client.Delete(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "8597D1E5B2604F3C853F0A4BAF84F901")
		}
	}
	// Delete /zones/{ZoneId}/records/{RecordId}/changes
	for _, tc := range testcase {
		reqId, err := client.Cancel(&tc.Spec)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "314ACF75BD3C4D968C28ACB03AC17BEB")
		}
	}
	createTestCase := []map[string]interface{}{
		{
			"name":   "www.example.jp.",
			"ttl":    30,
			"rrtype": "A",
			"rdata": []map[string]interface{}{
				{"value": "192.168.1.1"},
				{"value": "192.168.1.2"},
			},
			"description": "SERVER(IPv4)",
		},
		{
			"name":   "www.example.jp.",
			"ttl":    30,
			"rrtype": "AAAA",
			"rdata": []map[string]interface{}{
				{"value": "2001:db8::1"},
				{"value": "2001:db8::2"},
			},
			"description": "SERVER(IPv6)",
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
			"ttl": 30,
			"rdata": []map[string]interface{}{
				{"value": "192.168.1.1"},
				{"value": "192.168.1.2"},
			},
			"description": "SERVER(IPv4)",
		},
		{
			"ttl": 30,
			"rdata": []map[string]interface{}{
				{"value": "2001:db8::1"},
				{"value": "2001:db8::2"},
			},
			"description": "SERVER(IPv6)",
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

func TestCurrentRecordList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/currents", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "16FF10FA0E68496CBCD62E618BCCD665",
		"results": [
			{
				"id": "r1",
				"name": "www.example.jp.",
				"ttl": 30,
				"rrtype": "A",
				"rdata": [
					{
						"value": "192.168.1.1"
					},
					{
						"value": "192.168.1.2"
					}
				],
				"state": 0,
				"description": "SERVER(IPv4)",
				"operator": "user1"
			},
			{
				"id": "r2",
				"name": "www.example.jp.",
				"ttl": 30,
				"rrtype": "AAAA",
				"rdata": [
					{
						"value": "2001:db8::1"
					},
					{
						"value": "2001:db8::2"
					}
				],
				"state": 1,
				"description": "SERVER(IPv6)",
				"operator": "user1"
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/currents/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "36542C85D8C64ED7B8E20A8E8DA72A80",
		"result": {
			"count": 2
		}
	}`)))
	testcase := []struct {
		RequestId string
		Spec      zones.Record
	}{
		{
			RequestId: "743C1E2428264E368D37233049AD49B5",
			Spec: zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneId: "m1",
				},
				Id:     "r1",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeA,
				RData: []zones.RecordRDATA{
					{Value: "192.168.1.1"},
					{Value: "192.168.1.2"},
				},
				State:       zones.RecordStateApplied,
				Description: "SERVER(IPv4)",
				Operator:    "user1",
			},
		},
		{
			RequestId: "6B69B707901B479AAF83477B31340FEE",
			Spec: zones.Record{
				AttributeMeta: zones.AttributeMeta{
					ZoneId: "m1",
				},
				Id:     "r2",
				Name:   "www.example.jp.",
				TTL:    30,
				RRType: zones.TypeAAAA,
				RData: []zones.RecordRDATA{
					{Value: "2001:db8::1"},
					{Value: "2001:db8::2"},
				},
				State:       zones.RecordStateToBeAdded,
				Description: "SERVER(IPv6)",
				Operator:    "user1",
			},
		},
	}
	client := api.NewClient("", "http://localhost", nil)
	// GET /zones/{ZoneId}/records/currents
	list := &zones.CurrentRecordList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "16FF10FA0E68496CBCD62E618BCCD665") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc.Spec)
				}
			}
		}
	}

	list = &zones.CurrentRecordList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "16FF10FA0E68496CBCD62E618BCCD665") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc.Spec)
					}
				}
			}
		}
	}
	// GET /zones/{ZoneId}/records/currents/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(2))
		assert.Equal(t, reqId, "36542C85D8C64ED7B8E20A8E8DA72A80")
	}
}
func TestRecordDiff(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/diffs", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "C36898E9087C4E2F85C94D146BC6FEDD",
		"results": [
			{
				"new": {
					"id": "r1",
					"name": "www.example.jp.",
					"ttl": 30,
					"rrtype": "A",
					"rdata": [
						{
							"value": "192.168.1.1"
						},
						{
							"value": "192.168.1.2"
						}
					],
					"state": 0,
					"description": "SERVER(IPv4)",
					"operator": "user1"	
				}
			},
			{
				"old": {
					"id": "r2",
					"name": "www.example.jp.",
					"ttl": 30,
					"rrtype": "AAAA",
					"rdata": [
						{
							"value": "2001:db8::1"
						},
						{
							"value": "2001:db8::2"
						}
					],
					"state": 1,
					"description": "SERVER(IPv6)",
					"operator": "user1"
				}
			},
			{
				"new": {
					"id": "r3",
					"name": "www.example.jp.",
					"ttl": 300,
					"rrtype": "TXT",
					"rdata": [
						{
							"value": "new"
						}
					],
					"state": 1,
					"description": "SERVER(TXT)",
					"operator": "user1"
				},
				"old": {
					"id": "r3",
					"name": "www.example.jp.",
					"ttl": 300,
					"rrtype": "TXT",
					"rdata": [
						{
							"value": "old"
						}
					],
					"state": 1,
					"description": "SERVER(TXT)",
					"operator": "user1"
				}
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/records/diffs/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "EE547488A8B74999AD630664802B4E25",
		"result": {
			"count": 3
		}
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	// GET /zones/{ZoneId}/records/diffs
	tc := &zones.RecordDiffList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Items: []zones.RecordDiff{
			{
				New: &zones.Record{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Id:     "r1",
					Name:   "www.example.jp.",
					TTL:    30,
					RRType: zones.TypeA,
					RData: []zones.RecordRDATA{
						{Value: "192.168.1.1"},
						{Value: "192.168.1.2"},
					},
					State:       zones.RecordStateApplied,
					Description: "SERVER(IPv4)",
					Operator:    "user1",
				},
			},
			{
				Old: &zones.Record{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Id:     "r2",
					Name:   "www.example.jp.",
					TTL:    30,
					RRType: zones.TypeAAAA,
					RData: []zones.RecordRDATA{
						{Value: "2001:db8::1"},
						{Value: "2001:db8::2"},
					},
					State:       zones.RecordStateToBeAdded,
					Description: "SERVER(IPv6)",
					Operator:    "user1",
				},
			},
			{
				New: &zones.Record{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Id:     "r3",
					Name:   "www.example.jp.",
					TTL:    300,
					RRType: zones.TypeTXT,
					RData: []zones.RecordRDATA{
						{Value: "new"},
					},
					State:       zones.RecordStateToBeAdded,
					Description: "SERVER(TXT)",
					Operator:    "user1",
				},
				Old: &zones.Record{
					AttributeMeta: zones.AttributeMeta{
						ZoneId: "m1",
					},
					Id:     "r3",
					Name:   "www.example.jp.",
					TTL:    300,
					RRType: zones.TypeTXT,
					RData: []zones.RecordRDATA{
						{Value: "old"},
					},
					State:       zones.RecordStateToBeAdded,
					Description: "SERVER(TXT)",
					Operator:    "user1",
				},
			},
		},
	}
	list := &zones.RecordDiffList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}

	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "C36898E9087C4E2F85C94D146BC6FEDD")
		assert.Equal(t, tc, list)
	}

	list = &zones.RecordDiffList{}
	if assert.NoError(t, list.SetParams("m1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "C36898E9087C4E2F85C94D146BC6FEDD")
			assert.Equal(t, tc, list)
		}
	}
	// GET /zones/{ZoneId}/records/diffs/count
	reqId, err = client.Count(list, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, list.GetCount(), int32(3))
		assert.Equal(t, reqId, "EE547488A8B74999AD630664802B4E25")
	}
}

func TestRecordListSearchKeywords(t *testing.T) {
	testcase := []struct {
		keyword zones.RecordListSearchKeywords
		values  url.Values
	}{
		{
			zones.RecordListSearchKeywords{
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
			zones.RecordListSearchKeywords{
				CommonSearchParams: api.CommonSearchParams{
					Type:   api.SearchTypeOR,
					Offset: int32(10),
					Limit:  int32(100),
				},
				FullText:    api.KeywordsString{"hogehoge", "🐰"},
				Name:        api.KeywordsString{"example.jp.", "128/0.0.168.192.in-addr.arpa."},
				TTL:         []int32{300, 2147483647},
				RRType:      zones.KeywordsType{"A", "CNAME", "HTTPS"},
				RData:       api.KeywordsString{"192.168.0.1", "\"hogehoge\""},
				Description: api.KeywordsString{"🐇", "🍺"},
				Operator:    api.KeywordsString{"rabbit@example.jp", "SA0000000"},
			},
			/*
				_keywords_full_text[]
				_keywords_name[]
				_keywords_ttl[]
				_keywords_rrtype[]
				_keywords_rdata[]
				_keywords_description[]
				_keywords_operator[]

			*/
			url.Values{
				"type":                    []string{"OR"},
				"offset":                  []string{"10"},
				"limit":                   []string{"100"},
				"_keywords_full_text[]":   []string{"hogehoge", "🐰"},
				"_keywords_name[]":        []string{"example.jp.", "128/0.0.168.192.in-addr.arpa."},
				"_keywords_ttl[]":         []string{"300", "2147483647"},
				"_keywords_rrtype[]":      []string{"A", "CNAME", "HTTPS"},
				"_keywords_rdata[]":       []string{"192.168.0.1", "\"hogehoge\""},
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

func TestRRDatas(t *testing.T) {
	mx11, _ := dns.NewRR("example.jp. 300 IN MX 1 mx1.example.jp.")
	mx12, _ := dns.NewRR("example.jp. 300 IN MX 2 mx2.example.jp.")
	mx13, _ := dns.NewRR("example.jp. 300 IN MX 3 mx3.example.jp.")

	testcase := []struct {
		rrs    []dns.RR
		rdatas zones.RecordRDATAs
	}{
		{
			[]dns.RR{mx11, mx12, mx13},
			zones.RecordRDATAs{
				{Value: "1 mx1.example.jp."},
				{Value: "2 mx2.example.jp."},
				{Value: "3 mx3.example.jp."},
			},
		},
	}
	for _, tc := range testcase {
		r := zones.Record{}
		r.AddRRs(tc.rrs)
		assert.Equal(t, r.RData, tc.rdatas)
	}
}

func TestType(t *testing.T) {
	assert.Equal(t, zones.TypeANAME.Uint16(), uint16(65280))
	assert.Equal(t, zones.TypeA.Uint16(), dns.TypeA)
	assert.Equal(t, zones.TypeA.String(), "A")
	assert.Equal(t, zones.TypeANAME.String(), "ANAME")
	assert.Equal(t, zones.Uint16ToType(dns.TypeA), zones.TypeA)
	assert.Equal(t, zones.Uint16ToType(65280), zones.TypeANAME)
}
