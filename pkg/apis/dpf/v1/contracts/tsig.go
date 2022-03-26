package contracts

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

var _ ChildSpec = &Tsig{}

type TsigAlgorithm int

const (
	TsigAlgorithmHMACSHA256 TsigAlgorithm = 0
)

func (c TsigAlgorithm) String() string {
	tsigAlgorithmToString := map[TsigAlgorithm]string{
		TsigAlgorithmHMACSHA256: "HMAC-SHA256",
	}
	return tsigAlgorithmToString[c]
}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type Tsig struct {
	AttributeMeta
	Id          int64         `read:"id"`
	Name        string        `read:"name" create:"name"`
	Algorithm   TsigAlgorithm `read:"algorithm"`
	Secret      string        `read:"secret"`
	Description string        `read:"description" create:"description" update:"description"`
}

func (c *Tsig) GetName() string { return "tsigs" }
func (c *Tsig) GetID() int64    { return c.Id }
func (c *Tsig) SetID(id int64)  { c.Id = id }
func (c *Tsig) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForChildSpec(action, c)
}

func (c *Tsig) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID, &c.Id)
}

var _ ListSpec = &TsigList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object
type TsigList struct {
	AttributeMeta
	api.Count
	Items []Tsig `read:"items"`
}

func (c *TsigList) GetName() string         { return "tsigs" }
func (c *TsigList) GetItems() interface{}   { return &c.Items }
func (c *TsigList) Len() int                { return len(c.Items) }
func (c *TsigList) Index(i int) interface{} { return c.Items[i] }
func (c *TsigList) GetMaxLimit() int32      { return 10000 }
func (c *TsigList) ClearItems()             { c.Items = []Tsig{} }
func (c *TsigList) AddItem(v interface{}) bool {
	if a, ok := v.(Tsig); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *TsigList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}

func (c *TsigList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

func (c *TsigList) SetPathParams(args ...interface{}) error {
	return apis.SetPathParams(args, &c.ContractID)
}

var _ api.SearchParams = &TsigListSearchKeywords{}

// +k8s:deepcopy-gen=false
type TsigListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	Name        api.KeywordsString `url:"_keywords_name[],omitempty"`
	Description api.KeywordsString `url:"_keywords_description[],omitempty"`
}

func (s *TsigListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	register(&Tsig{}, &TsigList{})
}
