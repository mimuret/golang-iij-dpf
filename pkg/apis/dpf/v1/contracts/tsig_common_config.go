package contracts

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ CountableListSpec = &TsigCommonConfigList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type TsigCommonConfigList struct {
	AttributeMeta
	api.Count
	ID    int64
	Items []CommonConfig `read:"items"`
}

func (c *TsigCommonConfigList) GetName() string {
	return fmt.Sprintf("tsigs/%d/common_configs", c.ID)
}
func (c *TsigCommonConfigList) GetID() int64   { return c.ID }
func (c *TsigCommonConfigList) SetID(id int64) { c.ID = id }

func (c *TsigCommonConfigList) GetItems() interface{}   { return &c.Items }
func (c *TsigCommonConfigList) Len() int                { return len(c.Items) }
func (c *TsigCommonConfigList) Index(i int) interface{} { return c.Items[i] }
func (c *TsigCommonConfigList) GetMaxLimit() int32      { return 10000 }
func (c *TsigCommonConfigList) ClearItems()             { c.Items = []CommonConfig{} }
func (c *TsigCommonConfigList) AddItem(v interface{}) bool {
	if a, ok := v.(CommonConfig); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *TsigCommonConfigList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *TsigCommonConfigList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func (c *TsigCommonConfigList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID, &c.ID)
}

func init() {
	register(&TsigCommonConfigList{})
}
