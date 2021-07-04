package zones

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ Spec = &HistoryText{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type HistoryText struct {
	AttributeMeta
	History
	Text string `read:"text"`
}

func (c *HistoryText) GetName() string { return "zone_histories" }
func (c *HistoryText) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionRead {
		return action.ToMethod(), fmt.Sprintf("/zones/%s/zone_histories/%d/text", c.GetZoneId(), c.Id)
	}
	return "", ""
}
func (c *HistoryText) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId, &c.Id)
}

type History struct {
	Id          int64      `read:"id"`
	CommittedAt types.Time `read:"committed_at"`
	Description string     `read:"description"`
	Operator    string     `read:"operator"`
}

var _ CountableListSpec = &HistoryList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type HistoryList struct {
	AttributeMeta
	api.Count
	Items []History `read:"items"`
}

func (c *HistoryList) GetName() string         { return "zone_histories" }
func (c *HistoryList) GetItems() interface{}   { return &c.Items }
func (c *HistoryList) Len() int                { return len(c.Items) }
func (c *HistoryList) Index(i int) interface{} { return c.Items[i] }
func (c *HistoryList) GetMaxLimit() int32      { return 100 }
func (c *HistoryList) ClearItems()             { c.Items = []History{} }
func (c *HistoryList) AddItem(v interface{}) bool {
	if a, ok := v.(History); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *HistoryList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *HistoryList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

func (c *HistoryList) Init() {}

var _ api.SearchParams = &HistoryListSearchKeywords{}

// +k8s:deepcopy-gen=false

type HistoryListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	Description api.KeywordsString `url:"_keywords_description[],omitempty"`
	Operator    api.KeywordsString `url:"_keywords_operator[],omitempty"`
}

func (s *HistoryListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	Register.Add(&HistoryList{})
}
