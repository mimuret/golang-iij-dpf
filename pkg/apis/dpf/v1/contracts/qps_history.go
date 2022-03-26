package contracts

import (
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type QpsValue struct {
	Month string `read:"month"`
	Qps   int    `read:"qps"`
}

type QpsHistory struct {
	ServiceCode string     `read:"service_code"`
	Name        string     `read:"name"`
	Values      []QpsValue `read:"values"`
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type QpsHistoryList struct {
	AttributeMeta
	Items []QpsHistory `read:"items"`
}

var _ ListSpec = &QpsHistoryList{}

func (c *QpsHistoryList) GetName() string         { return "qps/histories" }
func (c *QpsHistoryList) GetItems() interface{}   { return &c.Items }
func (c *QpsHistoryList) Len() int                { return len(c.Items) }
func (c *QpsHistoryList) Index(i int) interface{} { return c.Items[i] }

// /contracts/{ContractID}/qps/histories
func (c *QpsHistoryList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *QpsHistoryList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID)
}

func (c *QpsHistoryList) Init() {}

func init() {
	register(&QpsHistoryList{})
}
