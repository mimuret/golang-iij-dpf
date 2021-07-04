package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const DefaultEndpoint = "https://api.dns-platform.jp/dpf/v1"

type Client struct {
	Endpoint string
	Token    string
	logger   Logger

	watchInterval time.Duration
	watchTimeout  time.Duration

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
		logger = &StdLogger{LogLevel: 4}
	}
	return &Client{Endpoint: endpoint, Token: token, logger: logger, watchInterval: time.Second * 5, watchTimeout: time.Minute * 5}
}

func (c *Client) SetWatchInterval(d time.Duration) error {
	if d < time.Second*5 {
		return fmt.Errorf("WatchInterval must greater than equals to 5s")
	}
	if c.watchTimeout <= d {
		return fmt.Errorf("WatchInterval must less than WatchTimeout")
	}
	c.watchInterval = d
	return nil
}

func (c *Client) SetWatchTimeout(d time.Duration) error {
	if c.watchInterval >= d {
		return fmt.Errorf("WatchTimeout must greater than WatchInterval")
	}
	c.watchTimeout = d
	return nil
}

func (c *Client) SetLogger(l Logger) *Client {
	c.logger = l
	return c
}

func (c *Client) Do(spec Spec, action Action, body interface{}, params SearchParams) (requestId string, err error) {
	var r io.Reader
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
	client := &http.Client{}

	// request
	resp, err := client.Do(req)
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
		badRequest := &BadResponse{}
		if err := UnmarshalRead(bs, badRequest); err != nil {
			return "", fmt.Errorf("failed to request: rcode: %d body: %s err: %w", resp.StatusCode, string(bs), err)
		}
		return badRequest.RequestId, badRequest
	}

	// parse raw response
	rawResponse := &RawResponse{}
	if err := UnmarshalRead(bs, rawResponse); err != nil {
		return "", fmt.Errorf("failed to parse get response: %w", err)
	}

	if method == http.MethodGet {
		if action == ActionCount {
			counter, ok := spec.(CountableListSpec)
			if !ok {
				return rawResponse.RequestId, fmt.Errorf("failed to get Count interface")
			}
			countResp := &CountResponse{}
			if err := UnmarshalRead(rawResponse.Result, countResp); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result: %w", err)
			}
			counter.SetCount(countResp.Result.Count)
		}
		if rawResponse.Result != nil {
			if err := UnmarshalRead(rawResponse.Result, spec); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result: %w", err)
			}
		} else if rawResponse.Results != nil {
			listSpec, ok := spec.(ListSpec)
			if !ok {
				return rawResponse.RequestId, fmt.Errorf("not support ListSpec %s", spec.GetName())
			}
			items := listSpec.GetItems()
			if err := UnmarshalRead(rawResponse.Results, items); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result list: %w", err)
			}
		} else {
			if err := UnmarshalRead(bs, spec); err != nil {
				return rawResponse.RequestId, fmt.Errorf("failed to parse response result list: %w", err)
			}
		}
	}

	// initialize process
	if d, ok := spec.(Initializer); ok {
		d.Init()
	}

	return rawResponse.RequestId, nil
}

func (c *Client) Read(s Spec) (requestId string, err error) {
	return c.Do(s, ActionRead, nil, nil)
}

func (c *Client) List(s ListSpec, keywords SearchParams) (requestId string, err error) {
	return c.Do(s, ActionList, nil, keywords)
}

func (c *Client) ListALL(s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	req, err := c.Count(s, keywords)
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
		cList, _ := list.(ListSpec)
		keywords.SetOffset(offset)
		req, err = c.List(cList, keywords)
		if err != nil {
			return req, err
		}
		for i := 0; i < cList.Len(); i++ {
			s.AddItem(cList.Index(i))
		}
	}
	return req, nil
}

func (c *Client) Count(s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	return c.Do(s, ActionCount, nil, keywords)
}

func (c *Client) Update(s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(s, ActionUpdate, body, nil)
}

func (c *Client) Create(s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(s, ActionCreate, body, nil)
}

func (c *Client) Apply(s Spec, body interface{}) (requestId string, err error) {
	if body == nil {
		body = s
	}
	return c.Do(s, ActionApply, body, nil)
}

func (c *Client) Delete(s Spec) (requestId string, err error) {
	return c.Do(s, ActionDelete, nil, nil)
}

func (c *Client) Cancel(s Spec) (requestId string, err error) {
	return c.Do(s, ActionCancel, nil, nil)
}

func (c *Client) Watch(ctx context.Context, f func() (keep bool, err error)) error {
	timer := time.NewTimer(c.watchTimeout)
	defer timer.Stop()
	ticker := time.NewTicker(c.watchInterval)
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
		case <-timer.C:
			return fmt.Errorf("timeout")
		}
	}
	return nil

}
func (c *Client) WatchRead(ctx context.Context, s Spec) error {
	obj := s.DeepCopyObject()
	org, ok := obj.(Spec)
	if !ok {
		return fmt.Errorf("is not api.Spec")
	}
	err := c.Watch(ctx, func() (bool, error) {
		_, err := c.Read(s)
		if err != nil {
			return true, err
		}
		if s == org {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) WatchList(ctx context.Context, s ListSpec, keyword SearchParams) error {
	obj := s.DeepCopyObject()
	org, ok := obj.(ListSpec)
	if !ok {
		return fmt.Errorf("is not api.List")
	}
	err := c.Watch(ctx, func() (bool, error) {
		_, err := c.List(s, keyword)
		if err != nil {
			return true, err
		}
		if s == org {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) WatchListAll(ctx context.Context, s CountableListSpec, keyword SearchParams) error {
	obj := s.DeepCopyObject()
	copy, ok := obj.(CountableListSpec)
	if !ok {
		return fmt.Errorf("is not api.List")
	}
	copy.ClearItems()
	err := c.Watch(ctx, func() (bool, error) {
		_, err := c.ListALL(copy, keyword)
		if err != nil {
			return true, err
		}
		if s == copy {
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
	return nil

}
