package zones

import (
	"fmt"
	"net/url"

	"strings"

	"github.com/google/go-querystring/query"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis"
)

type RecordState int

const (
	RecordStateApplied      RecordState = 0
	RecordStateToBeAdded    RecordState = 1
	RecordStateToBeDeleted  RecordState = 2
	RecordStateToBeUpdate   RecordState = 3
	RecordStateBeforeUpdate RecordState = 5
)

var RecordStateToString = map[RecordState]string{
	RecordStateApplied:      "Applied",
	RecordStateToBeAdded:    "ToBeAdded",
	RecordStateToBeDeleted:  "ToBeDeleted",
	RecordStateToBeUpdate:   "ToBeUpdate",
	RecordStateBeforeUpdate: "BeforeUpdate",
}

func (c RecordState) String() string {
	return RecordStateToString[c]
}

type Type string

const (
	TypeSOA   Type = "SOA"
	TypeA     Type = "A"
	TypeAAAA  Type = "AAAA"
	TypeCAA   Type = "CAA"
	TypeCNAME Type = "CNAME"
	TypeDS    Type = "DS"
	TypeNS    Type = "NS"
	TypeMX    Type = "MX"
	TypeNAPTR Type = "NAPTR"
	TypeSRV   Type = "SRV"
	TypeTXT   Type = "TXT"
	TypeTLSA  Type = "TLSA"
	TypePTR   Type = "PTR"

	TypeANAME Type = "ANAME"
)

func (c Type) String() string {
	return string(c)
}

// +k8s:deepcopy-gen=false
type KeywordsType []Type

type RecordRDATA struct {
	Value string `read:"value" create:"value" update:"value"`
}

func (c *RecordRDATA) String() string {
	return c.Value
}

type RecordRDATAs []RecordRDATA

func (c RecordRDATAs) String() string {
	var res []string
	for _, value := range c {
		res = append(res, value.String())
	}
	return strings.Join(res, ",")
}

var _ Spec = &Record{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type Record struct {
	AttributeMeta

	Id          string       `read:"id"`
	Name        string       `read:"name" create:"name"`
	TTL         int32        `read:"ttl"  create:"ttl" update:"ttl"`
	RRType      Type         `read:"rrtype"  create:"rrtype"`
	RData       RecordRDATAs `read:"rdata"  create:"rdata" update:"rdata"`
	State       RecordState  `read:"state"`
	Description string       `read:"description"  create:"description" update:"description"`
	Operator    string       `read:"operator"`
}

func (c *Record) GetName() string { return "records" }
func (c *Record) GetPathMethod(action api.Action) (string, string) {
	switch action {
	case api.ActionCreate:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s", c.GetZoneId(), c.GetName())
	case api.ActionRead, api.ActionUpdate, api.ActionDelete:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s/%s", c.GetZoneId(), c.GetName(), c.Id)
	case api.ActionCancel:
		return action.ToMethod(), fmt.Sprintf("/zones/%s/%s/%s/changes", c.GetZoneId(), c.GetName(), c.Id)
	}
	return "", ""
}
func (c *Record) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId, &c.Id)
}

var _ CountableListSpec = &RecordList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type RecordList struct {
	AttributeMeta
	api.Count
	Items []Record `read:"items"`
}

func (c *RecordList) GetName() string         { return "records" }
func (c *RecordList) GetItems() interface{}   { return &c.Items }
func (c *RecordList) Len() int                { return len(c.Items) }
func (c *RecordList) Index(i int) interface{} { return c.Items[i] }
func (c *RecordList) GetMaxLimit() int32      { return 10000 }
func (c *RecordList) ClearItems()             { c.Items = []Record{} }
func (c *RecordList) AddItem(v interface{}) bool {
	if a, ok := v.(Record); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *RecordList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *RecordList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}
func (c *RecordList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}

var _ CountableListSpec = &CurrentRecordList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type CurrentRecordList struct {
	AttributeMeta
	api.Count
	Items []Record `read:"items"`
}

func (c *CurrentRecordList) GetName() string         { return "records/currents" }
func (c *CurrentRecordList) Len() int                { return len(c.Items) }
func (c *CurrentRecordList) GetItems() interface{}   { return &c.Items }
func (c *CurrentRecordList) Index(i int) interface{} { return c.Items[i] }
func (c *CurrentRecordList) GetMaxLimit() int32      { return 10000 }
func (c *CurrentRecordList) ClearItems()             { c.Items = []Record{} }
func (c *CurrentRecordList) AddItem(v interface{}) bool {
	if a, ok := v.(Record); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *CurrentRecordList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *CurrentRecordList) Init() {
	for i := range c.Items {
		c.Items[i].AttributeMeta = c.AttributeMeta
	}
}
func (c *CurrentRecordList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

type RecordDiff struct {
	New *Record `read:"new"`
	Old *Record `read:"old"`
}

var _ CountableListSpec = &RecordDiffList{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/api.Object

type RecordDiffList struct {
	AttributeMeta
	api.Count
	Items []RecordDiff `read:"items"`
}

func (c *RecordDiffList) GetName() string         { return "records/diffs" }
func (c *RecordDiffList) GetItems() interface{}   { return &c.Items }
func (c *RecordDiffList) Len() int                { return len(c.Items) }
func (c *RecordDiffList) Index(i int) interface{} { return c.Items[i] }
func (c *RecordDiffList) GetMaxLimit() int32      { return 10000 }
func (c *RecordDiffList) ClearItems()             { c.Items = []RecordDiff{} }
func (c *RecordDiffList) AddItem(v interface{}) bool {
	if a, ok := v.(RecordDiff); ok {
		c.Items = append(c.Items, a)
		return true
	}
	return false
}

func (c *RecordDiffList) GetPathMethod(action api.Action) (string, string) {
	return GetPathMethodForListSpec(action, c)
}
func (c *RecordDiffList) SetParams(args ...interface{}) error {
	return apis.SetParams(args, &c.ZoneId)
}

func (c *RecordDiffList) Init() {
	for i := range c.Items {
		if c.Items[i].New != nil {
			c.Items[i].New.AttributeMeta = c.AttributeMeta
		}
		if c.Items[i].Old != nil {
			c.Items[i].Old.AttributeMeta = c.AttributeMeta
		}
	}
}

var _ api.SearchParams = &RecordListSearchKeywords{}

// +k8s:deepcopy-gen=false

type RecordListSearchKeywords struct {
	api.CommonSearchParams
	FullText    api.KeywordsString `url:"_keywords_full_text[],omitempty"`
	Name        api.KeywordsString `url:"_keywords_name[],omitempty"`
	TTL         []int32            `url:"_keywords_ttl[],omitempty"`
	RRType      KeywordsType       `url:"_keywords_rrtype[],omitempty"`
	RData       api.KeywordsString `url:"_keywords_rdata[],omitempty"`
	Description api.KeywordsString `url:"_keywords_description[],omitempty"`
	Operator    api.KeywordsString `url:"_keywords_operator[],omitempty"`
}

func (s *RecordListSearchKeywords) GetValues() (url.Values, error) { return query.Values(s) }

func init() {
	Register.Add(&Record{}, &RecordList{})
	Register.Add(&CurrentRecordList{})
	Register.Add(&RecordDiffList{})
}
