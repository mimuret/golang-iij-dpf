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
	Context("SetParams", func() {
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
				err = apis.SetParams(nil, &strValue, &int64Value)
			})
			It("nothing to do", func() {
				Expect(err).To(Succeed())
			})
		})
		When("args len not equals to ids", func() {
			BeforeEach(func() {
				err = apis.SetParams([]interface{}{"string"}, &strValue, &int64Value)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("SetParams: args need 2 items"))
			})
		})
		When("failed to cast int64", func() {
			BeforeEach(func() {
				err = apis.SetParams([]interface{}{"string"}, &int64Value)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to cast to int64"))
			})
		})
		When("failed to cast string", func() {
			BeforeEach(func() {
				err = apis.SetParams([]interface{}{10}, &strValue)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to cast to string"))
			})
		})
		When("success", func() {
			BeforeEach(func() {
				err = apis.SetParams([]interface{}{10, "hoge"}, &int64Value, &strValue)
			})
			It("return err", func() {
				Expect(err).To(Succeed())
				Expect(int64Value).To(Equal(int64(10)))
				Expect(strValue).To(Equal("hoge"))
			})
		})
	})
})
