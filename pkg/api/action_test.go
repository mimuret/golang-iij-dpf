package api_test

import (
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Action", func() {
	Context("Actions", func() {
		It("returns actions", func() {
			Expect(api.Actions()).NotTo(BeEmpty())
		})
	})
	Context("ToMethod", func() {
		When("ActionCreate", func() {
			It("returns http.MethodPost", func() {
				Expect(api.ActionCreate.ToMethod()).To(Equal(http.MethodPost))
			})
		})
		When("ActionRead", func() {
			It("returns http.MethodGet", func() {
				Expect(api.ActionRead.ToMethod()).To(Equal(http.MethodGet))
			})
		})
		When("ActionCreate", func() {
			It("returns http.MethodGet", func() {
				Expect(api.ActionList.ToMethod()).To(Equal(http.MethodGet))
			})
		})
		When("ActionUpdate", func() {
			It("returns http.MethodPatch", func() {
				Expect(api.ActionUpdate.ToMethod()).To(Equal(http.MethodPatch))
			})
		})
		When("ActionDelete", func() {
			It("returns http.MethodDelete", func() {
				Expect(api.ActionDelete.ToMethod()).To(Equal(http.MethodDelete))
			})
		})
		When("ActionCount", func() {
			It("returns http.MethodGet", func() {
				Expect(api.ActionCount.ToMethod()).To(Equal(http.MethodGet))
			})
		})
		When("ActionCancel", func() {
			It("returns http.MethodDelete", func() {
				Expect(api.ActionCancel.ToMethod()).To(Equal(http.MethodDelete))
			})
		})
		When("ActionApply", func() {
			It("returns http.MethodPatch", func() {
				Expect(api.ActionApply.ToMethod()).To(Equal(http.MethodPatch))
			})
		})
		When("OTHER", func() {
			It("returns empty", func() {
				Expect(api.Action("hoge").ToMethod()).To(Equal(""))
			})
		})
	})
})
