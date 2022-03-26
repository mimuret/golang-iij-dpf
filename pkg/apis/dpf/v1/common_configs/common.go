package common_configs

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

const groupName = "common-configs.api.dns-platform.jp/v1"

func register(items ...apis.Spec) {
	schema.NewRegister(groupName).Add(items...)
}

type Spec interface {
	apis.Spec
	SetCommonConfigId(int64)
	GetCommonConfigId() int64
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

type AttributeMeta struct {
	CommonConfigId int64 `read:"-" id:"1,required"`
}

func (s *AttributeMeta) GetGroup() string           { return groupName }
func (s *AttributeMeta) SetCommonConfigId(id int64) { s.CommonConfigId = id }
func (s *AttributeMeta) GetCommonConfigId() int64   { return s.CommonConfigId }

func GetPathMethodForChildSpec(action api.Action, s ChildSpec) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/common_configs/%d/%s", s.GetCommonConfigId(), s.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/common_configs/%d/%s/%d", s.GetCommonConfigId(), s.GetName(), s.GetId())
	}
	return "", ""
}

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	if action == api.ActionList {
		return action.ToMethod(), fmt.Sprintf("/common_configs/%d/%s", s.GetCommonConfigId(), s.GetName())
	}
	return "", ""
}
