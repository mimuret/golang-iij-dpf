package testtool

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	. "github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ ClientInterface = &TestClient{}

type TestClient struct {
	Client         ClientInterface
	Transport      *httpmock.MockTransport
	RequestHeaders map[string]http.Header
	RequestBody    map[string]string

	ReadFunc         func(s api.Spec) (requestId string, err error)
	ListFunc         func(s api.ListSpec, keywords api.SearchParams) (requestId string, err error)
	ListAllFunc      func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error)
	CountFunc        func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error)
	UpdateFunc       func(s api.Spec, body interface{}) (requestId string, err error)
	CreateFunc       func(s api.Spec, body interface{}) (requestId string, err error)
	ApplyFunc        func(s api.Spec, body interface{}) (requestId string, err error)
	DeleteFunc       func(s api.Spec) (requestId string, err error)
	CancelFunc       func(s api.Spec) (requestId string, err error)
	WatchReadFunc    func(ctx context.Context, interval time.Duration, s api.Spec) error
	WatchListFunc    func(ctx context.Context, interval time.Duration, s api.ListSpec, keyword api.SearchParams) error
	WatchListAllFunc func(ctx context.Context, interval time.Duration, s api.CountableListSpec, keyword api.SearchParams) error
}

type ResponseSpec struct {
	Code int
	Spec api.Spec
	Err  *api.BadResponse
}

func NewTestClient(token, endpoint string, logger Logger) *TestClient {
	cl := NewClient(token, endpoint, logger)
	nop := &TestClient{
		RequestHeaders: make(map[string]http.Header),
		RequestBody:    make(map[string]string),
	}
	cl.Client.Transport = nop
	nop.Client = cl
	return nop
}

func (n *TestClient) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Body != nil {
		bs, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		n.RequestBody[req.URL.Path] = string(bs)
		n.RequestHeaders[req.URL.Path] = req.Header
	}
	return httpmock.DefaultTransport.RoundTrip(req)
}

func (n *TestClient) Read(s api.Spec) (requestId string, err error) {
	if n.ReadFunc == nil {
		if n.Client != nil {
			return n.Client.Read(s)
		}
		return "", nil
	}
	return n.ReadFunc(s)
}
func (n *TestClient) List(s api.ListSpec, keywords api.SearchParams) (requestId string, err error) {
	if n.ListFunc == nil {
		if n.Client != nil {
			return n.Client.List(s, keywords)
		}
		return "", nil
	}
	return n.ListFunc(s, keywords)
}
func (n *TestClient) ListAll(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
	if n.ListAllFunc == nil {
		if n.Client != nil {
			return n.Client.ListAll(s, keywords)
		}
		return "", nil
	}
	return n.ListAllFunc(s, keywords)
}
func (n *TestClient) Count(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
	if n.CountFunc == nil {
		if n.Client != nil {
			return n.Client.Count(s, keywords)
		}
		return "", nil
	}
	return n.CountFunc(s, keywords)
}
func (n *TestClient) Update(s api.Spec, body interface{}) (requestId string, err error) {
	if n.UpdateFunc == nil {
		if n.Client != nil {
			return n.Client.Update(s, body)
		}
		return "", nil
	}
	return n.UpdateFunc(s, body)
}
func (n *TestClient) Create(s api.Spec, body interface{}) (requestId string, err error) {
	if n.CreateFunc == nil {
		if n.Client != nil {
			return n.Client.Create(s, body)
		}
		return "", nil
	}
	return n.CreateFunc(s, body)
}
func (n *TestClient) Apply(s api.Spec, body interface{}) (requestId string, err error) {
	if n.ApplyFunc == nil {
		if n.Client != nil {
			return n.Client.Apply(s, body)
		}
		return "", nil
	}
	return n.ApplyFunc(s, body)
}
func (n *TestClient) Delete(s api.Spec) (requestId string, err error) {
	if n.DeleteFunc == nil {
		if n.Client != nil {
			return n.Client.Delete(s)
		}
		return "", nil
	}
	return n.DeleteFunc(s)
}
func (n *TestClient) Cancel(s api.Spec) (requestId string, err error) {
	if n.CancelFunc == nil {
		if n.Client != nil {
			return n.Client.Cancel(s)
		}
		return "", nil
	}
	return n.CancelFunc(s)
}
func (n *TestClient) WatchRead(ctx context.Context, interval time.Duration, s api.Spec) error {
	if n.WatchReadFunc == nil {
		if n.Client != nil {
			return n.Client.WatchRead(ctx, interval, s)
		}
		return nil
	}
	return n.WatchReadFunc(ctx, interval, s)
}
func (n *TestClient) WatchList(ctx context.Context, interval time.Duration, s api.ListSpec, keyword api.SearchParams) error {
	if n.WatchListFunc == nil {
		if n.Client != nil {
			return n.Client.WatchList(ctx, interval, s, keyword)
		}
		return nil
	}
	return n.WatchListFunc(ctx, interval, s, keyword)
}
func (n *TestClient) WatchListAll(ctx context.Context, interval time.Duration, s api.CountableListSpec, keyword api.SearchParams) error {
	if n.WatchListAllFunc == nil {
		if n.Client != nil {
			return n.Client.WatchListAll(ctx, interval, s, keyword)
		}
		return nil
	}
	return n.WatchListAllFunc(ctx, interval, s, keyword)
}
