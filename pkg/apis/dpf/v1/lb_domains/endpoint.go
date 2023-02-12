package lb_domains

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ ChildSpec = &Site{}

type SiteAttributeMeta struct {
	AttributeMeta
	SiteResourceName string
}

type MonitoringEndpoint struct {
	MonitoringResourceName string `update:"resource_name" create:"resource_name" apply:"resource_name"`
	Enabled                bool   `create:"enabled" update:"enabled" apply:"enabled"`
	Monitoring             *Monitoring
}

func (m *MonitoringEndpoint) UnmarshalJSON(bs []byte) error {
	enabled := struct {
		Enabled bool `read:"enabled"`
	}{}
	var monitoring Monitoring
	if err := api.UnmarshalRead(bs, &enabled); err != nil {
		return fmt.Errorf("failed to parse MonitoringEndpoint")
	}
	if err := api.UnmarshalRead(bs, &monitoring); err != nil {
		return fmt.Errorf("failed to parse MonitoringEndpoint")
	}
	m.MonitoringResourceName = monitoring.ResourceName
	m.Enabled = enabled.Enabled
	m.Monitoring = &monitoring

	return nil
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Endpoint struct {
	SiteAttributeMeta
	ResourceName     string               `read:"resource_name" create:"resource_name,omitempty" apply:"resource_name"`
	Name             string               `read:"name" create:"name" update:"name" apply:"name"`
	MonitoringTarget string               `read:"monitoring_target" create:"monitoring_target" update:"monitoring_target" apply:"monitoring_target"`
	Description      string               `read:"description" create:"description" update:"description" apply:"description"`
	Weight           uint8                `read:"weight" create:"weight" update:"weight" apply:"weight"`
	ManualFaillback  bool                 `read:"manual_failback" create:"manual_failback" update:"manual_failback" apply:"manual_failback"`
	ManualFaillOver  bool                 `read:"manual_failover" create:"manual_failover" update:"manual_failover" apply:"manual_failover"`
	Enabled          bool                 `read:"enabled" create:"enabled" update:"enabled" apply:"enabled"`
	LiveStatus       Status               `read:"live_status"`
	ReadyStatus      Status               `read:"ready_status"`
	Rdata            []string             `read:"rdata" create:"rdata" update:"rdata" apply:"rdata"`
	Monitorings      []MonitoringEndpoint `read:"monitorings" create:"monitorings" update:"monitorings" apply:"monitorings"`
}

func (c *Endpoint) Fix() {
	if c.Weight == 0 {
		c.Weight = uint8(1)
	}
}

func (c *Endpoint) Init() {
	for i := range c.Monitorings {
		if c.Monitorings[i].Monitoring != nil {
			c.Monitorings[i].Monitoring.AttributeMeta = c.AttributeMeta
		}
	}
}

func (c *Endpoint) GetName() string                     { return "endpoints" }
func (c *Endpoint) GetResourceName() string             { return c.ResourceName }
func (c *Endpoint) SetResourceName(resourceName string) { c.ResourceName = resourceName }
func (c *Endpoint) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/sites/%s/endpoints", c.GetLBDoaminID(), c.SiteResourceName)
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/sites/%s/endpoints/%s", c.GetLBDoaminID(), c.SiteResourceName, c.ResourceName)
	}
	return "", ""
}

func (c *Endpoint) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.SiteResourceName, &c.ResourceName)
}

var _ ListSpec = &EndpointList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type EndpointList struct {
	SiteAttributeMeta
	Items []Endpoint `read:"items"`
}

func (c *EndpointList) GetName() string         { return "endpoints" }
func (c *EndpointList) GetItems() interface{}   { return &c.Items }
func (c *EndpointList) Len() int                { return len(c.Items) }
func (c *EndpointList) Index(i int) interface{} { return c.Items[i] }

func (c *EndpointList) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionList {
		return http.MethodGet, fmt.Sprintf("/lb_domains/%s/sites/%s/endpoints", c.GetLBDoaminID(), c.SiteResourceName)
	}
	return "", ""
}

func (c *EndpointList) Init() {
	for i := range c.Items {
		c.Items[i].SiteAttributeMeta = c.SiteAttributeMeta
		c.Items[i].Init()
	}
}

func (c *EndpointList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.SiteResourceName)
}

func init() {
	register(&Endpoint{}, &EndpointList{})
}
