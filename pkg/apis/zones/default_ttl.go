package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type DefaultTTLState int

const (
	DefaultTTLStateApplied      DefaultTTLState = 0
	DefaultTTLStateToBeUpdate   DefaultTTLState = 3
	DefaultTTLStateBeforeUpdate DefaultTTLState = 5
)

var DefaultTTLStateToString = map[DefaultTTLState]string{
	DefaultTTLStateApplied:      "Applied",
	DefaultTTLStateToBeUpdate:   "ToBeUpdate",
	DefaultTTLStateBeforeUpdate: "BeforeUpdate",
}

func (c DefaultTTLState) String() string {
	return DefaultTTLStateToString[c]
}

var _ Spec = &DefaultTTL{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type DefaultTTL struct {
	AttributeMeta
	Value    int64           `read:"value" update:"value"`
	State    DefaultTTLState `read:"state"`
	Operator string          `read:"operator"`
}

func (c *DefaultTTL) GetName() string { return "default_ttl" }
func (c *DefaultTTL) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/default_ttl", c.GetZoneId())
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/default_ttl/changes", c.GetZoneId())
	}
	return "", ""
}
func (c *DefaultTTL) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

type DefaultTTLDiff struct {
	New *DefaultTTL `read:"new"`
	Old *DefaultTTL `read:"old"`
}

var _ ListSpec = &DefaultTTLDiffList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type DefaultTTLDiffList struct {
	AttributeMeta
	Items []DefaultTTLDiff `read:"items"`
}

func (c *DefaultTTLDiffList) GetName() string         { return "default_ttl/diffs" }
func (c *DefaultTTLDiffList) GetItems() interface{}   { return &c.Items }
func (c *DefaultTTLDiffList) Len() int                { return len(c.Items) }
func (c *DefaultTTLDiffList) Index(i int) interface{} { return c.Items[i] }

func (c *DefaultTTLDiffList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *DefaultTTLDiffList) Init() {
	for i := range c.Items {
		if c.Items[i].New != nil {
			c.Items[i].New.AttributeMeta = c.AttributeMeta
		}
		if c.Items[i].Old != nil {
			c.Items[i].Old.AttributeMeta = c.AttributeMeta
		}
	}
}
func (c *DefaultTTLDiffList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

func init() {
	Register.Add(&DefaultTTL{})
	Register.Add(&DefaultTTLDiffList{})
}
