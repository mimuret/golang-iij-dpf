package lb_domains

type MonitoringPropsLocation string

const (
	MonitoringPropsLocationAll MonitoringPropsLocation = "all"
	MonitoringPropsLocationJP  MonitoringPropsLocation = "jp"
	MonitoringPropsLocationUS  MonitoringPropsLocation = "us"
)

type MonitoringPorpsCommon struct {
	Location MonitoringPropsLocation `read:"location" update:"location" create:"location,omitempty" apply:"location,omitempty"`
	Interval uint16                  `read:"interval" update:"interval" create:"interval,omitempty" apply:"interval,omitempty"`
	Holdtime uint16                  `read:"holdtime" update:"holdtime" create:"holdtime" apply:"holdtime"`
	Timeout  uint16                  `read:"timeout" update:"timeout" create:"timeout,omitempty" apply:"timeout,omitempty"`
}

var _ MonitoringPorps = &MonitoringPorpsPING{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.MonitoringPorps
type MonitoringPorpsPING struct {
	MonitoringPorpsCommon
}

func (m *MonitoringPorpsPING) GetMtype() MonitoringMtype {
	return MonitoringMtypePing
}

var _ MonitoringPorps = &MonitoringPorpsTCP{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.MonitoringPorps
type MonitoringPorpsTCP struct {
	MonitoringPorpsCommon
	Port       uint16 `read:"port" update:"port" create:"port" apply:"port"`
	TLSEnabled bool   `read:"tls_enabled" update:"tls_enabled" create:"tls_enabled" apply:"tls_enalbed"`
	TLSSNI     string `read:"tls_sni" update:"tls_sni" create:"tls_sni,omitempty" apply:"tls_sni,omitempty"`
}

func (m *MonitoringPorpsTCP) GetMtype() MonitoringMtype {
	return MonitoringMtypeTCP
}

type MonitoringPorpsHeader struct {
	FieldName  string `read:"field_name" update:"field_name" create:"field_name" apply:"field_name"`
	FieldValue string `read:"field_value" update:"field_value" create:"field_value" apply:"field_value"`
}

var _ MonitoringPorps = &MonitoringPorpsHTTP{}

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.MonitoringPorps
type MonitoringPorpsHTTP struct {
	MonitoringPorpsCommon
	Port          uint16   `read:"port" update:"port" create:"port,omitempty" apply:"port,omitempty"`
	TLSSNI        string   `read:"tls_sni" update:"tls_sni" create:"tls_sni,omitempty" apply:"tls_sni,omitempty"`
	ResponseMatch string   `read:"response_match" update:"response_match" create:"response_match,omitempty" apply:"response_match,omitempty"`
	HTTPS         bool     `read:"https" update:"https" create:"https" apply:"https"`
	Path          string   `read:"path" update:"path" create:"path,omitempty" apply:"path,omitempty"`
	StatusCode    []string `read:"status_codes" update:"status_codes" create:"status_codes,omitempty" apply:"status_codes,omitempty"`
}

func (m *MonitoringPorpsHTTP) GetMtype() MonitoringMtype {
	return MonitoringMtypeHTTP
}

type MonitoringPorpsStaticStatus string

const (
	MonitoringPorpsStaticStatusUp     MonitoringPorpsStaticStatus = "up"
	MonitoringPorpsStaticStatusDown   MonitoringPorpsStaticStatus = "down"
	MonitoringPorpsStaticStatusUnkown MonitoringPorpsStaticStatus = "unknown"
)

// +k8s:deepcopy-gen:interfaces=github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains.MonitoringPorps
type MonitoringPorpsStatic struct {
	Result MonitoringPorpsStaticStatus `read:"result" update:"result" create:"result" apply:"result"`
}

func (m *MonitoringPorpsStatic) GetMtype() MonitoringMtype {
	return MonitoringMtypeStatic
}
