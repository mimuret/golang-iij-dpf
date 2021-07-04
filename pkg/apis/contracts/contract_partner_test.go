package contracts_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/contracts"
	"github.com/stretchr/testify/assert"
)

func TestContractPartner(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/contracts/f1/contract_partners", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0EF26E92E5D54352A78C9EA4551BE502",
		"results": [
			{
				"service_code": "svc001"
			},
			{
				"service_code": "svc002"
			}
		]
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := []contracts.ContractPartner{
		{"svc001"},
		{"svc002"},
	}
	// GET /contracts/{ContractId}/common_configs
	list := &contracts.ContractPartnerList{
		AttributeMeta: contracts.AttributeMeta{
			ContractId: "f1",
		},
	}
	reqId, err := client.List(list, nil)
	if assert.NoError(t, err) {
		if assert.Equal(t, reqId, "0EF26E92E5D54352A78C9EA4551BE502") {
			if assert.Len(t, list.Items, 2) {
				for i, tc := range testcase {
					assert.Equal(t, list.Items[i], tc)
				}
			}
		}
	}
	list = &contracts.ContractPartnerList{}
	if assert.NoError(t, list.SetParams("f1")) {
		reqId, err := client.List(list, nil)
		if assert.NoError(t, err) {
			if assert.Equal(t, reqId, "0EF26E92E5D54352A78C9EA4551BE502") {
				if assert.Len(t, list.Items, 2) {
					for i, tc := range testcase {
						assert.Equal(t, list.Items[i], tc)
					}
				}
			}
		}
	}
}
