package lb_domains

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type SiteRRType string

const (
	SiteRRTypeA     SiteRRType = "A"
	SiteRRTypeAAAA  SiteRRType = "AAAA"
	SiteRRTypeCNAME SiteRRType = "CNAME"
)

var _ ChildSpec = &Site{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Site struct {
	AttributeMeta
	ResourceName string     `read:"resource_name" create:"resource_name,omitempty" apply:"resource_name"`
	Name         string     `read:"name" create:"name" update:"name" apply:"name"`
	RRType       SiteRRType `read:"rrtype" create:"rrtype" apply:"rrtype"`
	Description  string     `read:"description" create:"description" update:"description" apply:"description"`
	LiveStatus   Status     `read:"live_status"`
	Endpoints    []Endpoint `read:"endpoints" apply:"endpoints"`
}

func (c *Site) Init() {
	for i := range c.Endpoints {
		c.Endpoints[i].AttributeMeta = c.AttributeMeta
		c.Endpoints[i].SiteResourceName = c.ResourceName
		c.Endpoints[i].Init()
	}
}

func (c *Site) GetName() string                     { return "sites" }
func (c *Site) GetResourceName() string             { return c.ResourceName }
func (c *Site) SetResourceName(resourceName string) { c.ResourceName = resourceName }
func (c *Site) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

func (c *Site) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.ResourceName)
}

var _ ListSpec = &SiteList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type SiteList struct {
	AttributeMeta
	api.Count
	Items []Site `read:"items"`
}

func (c *SiteList) GetName() string         { return "sites" }
func (c *SiteList) GetItems() interface{}   { return &c.Items }
func (c *SiteList) Len() int                { return len(c.Items) }
func (c *SiteList) Index(i int) interface{} { return c.Items[i] }

func (c *SiteList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *SiteList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
		c.Items[i].Init()
	}
}

func (c *SiteList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID)
}

func init() {
	register(&Site{}, &SiteList{})
}
