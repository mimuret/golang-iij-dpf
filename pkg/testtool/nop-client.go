package testtool

import (
	"io/ioutil"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ ClientInterface = &NopClient{}

type NopClient struct {
	*Client
	Transport      *httpmock.MockTransport
	RequestHeaders map[string]http.Header
	RequestBody    map[string]string
	Response       map[string]ResponseSpec
}

type ResponseSpec struct {
	Code int
	Msg  interface{}
}

func NewNopClient(token, endpoint string, logger Logger) *NopClient {
	nop := &NopClient{
		Client:         NewClient(token, endpoint, logger),
		Transport:      httpmock.NewMockTransport(),
		RequestHeaders: make(map[string]http.Header),
		RequestBody:    make(map[string]string),
		Response:       make(map[string]ResponseSpec),
	}
	httpmock.ActivateNonDefault(nop.Client.Client)
	nop.Transport.RegisterNoResponder(func(req *http.Request) (*http.Response, error) {
		defer req.Body.Close()
		bs, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		nop.RequestBody[req.URL.Path] = string(bs)
		nop.RequestHeaders[req.URL.Path] = req.Header

		if response, ok := nop.Response[req.URL.Path]; ok {
			return httpmock.NewJsonResponse(response.Code, response.Msg)
		}
		return httpmock.NewJsonResponse(200, AsyncResponse{
			ResponseCommon: ResponseCommon{
				RequestId: "662F71C4C06441ACB72A29DD9BFFF089",
			},
			JobsUrl: "http://api.dns-platform.jp/dpf/v1/jobs/662F71C4C06441ACB72A29DD9BFFF089",
		})
	})
	return nop
}
