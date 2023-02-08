package lb_domains

import (
	"encoding/json"
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ ChildSpec = &Monitoring{}

type MonitoringMtype string

const (
	MonitoringMtypePing   MonitoringMtype = "ping"
	MonitoringMtypeTCP    MonitoringMtype = "tcp"
	MonitoringMtypeHTTP   MonitoringMtype = "http"
	MonitoringMtypeStatic MonitoringMtype = "static"
)

type MonitoringPorps interface {
	GetMtype() MonitoringMtype
	DeepCopyMonitoringPorps() MonitoringPorps
}

type MonitoringCommon struct {
	ResourceName string          `read:"resource_name" create:"resource_name" apply:"resource_name"`
	Name         string          `read:"name" create:"name" update:"name" apply:"name"`
	MType        MonitoringMtype `read:"mtype" create:"mtype" apply:"mtype"`
	Description  string          `read:"description" create:"description" update:"description" apply:"description"`
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Monitoring struct {
	AttributeMeta
	MonitoringCommon
	Props MonitoringPorps `create:"props" update:"props" apply:"props"`
}

func (m *Monitoring) Fix() {
	if m.Props != nil {
		m.MType = m.Props.GetMtype()
	}
}

func (m *Monitoring) SetProps(props MonitoringPorps) {
	m.MType = props.GetMtype()
	m.Props = props
}

func (m *Monitoring) UnmarshalJSON(bs []byte) error {
	c := struct {
		MonitoringCommon
		Props json.RawMessage `read:"props"`
	}{}
	if err := api.UnmarshalRead(bs, &c); err != nil {
		return fmt.Errorf("failed to parse Monitoring: %w", err)
	}
	var props MonitoringPorps
	switch c.MType {
	case MonitoringMtypePing:
		props = &MonitoringPorpsPING{}
	case MonitoringMtypeTCP:
		props = &MonitoringPorpsTCP{}
	case MonitoringMtypeHTTP:
		props = &MonitoringPorpsHTTP{}
	case MonitoringMtypeStatic:
		props = &MonitoringPorpsStatic{}
	default:
		return fmt.Errorf("unknown mtype `%s`", c.MType)
	}
	if err := api.UnmarshalRead(c.Props, props); err != nil {
		return fmt.Errorf("failed to parse props: %w", err)
	}
	m.MonitoringCommon = c.MonitoringCommon
	m.SetProps(props)
	return nil
}

func (c *Monitoring) GetName() string                     { return "monitorings" }
func (c *Monitoring) GetResourceName() string             { return c.ResourceName }
func (c *Monitoring) SetResourceName(resourceName string) { c.ResourceName = resourceName }
func (c *Monitoring) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

func (c *Monitoring) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.ResourceName)
}

var _ ListSpec = &MonitoringList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type MonitoringList struct {
	AttributeMeta
	Items []Monitoring `read:"items"`
}

func (c *MonitoringList) GetName() string         { return "monitorings" }
func (c *MonitoringList) GetItems() interface{}   { return &c.Items }
func (c *MonitoringList) Len() int                { return len(c.Items) }
func (c *MonitoringList) Index(i int) interface{} { return c.Items[i] }

func (c *MonitoringList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *MonitoringList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func (c *MonitoringList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID)
}

func init() {
	register(&Monitoring{}, &MonitoringList{})
}
