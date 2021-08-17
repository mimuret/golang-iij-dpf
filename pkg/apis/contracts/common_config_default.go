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
	Id int64 `apply:"common_config_id"`
}

func (c *CommonConfigDefault) GetName() string { return "common_configs/default" }
func (c *CommonConfigDefault) GetId() int64    { return c.Id }
func (c *CommonConfigDefault) SetId(id int64)  { c.Id = id }
func (c *CommonConfigDefault) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionApply:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/common_configs/default", c.GetContractId())
	}
	return "", ""
}
func (c *CommonConfigDefault) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractId)
}

func init() {
	Register.Add(&CommonConfigDefault{})
}
