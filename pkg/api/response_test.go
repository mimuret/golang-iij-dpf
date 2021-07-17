package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ = Describe("response", func() {
	Context("BadResponse", func() {
		var (
			bad *api.BadResponse
		)
		Context("Error", func() {
			When("Auth error", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   400,
						ErrorType:    api.ErrorTypeParamaterError,
						ErrorMessage: "There are invalid parameters.",
						ErrorDetails: api.ErrorDetails{
							{Code: "invalid", Attribute: "access_token"},
						},
					}
				})
				It("return `Auth error`", func() {
					Expect(bad.Error()).Should(MatchRegexp("Auth error"))
				})
			})
			When("request format error", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   400,
						ErrorType:    api.ErrorTypeParamaterError,
						ErrorMessage: "JSON parse error occurred.",
					}
				})
				It("return `Invalid request format`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: ParameterError Message: JSON parse error occurred."))
				})
			})
			When("invalid paramaters.", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   400,
						ErrorType:    api.ErrorTypeParamaterError,
						ErrorMessage: "There are invalid parameters.",
						ErrorDetails: api.ErrorDetails{
							{Code: "invalid", Attribute: "name"},
							{Code: "notfound", Attribute: "system_id"},
						},
					}
				})
				It("return `ErrorType: ParameterError with Details`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: ParameterError Message: There are invalid parameters. Detail: invalid=name, notfound=system_id"))
				})
			})
			When("NotFound", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   404,
						ErrorType:    api.ErrorTypeNotFound,
						ErrorMessage: "Specified resource not found.",
					}
				})
				It("return `ErrorType: NotFound`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: NotFound Message: Specified resource not found."))
				})
			})
			When("TooManyRequests", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   429,
						ErrorType:    api.ErrorTypeTooManyRequests,
						ErrorMessage: "Too many requests.",
					}
				})
				It("return `TooManyRequests`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: TooManyRequests Message: Too many requests."))
				})
			})
			When("SystemError", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   500,
						ErrorType:    api.ErrorTypeSystemError,
						ErrorMessage: "System error occurred.",
					}
				})
				It("return `SystemError`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: SystemError Message: System error occurred."))
				})
			})
			When("GatewayTimeout", func() {
				BeforeEach(func() {
					bad = &api.BadResponse{
						StatusCode:   504,
						ErrorType:    api.ErrorTypeGatewayTimeout,
						ErrorMessage: "Gateway timeout.",
					}
				})
				It("return `GatewayTimeout`", func() {
					Expect(bad.Error()).Should(MatchRegexp("ErrorType: GatewayTimeout Message: Gateway timeout."))
				})
			})
		})
	})
})
