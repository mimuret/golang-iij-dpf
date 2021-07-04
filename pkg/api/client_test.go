package api_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/stretchr/testify/assert"
)

var _ api.Spec = &TestSpec{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type TestSpec struct {
	Id     string `read:"id"`
	Name   string `read:"name"`
	Number int64  `read:"number"`
}

func (t *TestSpec) GetGroup() string { return "test" }
func (t *TestSpec) GetName() string  { return "tests" }
func (t *TestSpec) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), "/tests"
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/tests/%s", t.Id)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/tests/%s/cancel", t.Id)
	}
	return "", ""
}
func (t *TestSpec) DeepCopyTestSpec() *TestSpec {
	res := &TestSpec{}
	*res = *t
	return res
}

func (t *TestSpec) DeepCopyObject() api.Object {
	return t.DeepCopyTestSpec()
}

var _ api.CountableListSpec = &TestSpecList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type TestSpecList struct {
	api.Count
	Items []TestSpec `read:"items"`
}

func (t *TestSpecList) DeepCopyObject() api.Object {
	res := &TestSpecList{}
	res.Count = t.Count
	for _, item := range t.Items {
		s := item.DeepCopyTestSpec()
		res.Items = append(res.Items, *s)
	}
	return res
}
func (t *TestSpecList) GetGroup() string        { return "test" }
func (t *TestSpecList) GetName() string         { return "tests" }
func (t *TestSpecList) GetItems() interface{}   { return &t.Items }
func (c *TestSpecList) Len() int                { return len(c.Items) }
func (c *TestSpecList) Index(i int) interface{} { return c.Items[i] }
func (c *TestSpecList) GetMaxLimit() int32      { return 10000 }
func (c *TestSpecList) ClearItems()             { c.Items = []TestSpec{} }
func (c *TestSpecList) AddItem(v interface{}) bool {
	if a, ok := v.(TestSpec); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (t *TestSpecList) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionList:
		return action.ToMethod(), "/tests"
	case api.ActionCount:
		return action.ToMethod(), "/tests/count"
	}
	return "", ""
}
func (t *TestSpecList) Init() {}

func TestGet(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
		"result": {
			"id": "F2246BD8617444329A40470AEC7B00B9",
			"name": "test1",
			"number": 1
		}
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/C79CE3B1C87B47FA9BC618E6C40C3BD1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "CC84F36FC71749E4A50BD208DD20E535",
		"result": {
			"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
			"name": "test2",
			"number": 2
		}
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	testcase := []struct {
		RequestId string
		Id        string
		Name      string
	}{
		{"0CAC3AEFB6334ECCA7B70AF76D73508B", "F2246BD8617444329A40470AEC7B00B9", "test1"},
		{"CC84F36FC71749E4A50BD208DD20E535", "C79CE3B1C87B47FA9BC618E6C40C3BD1", "test2"},
	}
	for _, tc := range testcase {
		s := &TestSpec{Id: tc.Id}
		reqId, err := client.Read(s)
		if assert.NoError(t, err) {
			assert.Equal(t, reqId, tc.RequestId)
			assert.Equal(t, s.Id, tc.Id)
			assert.Equal(t, s.Name, tc.Name)
		}
	}

}
func TestGetList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
		"results": [
			{
				"id": "F2246BD8617444329A40470AEC7B00B9",
				"name": "test1",
				"number": 1
			},
			{
				"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
				"name": "test2",
				"number": 2
			}
		]
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	testcase := []TestSpec{
		{"F2246BD8617444329A40470AEC7B00B9", "test1", 1},
		{"C79CE3B1C87B47FA9BC618E6C40C3BD1", "test2", 2},
	}

	s := &TestSpecList{}
	_, err := client.List(s, nil)
	if assert.NoError(t, err) {
		if assert.Len(t, s.Items, 2) {
			for i, tc := range testcase {
				assert.Equal(t, s.Items[i], tc)
			}
		}
	}
}

func TestGetListAll(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "B7B871E6357E453287E2D7A9590D195D",
		"results": [
			{
				"id": "F2246BD8617444329A40470AEC7B00B9",
				"name": "test1",
				"number": 1
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "5F234DDF1D8E47D1BB26745D13FAE459",
		"results": [
			{
				"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
				"name": "test2",
				"number": 2
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
		"result": {
			"count": 2
		}
	}`)))

	client := api.NewClient("", "http://localhost", &api.StdLogger{LogLevel: 0})
	testcase := []TestSpec{
		{"F2246BD8617444329A40470AEC7B00B9", "test1", 1},
		{"C79CE3B1C87B47FA9BC618E6C40C3BD1", "test2", 2},
	}

	s := &TestSpecList{}
	reqId, err := client.ListALL(s, &api.CommonSearchParams{Limit: 1})
	if assert.NoError(t, err) {
		if assert.Len(t, s.Items, 2) {
			for i, tc := range testcase {
				assert.Equal(t, s.Items[i], tc)
			}
		}
	}
	assert.Equal(t, reqId, "5F234DDF1D8E47D1BB26745D13FAE459")
}
func TestGetCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
		"result": {
			"count": 2
		}
	}`)))
	client := api.NewClient("", "http://localhost", nil)

	s := &TestSpecList{}
	_, err := client.Count(s, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, s.GetCount(), int32(2))
	}
}

func TestPost(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPost, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "8C246D4DF69F4473AD9F90ADE0D4DAF2",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/8C246D4DF69F4473AD9F90ADE0D4DAF2"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	s := &TestSpec{"A99B06830D3B4D838EB7756B29CD302F", "test3", 1}
	reqId, err := client.Create(s, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "8C246D4DF69F4473AD9F90ADE0D4DAF2")
	}
}
func TestPatch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/F9BE7156EA4F4D83B69DBB067AD6858F", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0BC512B0B2FD45E9BDC3B60C683C901C",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/0BC512B0B2FD45E9BDC3B60C683C901C"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	s := &TestSpec{"F9BE7156EA4F4D83B69DBB067AD6858F", "test3", 2}
	reqId, err := client.Update(s, nil)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "0BC512B0B2FD45E9BDC3B60C683C901C")
	}

}
func TestDelete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodDelete, "/tests/F9BE7156EA4F4D83B69DBB067AD6858F", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "049B7CA403BC44309ECC6BE8D4739A8A",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/049B7CA403BC44309ECC6BE8D4739A8A"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	s := &TestSpec{"F9BE7156EA4F4D83B69DBB067AD6858F", "test3", 3}
	reqId, err := client.Delete(s)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "049B7CA403BC44309ECC6BE8D4739A8A")
	}

}
func TestCancel(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodDelete, "/tests/F9BE7156EA4F4D83B69DBB067AD6858F/cancel", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "049B7CA403BC44309ECC6BE8D4739A8A",
		"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/049B7CA403BC44309ECC6BE8D4739A8A"
	}`)))
	client := api.NewClient("", "http://localhost", nil)
	s := &TestSpec{"F9BE7156EA4F4D83B69DBB067AD6858F", "test3", 3}
	reqId, err := client.Delete(s)
	if assert.NoError(t, err) {
		assert.Equal(t, reqId, "049B7CA403BC44309ECC6BE8D4739A8A")
	}

}

func TestWatchRead(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
		"result": {
			"id": "F2246BD8617444329A40470AEC7B00B9",
			"name": "test1",
			"number": 1
		}
	}`)))

	client := api.NewClient("", "http://localhost", &api.StdLogger{LogLevel: 0})

	s := &TestSpec{Id: "F2246BD8617444329A40470AEC7B00B9"}
	_, err := client.Read(s)
	if assert.NoError(t, err) {
		go func() {
			time.Sleep(time.Second)
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"result": {
					"id": "F2246BD8617444329A40470AEC7B00B9",
					"name": "test2",
					"number": 1
				}
			}`)))
		}()
		assert.Equal(t, s.Name, "test1")
		err := client.WatchRead(context.Background(), s)
		if assert.NoError(t, err) {
			assert.Equal(t, s.Name, "test2")
		}
	}
}

func TestWatchList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
		"results": [
			{
				"id": "F2246BD8617444329A40470AEC7B00B9",
				"name": "test1",
				"number": 1
			}
		]
	}`)))

	client := api.NewClient("", "http://localhost", &api.StdLogger{LogLevel: 0})

	s := &TestSpecList{}
	_, err := client.List(s, nil)
	if assert.NoError(t, err) {
		go func() {
			time.Sleep(time.Second)
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"results": [
					{
						"id": "F2246BD8617444329A40470AEC7B00B9",
						"name": "test2",
						"number": 1
					}
				]
			}`)))
		}()
		assert.Equal(t, s.Items[0].Name, "test1")
		err := client.WatchList(context.Background(), s, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, s.Items[0].Name, "test2")
		}
	}
}

func TestWatchListAll(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "B7B871E6357E453287E2D7A9590D195D",
		"results": [
			{
				"id": "F2246BD8617444329A40470AEC7B00B9",
				"name": "test1",
				"number": 1
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "5F234DDF1D8E47D1BB26745D13FAE459",
		"results": [
			{
				"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
				"name": "test2",
				"number": 2
			}
		]
	}`)))
	httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
		"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
		"result": {
			"count": 2
		}
	}`)))

	client := api.NewClient("", "http://localhost", &api.StdLogger{LogLevel: 0})

	s := &TestSpecList{}
	_, err := client.ListALL(s, &api.CommonSearchParams{Limit: 1})
	if assert.NoError(t, err) {
		go func() {
			time.Sleep(time.Second)
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "B7B871E6357E453287E2D7A9590D195D",
				"results": [
					{
						"id": "F2246BD8617444329A40470AEC7B00B9",
						"name": "test3",
						"number": 1
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "5F234DDF1D8E47D1BB26745D13FAE459",
				"results": [
					{
						"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
						"name": "test4",
						"number": 2
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
				"result": {
					"count": 2
				}
			}`)))
		}()
		assert.Equal(t, s.Items[0].Name, "test1")
		assert.Equal(t, len(s.Items), 2)
		err := client.WatchListAll(context.Background(), s, &api.CommonSearchParams{Limit: 1})
		if assert.NoError(t, err) {
			assert.Equal(t, len(s.Items), 2)
			assert.Equal(t, s.Items[0].Name, "test3")
		}
	}
}
