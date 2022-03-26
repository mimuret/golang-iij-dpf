package testtool_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("response.go", func() {
	Context("GenReqId", func() {
		var reqId string
		BeforeEach(func() {
			reqId = testtool.GenReqId()
		})
		It("returns lowercase uuuid exclude hyphen", func() {
			Expect(reqId).To(MatchRegexp("^[a-f0-9]+$"))
		})
	})
	Context("GenReqId", func() {
		var (
			reqId string
			bs    []byte
		)
		BeforeEach(func() {
			reqId, bs = testtool.CreateAsyncResponse()
		})
		It("returns RequestId", func() {
			Expect(reqId).To(MatchRegexp("^[a-f0-9]+$"))
		})
		It("returns AsyncResponseJson", func() {
			Expect(bs).To(MatchJSON(`{
				"request_id": "` + reqId + `",
				"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/` + reqId + `"
			}`))
		})
	})
})
