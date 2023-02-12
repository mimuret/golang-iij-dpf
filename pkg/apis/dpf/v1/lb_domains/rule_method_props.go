package lb_domains

type RuleMethodMType string

const (
	RuleMethodMTypeEntryA     RuleMethodMType = "entry_a"
	RuleMethodMTypeEntryAAAA  RuleMethodMType = "entry_aaaa"
	RuleMethodMTypeEntryCNAME RuleMethodMType = "entry_cname"
	RuleMethodMTypeExitSite   RuleMethodMType = "exit_site"
	RuleMethodMTypeExitSorry  RuleMethodMType = "exit_sorry"
	RuleMethodMTypeFailover   RuleMethodMType = "exit_failover"
)

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodPropsCommon struct {
	ResourceName string          `read:"resource_name" create:"resource_name" apply:"resource_name"`
	MType        RuleMethodMType `read:"mtype" create:"mtype" apply:"mtype"`
	Enabled      bool            `read:"enabled" create:"enabled" update:"enabled" apply:"enabled"`
	LiveStatus   Status          `read:"live_status"`
	ReadyStatus  Status          `read:"ready_status"`
}

func (r *RuleMethodPropsCommon) Fix() {}

func (r *RuleMethodPropsCommon) GetMType() RuleMethodMType {
	return r.MType
}

func (r *RuleMethodPropsCommon) GetMethodResourceName() string {
	return r.ResourceName
}

func (r *RuleMethodPropsCommon) SetMethodResourceName(resourceName string) {
	r.ResourceName = resourceName
}

var _ RuleMethodProps = &RuleMethodEntryA{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodEntryA struct {
	RuleMethodPropsCommon
}

func (m *RuleMethodEntryA) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeEntryA
}

var _ RuleMethodProps = &RuleMethodEntryAAAA{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodEntryAAAA struct {
	RuleMethodPropsCommon
}

func (m *RuleMethodEntryAAAA) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeEntryAAAA
}

var _ RuleMethodProps = &RuleMethodEntryCNAME{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodEntryCNAME struct {
	RuleMethodPropsCommon
}

func (m *RuleMethodEntryCNAME) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeEntryCNAME
}

var _ RuleMethodProps = &RuleMethodExitSite{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodExitSite struct {
	RuleMethodPropsCommon
	ParentResourceName string `read:"parent_resource_name" create:"parent_resource_name" apply:"parent_resource_name"`
	SiteResourceName   string `read:"site_resource_name" create:"site_resource_name" apply:"site_resource_name"`
}

func (m *RuleMethodExitSite) GetParentResourceName() string {
	return m.ParentResourceName
}

func (m *RuleMethodExitSite) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeExitSite
}

var _ RuleMethodProps = &RuleMethodExitSorry{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodExitSorry struct {
	RuleMethodPropsCommon
	ParentResourceName string `read:"parent_resource_name" create:"parent_resource_name" apply:"parent_resource_name"`
}

func (m *RuleMethodExitSorry) GetParentResourceName() string {
	return m.ParentResourceName
}

func (m *RuleMethodExitSorry) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeExitSorry
}

var _ RuleMethodProps = &RuleMethodFailover{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.RuleMethodProps
type RuleMethodFailover struct {
	RuleMethodPropsCommon
	ParentResourceName string `read:"parent_resource_name" create:"parent_resource_name" apply:"parent_resource_name"`
}

func (m *RuleMethodFailover) GetParentResourceName() string {
	return m.ParentResourceName
}

func (m *RuleMethodFailover) Fix() {
	m.RuleMethodPropsCommon.MType = RuleMethodMTypeFailover
}
