package api_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ = Describe("errors.go", func() {
	Context("IsStatusCode", func() {
		When("err is not BadResponse", func() {
			It("returns false", func() {
				Expect(api.IsStatusCode(errors.New(""), 400)).To(BeFalse())
			})
		})
		When("err not BadResponse, not match value", func() {
			It("returns false", func() {
				Expect(api.IsStatusCode(&api.BadResponse{StatusCode: 200}, 400)).To(BeFalse())
			})
		})
		When("err not BadResponse, match value", func() {
			It("returns true", func() {
				Expect(api.IsStatusCode(&api.BadResponse{StatusCode: 200}, 200)).To(BeTrue())
			})
		})
	})
	Context("IsErrType", func() {
		When("err is not BadResponse", func() {
			It("returns false", func() {
				Expect(api.IsErrType(errors.New(""), "hoge")).To(BeFalse())
			})
		})
		When("err not BadResponse, not match value", func() {
			It("returns false", func() {
				Expect(api.IsErrType(&api.BadResponse{ErrorType: "book"}, "hoge")).To(BeFalse())
			})
		})
		When("err not BadResponse, match value", func() {
			It("returns true", func() {
				Expect(api.IsErrType(&api.BadResponse{ErrorType: "hoge"}, "hoge")).To(BeTrue())
			})
		})
	})
	Context("IsErrMsg", func() {
		When("err is not BadResponse", func() {
			It("returns false", func() {
				Expect(api.IsErrMsg(errors.New(""), "hoge")).To(BeFalse())
			})
		})
		When("err not BadResponse, not match value", func() {
			It("returns false", func() {
				Expect(api.IsErrMsg(&api.BadResponse{ErrorMessage: "book"}, "hoge")).To(BeFalse())
			})
		})
		When("err not BadResponse, match value", func() {
			It("returns true", func() {
				Expect(api.IsErrMsg(&api.BadResponse{ErrorMessage: "hoge"}, "hoge")).To(BeTrue())
			})
		})
	})
	Context("IsErrorCode", func() {
		When("err is not BadResponse", func() {
			It("returns false", func() {
				Expect(api.IsErrorCode(errors.New(""), "hoge")).To(BeFalse())
			})
		})
		When("err not BadResponse, not match value", func() {
			It("returns false", func() {
				ok, _ := api.IsErrorCode(&api.BadResponse{ErrorDetails: api.ErrorDetails{{Code: "code", Attribute: "ruby"}}}, "hoge")
				Expect(ok).To(BeFalse())
			})
		})
		When("err not BadResponse, match value", func() {
			It("returns true", func() {
				ok, attr := api.IsErrorCode(&api.BadResponse{ErrorDetails: api.ErrorDetails{{Code: "code", Attribute: "ruby"}}}, "code")
				Expect(ok).To(BeTrue())
				Expect(attr).To(Equal("ruby"))
			})
		})
	})
	Context("IsErrorCodeAttribute", func() {
		When("err is not BadResponse", func() {
			It("returns false", func() {
				Expect(api.IsErrorCodeAttribute(errors.New(""), "code", "attribute")).To(BeFalse())
			})
		})
		When("err not BadResponse, not match details", func() {
			It("returns false", func() {
				Expect(api.IsErrorCodeAttribute(&api.BadResponse{ErrorDetails: api.ErrorDetails{{Code: "code", Attribute: "ruby"}}}, "code", "attribute")).To(BeFalse())
			})
		})
		When("err not BadResponse, include details", func() {
			It("returns true", func() {
				Expect(api.IsErrorCodeAttribute(&api.BadResponse{ErrorDetails: api.ErrorDetails{{Code: "book", Attribute: "one"}, {Code: "code", Attribute: "attribute"}}}, "code", "attribute")).To(BeTrue())
			})
		})
	})
	Context("Bad", func() {
		When("err is not BadResponse", func() {
			It("return false", func() {
				ok := api.IsBadResponse(fmt.Errorf("failed"), nil)
				Expect(ok).To(BeFalse())
			})
		})
		When("err is BadResponse, func not set", func() {
			It("return true", func() {
				ok := api.IsBadResponse(&api.BadResponse{}, nil)
				Expect(ok).To(BeTrue())
			})
		})
		When("err is BadResponse, func set", func() {
			It("return func result", func() {
				ok := api.IsBadResponse(&api.BadResponse{}, func(b *api.BadResponse) bool { return true })
				Expect(ok).To(BeTrue())
				ok = api.IsBadResponse(&api.BadResponse{}, func(b *api.BadResponse) bool { return false })
				Expect(ok).To(BeFalse())
			})
		})
	})
	Context("IsAuthError", func() {
		When("status code is 400, ErrType is ParameterError, ErrorDtail include invalid=access_token", func() {
			It("return true", func() {
				Expect(api.IsAuthError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "access_token"}},
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsAuthError(&api.BadResponse{
					StatusCode:   200,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "access_token"}},
				})).To(BeFalse())
				Expect(api.IsAuthError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "NotFound",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "access_token"}},
				})).To(BeFalse())
				Expect(api.IsAuthError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "schema"}},
				})).To(BeFalse())
			})
		})
	})
	Context("IsRequestFormatError", func() {
		When("status code is 400, ErrType is ParameterError, ErrMsg `JSON parse error occurred`", func() {
			It("return true", func() {
				Expect(api.IsRequestFormatError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorMessage: "JSON parse error occurred.",
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsRequestFormatError(&api.BadResponse{
					StatusCode:   200,
					ErrorType:    "ParameterError",
					ErrorMessage: "JSON parse error occurred.",
				})).To(BeFalse())
				Expect(api.IsRequestFormatError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "NotFound",
					ErrorMessage: "JSON parse error occurred.",
				})).To(BeFalse())
				Expect(api.IsRequestFormatError(&api.BadResponse{
					StatusCode: 400,
					ErrorType:  "ParameterError",
				})).To(BeFalse())
			})
		})
	})
	Context("IsParameterError", func() {
		When("status code is 400, ErrType ParameterError, NotAuthError, NotRequestFormatError", func() {
			It("return true", func() {
				Expect(api.IsParameterError(&api.BadResponse{
					StatusCode: 400,
					ErrorType:  "ParameterError",
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsParameterError(&api.BadResponse{
					StatusCode: 200,
					ErrorType:  "ParameterError",
				})).To(BeFalse())
				Expect(api.IsParameterError(&api.BadResponse{
					StatusCode: 400,
					ErrorType:  "NotFound",
				})).To(BeFalse())
				Expect(api.IsParameterError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorMessage: "JSON parse error occurred.",
				})).To(BeFalse())
				Expect(api.IsParameterError(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "access_token"}},
				})).To(BeFalse())
			})
		})
	})
	Context("IsNotFound", func() {
		When("status code is 404", func() {
			It("return true", func() {
				Expect(api.IsNotFound(&api.BadResponse{
					StatusCode: 404,
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsNotFound(&api.BadResponse{
					StatusCode: 200,
				})).To(BeFalse())
			})
		})
	})
	Context("IsTooManyRequests", func() {
		When("status code is 429", func() {
			It("return true", func() {
				Expect(api.IsTooManyRequests(&api.BadResponse{
					StatusCode: 429,
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsTooManyRequests(&api.BadResponse{
					StatusCode: 200,
				})).To(BeFalse())
			})
		})
	})
	Context("IsSystemError", func() {
		When("status code is 500", func() {
			It("return true", func() {
				Expect(api.IsSystemError(&api.BadResponse{
					StatusCode: 500,
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsSystemError(&api.BadResponse{
					StatusCode: 200,
				})).To(BeFalse())
			})
		})
	})
	Context("IsGatewayTimeout", func() {
		When("status code is 504", func() {
			It("return true", func() {
				Expect(api.IsGatewayTimeout(&api.BadResponse{
					StatusCode: 504,
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsGatewayTimeout(&api.BadResponse{
					StatusCode: 200,
				})).To(BeFalse())
			})
		})
	})
	Context("IsInvalidSchema", func() {
		When("IsInvalidSchema and ErrorDtail include invalid=schema", func() {
			It("return true", func() {
				Expect(api.IsInvalidSchema(&api.BadResponse{
					StatusCode:   400,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "schema"}},
				})).To(BeTrue())
			})
		})
		When("other", func() {
			It("return false", func() {
				Expect(api.IsInvalidSchema(&api.BadResponse{
					StatusCode: 400,
					ErrorType:  "ParameterError",
				})).To(BeFalse())
				Expect(api.IsInvalidSchema(&api.BadResponse{
					StatusCode:   404,
					ErrorType:    "ParameterError",
					ErrorDetails: api.ErrorDetails{{Code: "invalid", Attribute: "schema"}},
				})).To(BeFalse())
			})
		})
	})
})
