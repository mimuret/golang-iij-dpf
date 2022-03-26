package contracts

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &CommonConfigDefault{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CommonConfigDefault struct {
	AttributeMeta
	ID int64 `apply:"common_config_id"`
}

func (c *CommonConfigDefault) GetName() string { return "common_configs/default" }
func (c *CommonConfigDefault) GetID() int64    { return c.ID }
func (c *CommonConfigDefault) SetID(id int64)  { c.ID = id }
func (c *CommonConfigDefault) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/common_configs/default", c.GetContractID())
	}
	return "", ""
}

func (c *CommonConfigDefault) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID)
}

func init() {
	register(&CommonConfigDefault{})
}
