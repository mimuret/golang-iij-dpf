package common_configs

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

// for IDE
var _ ChildSpec = &CcSecTransferAcl{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcSecTransferAcl struct {
	AttributeMeta
	ID      int64                       `read:"id" id:"2,required"`
	Network *types.IPNet                `read:"network" create:"network" update:"network"`
	TsigID  types.NullablePositiveInt64 `read:"tsig_id"  create:"tsig_id" update:"tsig_id"`
}

func (c *CcSecTransferAcl) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID, &c.ID)
}

func (c *CcSecTransferAcl) GetID() int64    { return c.ID }
func (c *CcSecTransferAcl) SetID(id int64)  { c.ID = id }
func (c *CcSecTransferAcl) GetName() string { return "cc_sec_transfer_acls" }
func (c *CcSecTransferAcl) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

// for IDE
var _ ListSpec = &CcSecTransferAclList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcSecTransferAclList struct {
	AttributeMeta
	Items []CcSecTransferAcl `read:"items"`
}

func (c *CcSecTransferAclList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID)
}

func (c *CcSecTransferAclList) GetName() string         { return "cc_sec_transfer_acls" }
func (c *CcSecTransferAclList) GetItems() interface{}   { return &c.Items }
func (c *CcSecTransferAclList) Len() int                { return len(c.Items) }
func (c *CcSecTransferAclList) Index(i int) interface{} { return c.Items[i] }

func (c *CcSecTransferAclList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *CcSecTransferAclList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func init() {
	register(&CcSecTransferAcl{}, &CcSecTransferAclList{})
}
