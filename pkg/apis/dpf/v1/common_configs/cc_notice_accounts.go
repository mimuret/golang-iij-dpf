package common_configs

import (
	"fmt"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type CcNoticeLang string

const (
	CcNoticeLangJA   CcNoticeLang = "ja"
	CcNoticeLangENUS CcNoticeLang = "en_US"
)

type CcNoticePhone struct {
	CountryCode string `read:"country_code" create:"country_code" update:"country_code"`
	Number      string `read:"number" create:"number" update:"number"`
}
type CcNoticeProps struct {
	Mail  string         `read:"mail" create:"mail,omitempty" update:"mail,omitempty"`
	Phone *CcNoticePhone `read:"phone" create:"phone,omitempty" update:"phone,omitempty"`
}

// for IDE.
var _ Spec = &CcNoticeAccount{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcNoticeAccount struct {
	AttributeMeta
	ResourceName string        `read:"resource_name,omitempty" create:"resource_name,omitempty" id:"2,required"`
	Name         string        `read:"name" create:"name" update:"name"`
	Lang         CcNoticeLang  `read:"lang" create:"lang" update:"lang"`
	Props        CcNoticeProps `read:"props" create:"props" update:"props"`
}

func (c *CcNoticeAccount) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID, &c.ResourceName)
}

func (c *CcNoticeAccount) GetName() string { return "cc_notice_accounts" }
func (c *CcNoticeAccount) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/common_configs/%d/cc_notice_accounts", c.GetCommonConfigID())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/common_configs/%d/cc_notice_accounts/%s", c.GetCommonConfigID(), c.ResourceName)
	}
	return "", ""
}

// for IDE.
var _ ListSpec = &CcNoticeAccountList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type CcNoticeAccountList struct {
	AttributeMeta
	Items []CcNoticeAccount `read:"items"`
}

func (c *CcNoticeAccountList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.CommonConfigID)
}

func (c *CcNoticeAccountList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *CcNoticeAccountList) GetName() string         { return "cc_notice_accounts" }
func (c *CcNoticeAccountList) GetItems() interface{}   { return &c.Items }
func (c *CcNoticeAccountList) Len() int                { return len(c.Items) }
func (c *CcNoticeAccountList) Index(i int) interface{} { return c.Items[i] }

func (c *CcNoticeAccountList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func init() {
	register(&CcNoticeAccount{}, &CcNoticeAccountList{})
}
