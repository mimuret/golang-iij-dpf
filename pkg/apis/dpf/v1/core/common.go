package core

import "github.com/mimuret/golang-iij-dpf/pkg/schema"

var (
	groupName = "core.api.dns-platform.jp/v1"
	Register  = schema.NewRegister(groupName)
)

type AttributeMeta struct {
}

func (s *AttributeMeta) GetGroup() string { return groupName }
