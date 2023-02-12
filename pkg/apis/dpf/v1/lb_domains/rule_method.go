package lb_domains

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &RuleMethod{}

type RuleAttributeMeta struct {
	AttributeMeta
	RuleResourceName string
}

type RuleMethodProps interface {
	Fix()
	GetMType() RuleMethodMType
	GetMethodResourceName() string
	SetMethodResourceName(string)
	DeepCopyRuleMethodProps() RuleMethodProps
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type RuleMethod struct {
	RuleAttributeMeta
	Priority *uint           `create:"priority,omitempty" update:"priority,omitempty" apply:"priority,omitempty"`
	Method   RuleMethodProps `create:"method" update:"method" apply:"method"`
}

func (c *RuleMethod) Fix() {
	if c.Method != nil {
		c.Method.Fix()
	}
}

func (c *RuleMethod) UnmarshalJSON(bs []byte) error {
	r := struct {
		Priority *uint           `read:"priority"`
		Method   json.RawMessage `read:"method"`
	}{}
	if err := api.UnmarshalRead(bs, &r); err != nil {
		return fmt.Errorf("failed to parse RuleMethod: %w", err)
	}
	propsCommon := &RuleMethodPropsCommon{}
	if err := api.UnmarshalRead(r.Method, propsCommon); err != nil {
		return fmt.Errorf("failed to parse Method: %w", err)
	}

	var props RuleMethodProps
	switch propsCommon.MType {
	case "entry_a":
		props = &RuleMethodEntryA{}
	case "entry_aaaa":
		props = &RuleMethodEntryAAAA{}
	case "entry_cname":
		props = &RuleMethodEntryCNAME{}
	case "exit_site":
		props = &RuleMethodExitSite{}
	case "exit_sorry":
		props = &RuleMethodExitSorry{}
	case "failover":
		props = &RuleMethodFailover{}
	default:
		return fmt.Errorf("unknown mtype `%s`", propsCommon.MType)
	}
	if err := api.UnmarshalRead(r.Method, props); err != nil {
		return fmt.Errorf("failed to parse props: %w", err)
	}
	c.Priority = r.Priority
	c.Method = props
	return nil
}

func (c *RuleMethod) GetName() string               { return "rule_methods" }
func (c *RuleMethod) GetMethodResourceName() string { return c.Method.GetMethodResourceName() }
func (c *RuleMethod) SetMethodResourceName(resourceName string) {
	c.Method.SetMethodResourceName(resourceName)
}

func (c *RuleMethod) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/rules/%s/rule_methods", c.GetLBDoaminID(), c.RuleResourceName)
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/rules/%s/rule_methods/%s", c.GetLBDoaminID(), c.RuleResourceName, c.Method.GetMethodResourceName())
	}
	return "", ""
}

func (c *RuleMethod) SetPathParams(args ...interface{}) error {
	var methodResourceName string
	if err := apis.SetPathParams(args, &c.LBDomainID, &c.RuleResourceName, &methodResourceName); err != nil {
		return err
	}
	c.SetMethodResourceName(methodResourceName)
	return nil
}

var _ ListSpec = &MonitoringList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type RuleMethodList struct {
	RuleAttributeMeta
	Items []RuleMethod `read:"items"`
}

func (c *RuleMethodList) GetName() string         { return "rule_methods" }
func (c *RuleMethodList) GetItems() interface{}   { return &c.Items }
func (c *RuleMethodList) Len() int                { return len(c.Items) }
func (c *RuleMethodList) Index(i int) interface{} { return c.Items[i] }

func (c *RuleMethodList) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionList {
		return http.MethodGet, fmt.Sprintf("/lb_domains/%s/rules/%s/rule_methods", c.GetLBDoaminID(), c.RuleResourceName)
	}
	return "", ""
}

func (c *RuleMethodList) Init() {
	for i := range c.Items {
		c.Items[i].RuleAttributeMeta = c.RuleAttributeMeta
	}
}

func (c *RuleMethodList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.RuleResourceName)
}

func init() {
	register(&RuleMethod{}, &RuleMethodList{})
}
