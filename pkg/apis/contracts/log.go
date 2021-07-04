package contracts

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type LogList struct {
	AttributeMeta
	api.Count
	Items []core.Log `read:"items"`
}

var _ CountableListSpec = &LogList{}

func (c *LogList) GetName() string         { return "logs" }
func (c *LogList) GetItems() interface{}   { return &c.Items }
func (c *LogList) Len() int                { return len(c.Items) }
func (c *LogList) Index(i int) interface{} { return c.Items[i] }
func (c *LogList) GetMaxLimit() int32      { return 100 }
func (c *LogList) ClearItems()             { c.Items = []core.Log{} }
func (c *LogList) AddItem(v interface{}) bool {
	if a, ok := v.(core.Log); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *LogList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *LogList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ContractId)
}

func (c *LogList) Init() {}

func init() {
	Register.Add(&LogList{})
}
