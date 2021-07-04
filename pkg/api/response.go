package api

import (
	"encoding/json"
	"strings"
)

type ResponseCommon struct {
	RequestId string `read:"request_id"`
}

type RawResponse struct {
	ResponseCommon `read:",inline"`
	Result         json.RawMessage `read:"result"`
	Results        json.RawMessage `read:"results"`
}

type BadResponse struct {
	ResponseCommon `read:",inline"`
	ErrorType      string       `read:"error_type"`
	ErrorMessage   string       `read:"error_message"`
	ErrorDetails   ErrorDetails `read:"error_details"`
}

func (r *BadResponse) Error() string {
	errorDetail := ""
	if len(r.ErrorDetails) > 0 {
		errorDetail = " Detail: "
		if IsAuthError(r) {
			return "Auth error"
		} else if IsInvalidSchema(r) {
			return "Invalid request format"
		} else {
			errorDetail += r.ErrorDetails.Error()
		}
	}
	return "ErrorType: " + r.ErrorType + " Message: " + r.ErrorMessage + errorDetail
}

type ErrorDetails []ErrorDetail

func (e ErrorDetails) Error() string {
	res := []string{}
	for _, detail := range e {
		res = append(res, detail.Error())
	}
	return strings.Join(res, ", ")
}

type ErrorDetail struct {
	Code      string `read:"code"`
	Attribute string `read:"attribute"`
}

func (e ErrorDetail) Error() string {
	return e.Code + "=" + e.Attribute
}

type CountResponse struct {
	ResponseCommon `read:",inline"`
	Result         Count `read:"result"`
}

type Count struct {
	Count int32 `read:"count" json:"-"`
}

func (c *Count) SetCount(v int32) { c.Count = v }
func (c *Count) GetCount() int32  { return c.Count }

type AsyncResponse struct {
	ResponseCommon `read:",inline"`
	JobsUrl        string `read:"jobs_url"`
}
