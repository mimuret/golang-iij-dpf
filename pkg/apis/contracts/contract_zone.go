package contracts

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

var _ CountableListSpec = &ContractZoneList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type ContractZoneList struct {
	AttributeMeta
	api.Count
	Items []core.Zone `read:"items"`
}

func (c *ContractZoneList) GetName() string         { return "zones" }
func (c *ContractZoneList) GetItems() interface{}   { return &c.Items }
func (c *ContractZoneList) Len() int                { return len(c.Items) }
func (c *ContractZoneList) Index(i int) interface{} { return c.Items[i] }
func (c *ContractZoneList) GetMaxLimit() int32      { return 10000 }
func (c *ContractZoneList) ClearItems()             { c.Items = []core.Zone{} }
func (c *ContractZoneList) AddItem(v interface{}) bool {
	if a, ok := v.(core.Zone); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *ContractZoneList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *ContractZoneList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId)
}

func (c *ContractZoneList) Init() {}

var _ Spec = &ContractZoneCommonConfig{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type ContractZoneCommonConfig struct {
	AttributeMeta
	CommonConfigId int64    `update:"common_config_id"`
	ZoneIds        []string `update:"zone_ids"`
}

func (c *ContractZoneCommonConfig) GetName() string { return "zones/common_configs" }
func (c *ContractZoneCommonConfig) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/zones/common_configs", c.GetContractId())
	}
	return "", ""
}
func (c *ContractZoneCommonConfig) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId)
}

func init() {
	Register.Add(&ContractZoneList{})
	Register.Add(&ContractZoneCommonConfig{})
}
