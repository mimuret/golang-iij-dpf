package lb_domains

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ ChildSpec = &Site{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type EndpointManualFailover struct {
	SiteAttributeMeta
	EndpointResourceName string
}

func (c *EndpointManualFailover) GetName() string { return "failover" }
func (c *EndpointManualFailover) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return http.MethodPost, fmt.Sprintf("/lb_domains/%s/sites/%s/endpoints/%s/failover", c.LBDomainID, c.SiteResourceName, c.EndpointResourceName)
	}
	return "", ""
}

func (c *EndpointManualFailover) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.SiteResourceName, &c.EndpointResourceName)
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type EndpointManualFailback struct {
	SiteAttributeMeta
	EndpointResourceName string
}

func (c *EndpointManualFailback) GetName() string { return "failback" }
func (c *EndpointManualFailback) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return http.MethodPost, fmt.Sprintf("/lb_domains/%s/sites/%s/endpoints/%s/failback", c.LBDomainID, c.SiteResourceName, c.EndpointResourceName)
	}
	return "", ""
}

func (c *EndpointManualFailback) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.LBDomainID, &c.SiteResourceName, &c.EndpointResourceName)
}

func init() {
	register(&EndpointManualFailover{}, &EndpointManualFailback{})
}
