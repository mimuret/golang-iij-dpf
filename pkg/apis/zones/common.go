package zones

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

var (
	groupName = "zones"
	Register  = schema.NewRegister(groupName)
)

type Spec interface {
	apis.Spec
	SetZoneId(string)
	GetZoneId() string
}

type ChildSpec interface {
	Spec
	GetId() int64
	SetId(int64)
}

type ListSpec interface {
	api.ListSpec
	Spec
}

type CountableListSpec interface {
	api.CountableListSpec
	Spec
}

type AttributeMeta struct {
	ZoneId string `read:"-"`
}

// for ctl
func (s *AttributeMeta) GetGroup() string    { return groupName }
func (s *AttributeMeta) SetZoneId(id string) { s.ZoneId = id }
func (s *AttributeMeta) GetZoneId() string   { return s.ZoneId }

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	switch action {
	case api.ActionList:
		return http.MethodGet, fmt.Sprintf("/zones/%s/%s", s.GetZoneId(), s.GetName())
	case api.ActionCount:
		if _, ok := s.(api.CountableListSpec); ok {
			return http.MethodGet, fmt.Sprintf("/zones/%s/%s/count", s.GetZoneId(), s.GetName())
		}
	}
	return "", ""
}

func GetReadPathMethodForSpec(action api.Action, s Spec) (string, string) {
	switch action {
	case api.ActionRead:
		return http.MethodGet, fmt.Sprintf("/zones/%s/%s", s.GetZoneId(), s.GetName())
	}
	return "", ""
}
