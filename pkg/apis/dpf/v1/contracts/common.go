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
	GetContractId() string
	SetContractId(string)
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
	ContractId string
}

func (s *AttributeMeta) GetGroup() string        { return groupName }
func (s *AttributeMeta) GetContractId() string   { return s.ContractId }
func (s *AttributeMeta) SetContractId(id string) { s.ContractId = id }

func GetPathMethodForChildSpec(action api.Action, s ChildSpec) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/%s", s.GetContractId(), s.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/contracts/%s/%s/%d", s.GetContractId(), s.GetName(), s.GetId())
	}
	return "", ""
}

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	switch action {
	case api.ActionList:
		return http.MethodGet, fmt.Sprintf("/contracts/%s/%s", s.GetContractId(), s.GetName())
	case api.ActionCount:
		if _, ok := s.(api.CountableListSpec); ok {
			return http.MethodGet, fmt.Sprintf("/contracts/%s/%s/count", s.GetContractId(), s.GetName())
		}
	}
	return "", ""
}
