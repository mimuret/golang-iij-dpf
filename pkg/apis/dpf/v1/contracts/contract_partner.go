package contracts

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type ContractPartner struct {
	ServiceCode string `read:"service_code"`
}

var _ ListSpec = &ContractPartnerList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type ContractPartnerList struct {
	AttributeMeta
	Items []ContractPartner `read:"items"`
}

func (c *ContractPartnerList) GetName() string         { return "contract_partners" }
func (c *ContractPartnerList) GetItems() interface{}   { return &c.Items }
func (c *ContractPartnerList) Len() int                { return len(c.Items) }
func (c *ContractPartnerList) Index(i int) interface{} { return c.Items[i] }

func (c *ContractPartnerList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *ContractPartnerList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractId)
}

func (c *ContractPartnerList) Init() {}

func init() {
	register(&ContractPartnerList{})
}
