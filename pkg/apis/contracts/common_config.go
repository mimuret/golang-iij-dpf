package contracts

import (
	"fmt"
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
func (c *CommonConfig) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId, &c.Id)
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
func (c *CommonConfigList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId)
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

var _ Spec = &CommonConfigDefault{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfigDefault struct {
	AttributeMeta  `update:"-"`
	CommonConfigId int64 `update:"common_config_id"`
}

func (c *CommonConfigDefault) GetName() string { return "common_configs" }
func (c *CommonConfigDefault) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/common_configs/default", c.GetContractId())
	}
	return "", ""
}
func (c *CommonConfigDefault) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId)
}

var _ Spec = &CommonConfigManagedDns{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfigManagedDns struct {
	AttributeMeta
	Id                int64
	ManagedDnsEnabled types.Boolean `update:"managed_dns_enabled"`
}

func (c *CommonConfigManagedDns) GetName() string { return "common_configs" }
func (c *CommonConfigManagedDns) GetId() int64    { return c.Id }
func (c *CommonConfigManagedDns) SetId(id int64)  { c.Id = id }
func (c *CommonConfigManagedDns) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/common_configs/%d/managed_dns", c.GetContractId(), c.Id)
	}
	return "", ""
}
func (c *CommonConfigManagedDns) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId, &c.Id)
}

func init() {
	Register.Add(&CommonConfig{}, &CommonConfigList{})
	Register.Add(&CommonConfigDefault{})
	Register.Add(&CommonConfigManagedDns{})
}
