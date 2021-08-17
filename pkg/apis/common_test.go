package apis_test

import (
	"testing"

	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "apis Suite")
}

var _ = Describe("apis", func() {
	Context("SetPathParams", func() {
		var (
			err        error
			strValue   string
			int64Value int64
		)
		BeforeEach(func() {
			int64Value = 0
			strValue = ""
		})
		When("args is nothing", func() {
			BeforeEach(func() {
				err = apis.SetPathParams(nil, &strValue, &int64Value)
			})
			It("nothing to do", func() {
				Expect(err).To(Succeed())
			})
		})
		When("args len not equals to ids", func() {
			BeforeEach(func() {
				err = apis.SetPathParams([]interface{}{"string"}, &strValue, &int64Value)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("SetPathParams: args need 2 items"))
			})
		})
		When("failed to cast int64", func() {
			BeforeEach(func() {
				err = apis.SetPathParams([]interface{}{"string"}, &int64Value)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to cast to int64"))
			})
		})
		When("failed to cast string", func() {
			BeforeEach(func() {
				err = apis.SetPathParams([]interface{}{10}, &strValue)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to cast to string"))
			})
		})
		When("ids is not *int64 or *string", func() {
			var v int32
			It("raise panic", func() {
				Expect(func() { apis.SetPathParams([]interface{}{10}, &v) }).To(Panic())
			})
		})
		When("success", func() {
			BeforeEach(func() {
				err = apis.SetPathParams([]interface{}{10, "hoge"}, &int64Value, &strValue)
			})
			It("return err", func() {
				Expect(err).To(Succeed())
				Expect(int64Value).To(Equal(int64(10)))
				Expect(strValue).To(Equal("hoge"))
			})
		})
	})
})
