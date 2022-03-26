package contracts

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

const groupName = "contracts.api.dns-platform.jp/v1"

func register(items ...apis.Spec) {
	schema.NewRegister(groupName).Add(items...)
}

type Spec interface {
	api.Spec
	apis.Params
	GetContractID() string
	SetContractID(string)
}

type ChildSpec interface {
	Spec
	GetID() int64
	SetID(int64)
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
	ContractID string
}

func (s *AttributeMeta) GetGroup() string        { return groupName }
func (s *AttributeMeta) GetContractID() string   { return s.ContractID }
func (s *AttributeMeta) SetContractID(id string) { s.ContractID = id }

func GetPathMethodForChildSpec(action api.Action, s ChildSpec) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/%s", s.GetContractID(), s.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/%s/%d", s.GetContractID(), s.GetName(), s.GetID())
	}
	return "", ""
}

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	switch action {
	case api.ActionList:
		return http.MethodGet, fmt.Sprintf("/contracts/%s/%s", s.GetContractID(), s.GetName())
	case api.ActionCount:
		if _, ok := s.(api.CountableListSpec); ok {
			return http.MethodGet, fmt.Sprintf("/contracts/%s/%s/count", s.GetContractID(), s.GetName())
		}
	}
	return "", ""
}
