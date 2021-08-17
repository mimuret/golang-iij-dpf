package testtool_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("net.go", func() {
	Context("MustParseIPNet", func() {
		var (
			ipnet *types.IPNet
		)
		When("normal", func() {
			BeforeEach(func() {
				ipnet = testtool.MustParseIPNet("192.168.0.0/16")
			})
			It("returns *types.IPNet", func() {
				Expect(ipnet).NotTo(BeNil())
			})
		})
		When("nivalid data", func() {
			It("raise panic", func() {
				Expect(func() {
					testtool.MustParseIPNet("192.168.a.0/16")
				}).To(Panic())
			})
		})
	})
})
