package zones_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/zones"
	"github.com/stretchr/testify/assert"
)

func TestManagedDNS(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/m1/managed_dns_servers", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "F4F8591B01B84EFBBA0AD17029DB7DA4",
		"results": [
			"ns1.example.jp.",
			"ns1.example.net.",
			"ns1.example.com."
		]
	}`)))

	client := api.NewClient("", "http://localhost", nil)
	testcase := zones.ManagedDnsList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
		Items: []string{
			"ns1.example.jp.",
			"ns1.example.net.",
			"ns1.example.com.",
		},
	}
	s := zones.ManagedDnsList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: "m1",
		},
	}
	// GET /zones/{ZoneId}/managed_dns_servers
	reqId, err := client.List(&s, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "F4F8591B01B84EFBBA0AD17029DB7DA4")
		assert.Equal(t, s, testcase)
	}
	s = zones.ManagedDnsList{}
	if assert.NoError(t, s.SetParams("m1")) {
		reqId, err := client.List(&s, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, "F4F8591B01B84EFBBA0AD17029DB7DA4")
			assert.Equal(t, s, testcase)
		}
	}

}
