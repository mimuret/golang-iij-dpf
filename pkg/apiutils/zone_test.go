package apiutils_test

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apiutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("zone", func() {
	var (
		err      error
		c        *api.Client
		systemId string
	)
	BeforeEach(func() {
		httpmock.Reset()
		c = api.NewClient("token", "http://localhost", nil)
	})
	Context("ZonenameToSystemId", func() {
		When("failed to read", func() {
			BeforeEach(func() {
				systemId, err = apiutils.ZonenameToSystemId(c, "example.jp.")
			})
			It("return empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to search zone"))
				Expect(systemId).To(Equal(""))
			})
		})
		When("find", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones?_keywords_name%5B%5D=example.jp.", httpmock.NewBytesResponder(200, []byte(`{
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
							}
						]
					}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/count?_keywords_name%5B%5D=example.jp.", httpmock.NewBytesResponder(200, []byte(`{
						"request_id": "9C518C729E5541D999389C686FE8987D",
						"result": {
							"count": 1
						}
					}`)))
				systemId, err = apiutils.ZonenameToSystemId(c, "example.jp.")
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("return systemid", func() {
				Expect(err).To(Succeed())
				Expect(systemId).To(Equal("m1"))
			})
		})
		When("not find", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones?_keywords_name%5B%5D=example.jp.", httpmock.NewBytesResponder(200, []byte(`{
							"request_id": "2C18922432DC48D485613F5383A7ED8E",
							"results": []
						}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/zones/count?_keywords_name%5B%5D=example.jp.", httpmock.NewBytesResponder(200, []byte(`{
							"request_id": "9C518C729E5541D999389C686FE8987D",
							"result": {
								"count": 0
							}
						}`)))
				systemId, err = apiutils.ZonenameToSystemId(c, "example.jp.")
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("not return systemId", func() {
				Expect(err).To(Succeed())
				Expect(systemId).To(Equal(""))
			})
		})
	})
})
