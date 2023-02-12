package lb_domains

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &Config{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Config struct {
	AttributeMeta
	Monitorings []Monitoring `read:"monitorings" apply:"monitorings"`
	Sites       []Site       `read:"sites" apply:"sites"`
	Rules       []Rule       `read:"rules" apply:"rules"`
}

func (c *Config) Init() {
	for i := range c.Monitorings {
		c.Monitorings[i].AttributeMeta = c.AttributeMeta
	}
	for i := range c.Sites {
		c.Sites[i].AttributeMeta = c.AttributeMeta
		c.Sites[i].Init()
	}
	for i := range c.Rules {
		c.Rules[i].AttributeMeta = c.AttributeMeta
		c.Rules[i].Init()
	}
}

func (c *Config) GetName() string { return "sites" }
func (c *Config) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return http.MethodPut, fmt.Sprintf("/lb_domains/%s/config", c.GetLBDoaminID())
	case api.ActionRead:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/config", c.GetLBDoaminID())
	}
	return "", ""
}

func (c *Config) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID)
}

func init() {
	register(&Config{})
}
