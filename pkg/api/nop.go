package api

import (
	"context"
	"time"
)

var _ ClientInterface = &Client{}

type NopClient struct {
}

func (n *NopClient) SetWatchInterval(d time.Duration) error { return nil }
func (n *NopClient) SetWatchTimeout(d time.Duration) error  { return nil }
func (n *NopClient) Do(spec Spec, action Action, body interface{}, params SearchParams) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Read(s Spec) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) List(s ListSpec, keywords SearchParams) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) ListALL(s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Count(s CountableListSpec, keywords SearchParams) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Update(s Spec, body interface{}) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Create(s Spec, body interface{}) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Apply(s Spec, body interface{}) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Delete(s Spec) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Cancel(s Spec) (requestId string, err error) {
	return "B1BBA7357448495288DD914D559A0C94", nil
}
func (n *NopClient) Watch(ctx context.Context, f func() (keep bool, err error)) error {
	return nil
}
func (n *NopClient) WatchRead(ctx context.Context, s Spec) error {
	return nil
}
func (n *NopClient) WatchList(ctx context.Context, s ListSpec, keyword SearchParams) error {
	return nil
}
func (n *NopClient) WatchListAll(ctx context.Context, s CountableListSpec, keyword SearchParams) error {
	return nil
}
