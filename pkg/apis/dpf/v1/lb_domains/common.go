package lb_domains

import (
	"fmt"
	"net/http"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
	"github.com/mimuret/golang-iij-dpf/pkg/schema"
)

const groupName = "lb-domains.api.dns-platform.jp/v1"

func register(items ...apis.Spec) {
	schema.NewRegister(groupName).Add(items...)
}

type Spec interface {
	api.Spec
	apis.Params
	GetLBDoaminID() string
	SetLBDoaminID(string)
}

type ChildSpec interface {
	Spec
	GetResourceName() string
	SetResourceName(string)
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
	LBDomainID string
}

func (s *AttributeMeta) GetGroup() string        { return groupName }
func (s *AttributeMeta) GetLBDoaminID() string   { return s.LBDomainID }
func (s *AttributeMeta) SetLBDoaminID(id string) { s.LBDomainID = id }

func GetPathMethodForChildSpec(action api.Action, s ChildSpec) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/%s", s.GetLBDoaminID(), s.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/lb_domains/%s/%s/%s", s.GetLBDoaminID(), s.GetName(), s.GetResourceName())
	}
	return "", ""
}

func GetPathMethodForListSpec(action api.Action, s ListSpec) (string, string) {
	if action == api.ActionList {
		return http.MethodGet, fmt.Sprintf("/lb_domains/%s/%s", s.GetLBDoaminID(), s.GetName())
	}
	return "", ""
}

func GetPathMethodForCountableListSpec(action api.Action, s ListSpec) (string, string) {
	switch action {
	case api.ActionList:
		return http.MethodGet, fmt.Sprintf("/lb_domains/%s/%s", s.GetLBDoaminID(), s.GetName())
	case api.ActionCount:
		if _, ok := s.(api.CountableListSpec); ok {
			return http.MethodGet, fmt.Sprintf("/lb_domains/%s/%s/count", s.GetLBDoaminID(), s.GetName())
		}
	}
	return "", ""
}
