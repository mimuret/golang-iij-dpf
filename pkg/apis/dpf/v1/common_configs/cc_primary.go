package common_configs

import (
	"net"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

// for IDE.
var _ ChildSpec = &CcPrimary{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcPrimary struct {
	AttributeMeta
	ID      int64                       `read:"id,omitempty" id:"2,required"`
	Address net.IP                      `read:"address" create:"address" update:"address"`
	TsigID  types.NullablePositiveInt64 `read:"tsig_id"  create:"tsig_id" update:"tsig_id"`
	Enabled types.Boolean               `read:"enabled" update:"enabled"`
}

func (c *CcPrimary) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID, &c.ID)
}

func (c *CcPrimary) GetID() int64    { return c.ID }
func (c *CcPrimary) SetID(id int64)  { c.ID = id }
func (c *CcPrimary) GetName() string { return "cc_primaries" }
func (c *CcPrimary) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

// for IDE.
var _ ListSpec = &CcPrimaryList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcPrimaryList struct {
	AttributeMeta
	Items []CcPrimary `read:"items"`
}

func (c *CcPrimaryList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID)
}

func (c *CcPrimaryList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *CcPrimaryList) GetName() string         { return "cc_primaries" }
func (c *CcPrimaryList) GetItems() interface{}   { return &c.Items }
func (c *CcPrimaryList) Len() int                { return len(c.Items) }
func (c *CcPrimaryList) Index(i int) interface{} { return c.Items[i] }

func (c *CcPrimaryList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func init() {
	register(&CcPrimary{}, &CcPrimaryList{})
}
