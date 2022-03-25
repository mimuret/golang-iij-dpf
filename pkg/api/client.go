package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

const DefaultEndpoint = "https://api.dns-platform.jp/dpf/v1"

type ClientInterface interface {
	Read(ctx context.Context, s Spec) (requestId string, err error)
	List(ctx context.Context, s ListSpec, keywords SearchParams) (requestId string, err error)
	ListAll(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestId string, err error)
	Count(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestId string, err error)
	Update(ctx context.Context, s Spec, body interface{}) (requestId string, err error)
	Create(ctx context.Context, s Spec, body interface{}) (requestId string, err error)
	Apply(ctx context.Context, s Spec, body interface{}) (requestId string, err error)
	Delete(ctx context.Context, s Spec) (requestId string, err error)
	Cancel(ctx context.Context, s Spec) (requestId string, err error)
	WatchRead(ctx context.Context, interval time.Duration, s Spec) error
	WatchList(ctx context.Context, interval time.Duration, s ListSpec, keyword SearchParams) error
	WatchListAll(ctx context.Context, interval time.Duration, s CountableListSpec, keyword SearchParams) error
}

var _ ClientInterface = &Client{}

type Client struct {
	Endpoint string
	Token    string
	logger   Logger

	Client       *http.Client
	LastRequest  *RequestInfo
	LastResponse *ResponseInfo
}

type RequestInfo struct {
	Method string
	Url    string
	Body   []byte
}

type ResponseInfo struct {
	Response *http.Response
	Body     []byte
}

func NewClient(token string, endpoint string, logger Logger) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	if logger == nil {
		logger = NewStdLogger(os.Stderr, "dpf-client", 0, 4)
	}
	return &Client{Endpoint: endpoint, Token: token, logger: logger, Client: http.DefaultClient}
}

func (c *Client) Do(ctx context.Context, spec Spec, action Action, body interface{}, params SearchParams) (requestId string, err error) {
	var (
		ok            bool
		r             io.Reader
		countableList CountableListSpec
	)
	if action == ActionCount {
		countableList, ok = spec.(CountableListSpec)
		if !ok {
			return "", fmt.Errorf("spec is not CountableListSpec")
		}
	}
	c.LastRequest = &RequestInfo{}
	c.LastResponse = nil
	// create URL
	method, path := spec.GetPathMethod(action)
	if path == "" {
		return "", fmt.Errorf("not support action %s", action)
	}
	c.LastRequest.Method = method
	url := c.Endpoint + path
	if params != nil {
		p, err := params.GetValues()
		if err != nil {
			return "", fmt.Errorf("failed to get search params: %w", err)
		}
		url += "?" + p.Encode()
	}
	c.LastRequest.Url = url
	c.logger.Debugf("method: %s request-url: %s", method, url)
	// make request body
	if body != nil {
		var (
			jsonBody []byte
			err      error
		)
		switch action {
		case ActionCreate:
			jsonBody, err = MarshalCreate(body)
		case ActionUpdate:
			jsonBody, err = MarshalUpdate(body)
		case ActionApply:
			jsonBody, err = MarshalApply(body)
		default:
			return "", fmt.Errorf("not support action `%s` with body request", action)
		}
		if err != nil {
			return "", fmt.Errorf("failed to encode body to json: %w", err)
		}
		c.logger.Tracef("request-body: `%s`", string(jsonBody))
		c.LastRequest.Body = jsonBody
		r = bytes.NewBuffer(jsonBody)
	}

	// make request
	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}

	// authorized
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")

	// request
	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("failed to get http response: %w", err)
	}
	defer resp.Body.Close()
	c.LastResponse = &ResponseInfo{
		Response: resp,
	}
	// get body
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to get http response body: %w", err)
	}
	c.LastResponse.Body = bs
	c.logger.Debugf("status-code: `%d`", resp.StatusCode)
	c.logger.Tracef("response-body: `%s`", string(bs))

	// if statiscode is error, response body type is BadResponse or Plantext
	if resp.StatusCode >= 400 {
		badRequest := &BadResponse{StatusCode: resp.StatusCode}
		if err := UnmarshalRead(bs, badRequest); err != nil {
			return "", fmt.Errorf("failed to request: status code: %d body: %s err: %w", resp.StatusCode, string(bs), err)
		}
		return badRequest.RequestId, badRequest
	}

	// parse raw response
	rawResponse := &RawResponse{}
	if err := UnmarshalRead(bs, rawResponse); err != nil {
		// maybe not executed
		return "", fmt.Errorf("failed to parse get response: %w", err)
	}

	if method == http.MethodGet {
		if countableList != nil {
			// ActionCount
			count := &Count{}
			if err := UnmarshalRead(rawResponse.Result, count); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse count response result: %w", err)
			}
			countableList.SetCount(count.Count)
		} else if rawResponse.Result != nil {
			// ActionRead
			if err := UnmarshalRead(rawResponse.Result, spec); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result: %w", err)
			}
		} else if rawResponse.Results != nil {
			// ActionList
			listSpec, ok := spec.(ListSpec)
			if !ok {
				return rawResponse.RequestId, fmt.Errorf("not support ListSpec %s", spec.GetName())
			}
			items := listSpec.GetItems()
			if err := UnmarshalRead(rawResponse.Results, items); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse list response results: %w", err)
			}
		} else {
			// Other(Job)
			if err := UnmarshalRead(bs, spec); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result: %w", err)
			}
		}
	}

	// initialize process
	if d, ok := spec.(Initializer); ok {
		d.Init()
	}

	return rawResponse.RequestId, nil
}

func (c *Client) Read(ctx context.Context, s Spec) (requestId string, err error) {
	return c.Do(ctx, s, ActionRead, nil, nil)
}

func (c *Client) List(ctx context.Context, s ListSpec, keywords SearchParams) (requestId string, err error) {
	return c.Do(ctx, s, ActionList, nil, keywords)
}

func (c *Client) ListAll(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	req, err := c.Count(ctx, s, keywords)
	if err != nil {
		return req, err
	}

	if keywords == nil {
		keywords = &CommonSearchParams{}
		keywords.SetLimit(s.GetMaxLimit())
	}

	count := s.GetCount()
	cpObj := s.DeepCopyObject()

	for offset := int32(0); offset < count; offset += keywords.GetLimit() {
		list := cpObj.DeepCopyObject()
		cList := list.(ListSpec)
		keywords.SetOffset(offset)
		req, err = c.List(ctx, cList, keywords)
		if err != nil {
			return req, err
		}
		for i := 0; i < cList.Len(); i++ {
			s.AddItem(cList.Index(i))
		}
	}
	return req, nil
}

func (c *Client) Count(ctx context.Context, s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	return c.Do(ctx, s, ActionCount, nil, keywords)
}

func (c *Client) Update(ctx context.Context, s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionUpdate, body, nil)
}

func (c *Client) Create(ctx context.Context, s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionCreate, body, nil)
}

func (c *Client) Apply(ctx context.Context, s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(ctx, s, ActionApply, body, nil)
}

func (c *Client) Delete(ctx context.Context, s Spec) (requestId string, err error) {
	return c.Do(ctx, s, ActionDelete, nil, nil)
}

func (c *Client) Cancel(ctx context.Context, s Spec) (requestId string, err error) {
	return c.Do(ctx, s, ActionCancel, nil, nil)
}

func (c *Client) watch(ctx context.Context, interval time.Duration, f func() (keep bool, err error)) error {
	if interval < time.Second {
		return fmt.Errorf("interval must greater than equals to 1s")
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
LOOP:
	for {
		select {
		case <-ticker.C:
			loopBreak, err := f()
			if err != nil {
				return err
			}
			if loopBreak {
				break LOOP
			}
		case <-ctx.Done():
			break LOOP
		}
	}
	return ctx.Err()
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is Readable Spec
func (c *Client) WatchRead(ctx context.Context, interval time.Duration, s Spec) error {
	obj := s.DeepCopyObject()
	org := obj.(Spec)
	return c.watch(ctx, interval, func() (bool, error) {
		_, err := c.Read(ctx, s)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, org) {
			return false, nil
		}
		return true, nil
	})
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is ListAble Spec
func (c *Client) WatchList(ctx context.Context, interval time.Duration, s ListSpec, keyword SearchParams) error {
	obj := s.DeepCopyObject()
	org := obj.(ListSpec)
	return c.watch(ctx, interval, func() (bool, error) {
		_, err := c.List(ctx, s, keyword)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, org) {
			return false, nil
		}
		return true, nil
	})
}

// ctx should set Deadline or Timeout
// interval must be grater than equals to 1s
// s is CountableListSpec Spec
func (c *Client) WatchListAll(ctx context.Context, interval time.Duration, s CountableListSpec, keyword SearchParams) error {
	obj := s.DeepCopyObject()
	copy := obj.(CountableListSpec)
	copy.ClearItems()
	err := c.watch(ctx, interval, func() (bool, error) {
		_, err := c.ListAll(ctx, copy, keyword)
		if err != nil {
			return true, err
		}
		if reflect.DeepEqual(s, copy) {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	s.ClearItems()
	for i := 0; i < copy.Len(); i++ {
		s.AddItem(copy.Index(i))
	}
	s.SetCount(int32(copy.Len()))
	return nil

}
