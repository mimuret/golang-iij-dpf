package zones

import (
	"fmt"
	"net"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ Spec = &DefaultTTL{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ZoneProxy struct {
	AttributeMeta
	Enabled types.Boolean `read:"enabled" update:"enabled"`
}

func (c *ZoneProxy) GetName() string { return "zone_proxy" }
func (c *ZoneProxy) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/zone_proxy", c.GetZoneId())
	}
	return "", ""
}
func (c *ZoneProxy) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

type ZoneProxyStatus string

var (
	ZoneProxyStatusSuccess ZoneProxyStatus = "success"
	ZoneProxyStatusFail    ZoneProxyStatus = "fail"
)

type ZoneProxyHealthCheck struct {
	Address  net.IP          `read:"address"`
	Status   ZoneProxyStatus `read:"status"`
	TsigName string          `read:"tsig_name"`
	Enabled  types.Boolean   `read:"enabled"`
}

var _ ListSpec = &ZoneProxyHealthCheckList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type ZoneProxyHealthCheckList struct {
	AttributeMeta
	Items []ZoneProxyHealthCheck `read:"items"`
}

func (c *ZoneProxyHealthCheckList) GetName() string         { return "zone_proxy/health_check" }
func (c *ZoneProxyHealthCheckList) Len() int                { return len(c.Items) }
func (c *ZoneProxyHealthCheckList) GetItems() interface{}   { return &c.Items }
func (c *ZoneProxyHealthCheckList) Index(i int) interface{} { return c.Items[i] }

func (c *ZoneProxyHealthCheckList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *ZoneProxyHealthCheckList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

func (c *ZoneProxyHealthCheckList) Init() {}

func init() {
	Register.Add(&ZoneProxy{})
	Register.Add(&ZoneProxyHealthCheckList{})
}
