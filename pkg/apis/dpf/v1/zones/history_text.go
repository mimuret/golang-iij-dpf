package zones

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ Spec = &HistoryText{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type HistoryText struct {
	AttributeMeta
	History
	Text string `read:"text"`
}

func (c *HistoryText) GetName() string { return fmt.Sprintf("zone_histories/%d/text", c.Id) }
func (c *HistoryText) GetPathMethod(action api.Action) (string, string) {
	if action == api.ActionRead {
		return action.ToMethod(), fmt.Sprintf("/zones/%s/zone_histories/%d/text", c.GetZoneId(), c.Id)
	}
	return "", ""
}

func (c *HistoryText) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ZoneId, &c.Id)
}

func init() {
	register(&HistoryText{})
}
