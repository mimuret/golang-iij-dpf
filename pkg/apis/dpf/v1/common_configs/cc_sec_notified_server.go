package common_configs

import (
	"net"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

// for IDE.
var _ ChildSpec = &CcSecNotifiedServer{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcSecNotifiedServer struct {
	AttributeMeta `read:"-" create:"-" update:"-"`
	Id            int64  `read:"id" create:"-" update:"-"  id:"2,required"`
	Address       net.IP `read:"address" create:"address" update:"address"`
	TsigId        int64  `read:"tsig_id,omitempty" create:"tsig_id,omitempty" update:"tsig_id,omitempty"`
}

func (c *CcSecNotifiedServer) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID, &c.Id)
}

func (c *CcSecNotifiedServer) GetID() int64    { return c.Id }
func (c *CcSecNotifiedServer) SetID(id int64)  { c.Id = id }
func (c *CcSecNotifiedServer) GetName() string { return "cc_sec_notified_servers" }
func (c *CcSecNotifiedServer) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

// for IDE
var _ ListSpec = &CcSecNotifiedServerList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcSecNotifiedServerList struct {
	AttributeMeta
	Items []CcSecNotifiedServer `read:"items"`
}

func (c *CcSecNotifiedServerList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID)
}

func (c *CcSecNotifiedServerList) GetName() string         { return "cc_sec_notified_servers" }
func (c *CcSecNotifiedServerList) GetItems() interface{}   { return &c.Items }
func (c *CcSecNotifiedServerList) Len() int                { return len(c.Items) }
func (c *CcSecNotifiedServerList) Index(i int) interface{} { return c.Items[i] }

func (c *CcSecNotifiedServerList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *CcSecNotifiedServerList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func init() {
	register(&CcSecNotifiedServer{}, &CcSecNotifiedServerList{})
}
