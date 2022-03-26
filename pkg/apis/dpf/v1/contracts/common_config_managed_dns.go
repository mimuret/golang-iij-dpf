package contracts

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ Spec = &CommonConfigManagedDns{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfigManagedDns struct {
	AttributeMeta
	ID                int64
	ManagedDnsEnabled types.Boolean `apply:"managed_dns_enabled"`
}

func (c *CommonConfigManagedDns) GetName() string {
	return fmt.Sprintf("common_configs/%d/managed_dns", c.ID)
}
func (c *CommonConfigManagedDns) GetID() int64   { return c.ID }
func (c *CommonConfigManagedDns) SetID(id int64) { c.ID = id }
func (c *CommonConfigManagedDns) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/common_configs/%d/managed_dns", c.GetContractID(), c.ID)
	}
	return "", ""
}

func (c *CommonConfigManagedDns) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID, &c.ID)
}

func init() {
	register(&CommonConfigManagedDns{})
}
