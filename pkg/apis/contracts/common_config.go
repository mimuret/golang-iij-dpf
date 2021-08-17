package contracts

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ ChildSpec = &CommonConfig{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfig struct {
	AttributeMeta
	Id int64 `read:"id"`
	// patchable
	Name        string `read:"name"  create:"name" update:"name"`
	Description string `read:"description"  create:"description" update:"description"`
	// not patchable
	ManagedDNSEnabled types.Boolean `read:"managed_dns_enabled"`
	Default           types.Boolean `read:"default"`
}

func (c *CommonConfig) GetName() string { return "common_configs" }
func (c *CommonConfig) GetId() int64    { return c.Id }
func (c *CommonConfig) SetId(id int64)  { c.Id = id }
func (c *CommonConfig) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}
func (c *CommonConfig) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractId, &c.Id)
}

var _ CountableListSpec = &CommonConfigList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfigList struct {
	AttributeMeta
	api.Count
	Items []CommonConfig `read:"items"`
}

func (c *CommonConfigList) GetName() string         { return "common_configs" }
func (c *CommonConfigList) GetItems() interface{}   { return &c.Items }
func (c *CommonConfigList) Len() int                { return len(c.Items) }
func (c *CommonConfigList) Index(i int) interface{} { return c.Items[i] }
func (c *CommonConfigList) GetMaxLimit() int32      { return 10000 }
func (c *CommonConfigList) ClearItems()             { c.Items = []CommonConfig{} }
func (c *CommonConfigList) AddItem(v interface{}) bool {
	if a, ok := v.(CommonConfig); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *CommonConfigList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *CommonConfigList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}
func (c *CommonConfigList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractId)
}

var _ api.SearchParams = &CommonConfigListSearchKeywords{}

// +k8s:deepcopy-gen=false
type CommonConfigListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	Name        api.KeywordsString `url:"_keywords_name[],omitempty"`
	Description api.KeywordsString `url:"_keywords_description[],omitempty"`
}

func (s *CommonConfigListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	Register.Add(&CommonConfigList{})
}
