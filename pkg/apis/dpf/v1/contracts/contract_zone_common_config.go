package contracts

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &ContractZoneCommonConfig{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type ContractZoneCommonConfig struct {
	AttributeMeta
	CommonConfigID int64    `apply:"common_config_id"`
	ZoneIDs        []string `apply:"zone_ids"`
}

func (c *ContractZoneCommonConfig) GetName() string { return "zones/common_configs" }
func (c *ContractZoneCommonConfig) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionApply {
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/zones/common_configs", c.GetContractID())
	}
	return "", ""
}

func (c *ContractZoneCommonConfig) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID)
}

func init() {
	register(&ContractZoneCommonConfig{})
}
