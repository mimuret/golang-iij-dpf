package contracts_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	"github.com/stretchr/testify/assert"
)

func TestQps(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/qps/histories", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "3649092EAD70427A854CE933973AFFF8",
		"results": [
			{
				"service_code": "f00001",
				"name": "hoge",
				"values": [
					{
						"month": "2021-01",
						"qps": 100
					},
					{
						"month": "2021-02",
						"qps": 100
					}
				]
			},
			{
				"service_code": "f00002",
				"name": "hogehoge",
				"values": [
					{
						"month": "2021-07",
						"qps": 20
					},
					{
						"month": "2021-09",
						"qps": 100
					}
				]
			}
		]
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	testcase := []contracts.QpsHistory{
		{
			ServiceCode: "f00001",
			Name:        "hoge",
			Values: []contracts.QpsValue{
				{
					Month: "2021-01",
					Qps:   100,
				},
				{
					Month: "2021-02",
					Qps:   100,
				},
			},
		},
		{
			ServiceCode: "f00002",
			Name:        "hogehoge",
			Values: []contracts.QpsValue{
				{
					Month: "2021-07",
					Qps:   20,
				},
				{
					Month: "2021-09",
					Qps:   100,
				},
			},
		},
	}
	// GET /contracts/{ContractId}/qps/histories
	list := &contracts.QpsHistoryList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "3649092EAD70427A854CE933973AFFF8") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &contracts.QpsHistoryList{}
	if assert.NoError(t, list.SetParams("f1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "3649092EAD70427A854CE933973AFFF8") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}
}
