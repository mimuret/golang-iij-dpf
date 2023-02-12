package lb_domains

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type RuleRRType string

const (
	RuleRRTypeA     RuleRRType = "A"
	RuleRRTypeAAAA  RuleRRType = "AAAA"
	RuleRRTypeCNAME RuleRRType = "CNAME"
)

var _ ChildSpec = &Rule{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Rule struct {
	AttributeMeta
	ResourceName string       `read:"resource_name" create:"resource_name,omitempty" apply:"resource_name"`
	Name         string       `read:"name" create:"name" update:"name" apply:"name"`
	Description  string       `read:"description" create:"description" update:"description" apply:"description"`
	Methods      []RuleMethod `read:"methods" apply:"methods"`
}

func (c *Rule) Init() {
	for i := range c.Methods {
		c.Methods[i].AttributeMeta = c.AttributeMeta
		c.Methods[i].RuleResourceName = c.ResourceName
	}
}

func (c *Rule) GetName() string                     { return "rules" }
func (c *Rule) GetResourceName() string             { return c.ResourceName }
func (c *Rule) SetResourceName(resourceName string) { c.ResourceName = resourceName }
func (c *Rule) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

func (c *Rule) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.ResourceName)
}

var _ ListSpec = &RuleList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type RuleList struct {
	AttributeMeta
	api.Count
	Items []Rule `read:"items"`
}

func (c *RuleList) GetName() string         { return "rules" }
func (c *RuleList) GetItems() interface{}   { return &c.Items }
func (c *RuleList) Len() int                { return len(c.Items) }
func (c *RuleList) Index(i int) interface{} { return c.Items[i] }
func (c *RuleList) GetMaxLimit() int32      { return 10000 }
func (c *RuleList) ClearItems()             { c.Items = []Rule{} }
func (c *RuleList) AddItem(v interface{}) bool {
	if a, ok := v.(Rule); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *RuleList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *RuleList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
		c.Items[i].Init()
	}
}

func (c *RuleList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID)
}

func init() {
	register(&Rule{}, &RuleList{})
}
