package zones

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

var _ Spec = &Contract{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type Contract struct {
	AttributeMeta
	core.Contract
}

func (c *Contract) GetName() string { return "contract" }
func (c *Contract) GetPathMethod(action api.Action) (string, string) {
	return GetReadPathMethodForSpec(action, c)
}
func (c *Contract) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

func init() {
	Register.Add(&Contract{})
}
