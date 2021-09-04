package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

type DnssecState int

const (
	DnssecStateZoneClosed DnssecState = 0
	DnssecStateEnabling   DnssecState = 1
	DnssecStateEnable     DnssecState = 2
	DnssecStateDisabling  DnssecState = 3
	DnssecStateDisable    DnssecState = 4
)

var DnssecStateToString = map[DnssecState]string{
	DnssecStateZoneClosed: "ZoneClosed",
	DnssecStateEnabling:   "Enabling",
	DnssecStateEnable:     "Enable",
	DnssecStateDisabling:  "Disabling",
	DnssecStateDisable:    "Disable",
}

func (c DnssecState) String() string {
	return DnssecStateToString[c]
}

type DSState int

const (
	DSStateClose                         DSState = 0
	DSStateBeforeRegistration            DSState = 1
	DSStateWaitClearCacheForRegistration DSState = 2
	DSStateDisclose                      DSState = 3
	DSStateBeforeChange                  DSState = 4
	DSStateWaitClearCacheForChanged      DSState = 5
	DSStateBeforeDelete                  DSState = 6
	DSStateWaitClearCacheForDelete       DSState = 7
)

var DSStateToSString = map[DSState]string{
	DSStateClose:                         "Close",
	DSStateBeforeRegistration:            "BeforeRegistration",
	DSStateWaitClearCacheForRegistration: "WaitRegistration",
	DSStateDisclose:                      "Disclose",
	DSStateBeforeChange:                  "BeforeChange",
	DSStateWaitClearCacheForChanged:      "WaitChange",
	DSStateBeforeDelete:                  "BeforeDelete",
	DSStateWaitClearCacheForDelete:       "WaitDelete",
}

func (c DSState) String() string {
	return DSStateToSString[c]
}

var _ Spec = &Dnssec{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type Dnssec struct {
	AttributeMeta
	Enabled types.Boolean `read:"enabled" update:"enabled"`
	State   DnssecState   `read:"state"`
	DsState DSState       `read:"ds_state"`
}

func (c *Dnssec) GetName() string { return "dnssec" }
func (c *Dnssec) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionRead, api.ActionUpdate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/dnssec", c.GetZoneId())
	}
	return "", ""
}
func (c *Dnssec) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneId)
}

func init() {
	Register.Add(&Dnssec{})
}