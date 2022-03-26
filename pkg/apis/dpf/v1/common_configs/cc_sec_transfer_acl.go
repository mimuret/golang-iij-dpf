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
	Id      int64        `read:"id" id:"2,required"`
	Network *types.IPNet `read:"network" create:"network" update:"network"`
	TsigId  int64        `read:"tsig_id,omitempty"  create:"tsig_id,omitempty" update:"tsig_id,omitempty"`
}

func (c *CcSecTransferAcl) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigId, &c.Id)
}

func (c *CcSecTransferAcl) GetId() int64    { return c.Id }
func (c *CcSecTransferAcl) SetId(id int64)  { c.Id = id }
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
	return apis.SetPathParams(args, &c.CommonConfigId)
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
