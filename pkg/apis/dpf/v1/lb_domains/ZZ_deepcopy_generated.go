//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
MIT License

Copyright (c) 2021 Manabu Sonoda

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package lb_domains

import (
	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeMeta) DeepCopyInto(out *AttributeMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeMeta.
func (in *AttributeMeta) DeepCopy() *AttributeMeta {
	if in == nil {
		return nil
	}
	out := new(AttributeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Monitorings != nil {
		in, out := &in.Monitorings, &out.Monitorings
		*out = make([]Monitoring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sites != nil {
		in, out := &in.Sites, &out.Sites
		*out = make([]Site, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Config) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	out.SiteAttributeMeta = in.SiteAttributeMeta
	if in.Rdata != nil {
		in, out := &in.Rdata, &out.Rdata
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Monitorings != nil {
		in, out := &in.Monitorings, &out.Monitorings
		*out = make([]MonitoringEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Endpoint) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointList) DeepCopyInto(out *EndpointList) {
	*out = *in
	out.SiteAttributeMeta = in.SiteAttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointList.
func (in *EndpointList) DeepCopy() *EndpointList {
	if in == nil {
		return nil
	}
	out := new(EndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *EndpointList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointManualFailback) DeepCopyInto(out *EndpointManualFailback) {
	*out = *in
	out.SiteAttributeMeta = in.SiteAttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointManualFailback.
func (in *EndpointManualFailback) DeepCopy() *EndpointManualFailback {
	if in == nil {
		return nil
	}
	out := new(EndpointManualFailback)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *EndpointManualFailback) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointManualFailover) DeepCopyInto(out *EndpointManualFailover) {
	*out = *in
	out.SiteAttributeMeta = in.SiteAttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointManualFailover.
func (in *EndpointManualFailover) DeepCopy() *EndpointManualFailover {
	if in == nil {
		return nil
	}
	out := new(EndpointManualFailover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *EndpointManualFailover) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogList) DeepCopyInto(out *LogList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]core.Log, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogList.
func (in *LogList) DeepCopy() *LogList {
	if in == nil {
		return nil
	}
	out := new(LogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *LogList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.MonitoringCommon = in.MonitoringCommon
	if in.Props != nil {
		out.Props = in.Props.DeepCopyMonitoringPorps()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Monitoring) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringCommon) DeepCopyInto(out *MonitoringCommon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringCommon.
func (in *MonitoringCommon) DeepCopy() *MonitoringCommon {
	if in == nil {
		return nil
	}
	out := new(MonitoringCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringEndpoint) DeepCopyInto(out *MonitoringEndpoint) {
	*out = *in
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(Monitoring)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringEndpoint.
func (in *MonitoringEndpoint) DeepCopy() *MonitoringEndpoint {
	if in == nil {
		return nil
	}
	out := new(MonitoringEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringList) DeepCopyInto(out *MonitoringList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Monitoring, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringList.
func (in *MonitoringList) DeepCopy() *MonitoringList {
	if in == nil {
		return nil
	}
	out := new(MonitoringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *MonitoringList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsCommon) DeepCopyInto(out *MonitoringPorpsCommon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsCommon.
func (in *MonitoringPorpsCommon) DeepCopy() *MonitoringPorpsCommon {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsHTTP) DeepCopyInto(out *MonitoringPorpsHTTP) {
	*out = *in
	out.MonitoringPorpsCommon = in.MonitoringPorpsCommon
	if in.StatusCode != nil {
		in, out := &in.StatusCode, &out.StatusCode
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsHTTP.
func (in *MonitoringPorpsHTTP) DeepCopy() *MonitoringPorpsHTTP {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMonitoringPorps is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorps.
func (in *MonitoringPorpsHTTP) DeepCopyMonitoringPorps() MonitoringPorps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsHeader) DeepCopyInto(out *MonitoringPorpsHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsHeader.
func (in *MonitoringPorpsHeader) DeepCopy() *MonitoringPorpsHeader {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsPING) DeepCopyInto(out *MonitoringPorpsPING) {
	*out = *in
	out.MonitoringPorpsCommon = in.MonitoringPorpsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsPING.
func (in *MonitoringPorpsPING) DeepCopy() *MonitoringPorpsPING {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsPING)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMonitoringPorps is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorps.
func (in *MonitoringPorpsPING) DeepCopyMonitoringPorps() MonitoringPorps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsStatic) DeepCopyInto(out *MonitoringPorpsStatic) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsStatic.
func (in *MonitoringPorpsStatic) DeepCopy() *MonitoringPorpsStatic {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsStatic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMonitoringPorps is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorps.
func (in *MonitoringPorpsStatic) DeepCopyMonitoringPorps() MonitoringPorps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringPorpsTCP) DeepCopyInto(out *MonitoringPorpsTCP) {
	*out = *in
	out.MonitoringPorpsCommon = in.MonitoringPorpsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorpsTCP.
func (in *MonitoringPorpsTCP) DeepCopy() *MonitoringPorpsTCP {
	if in == nil {
		return nil
	}
	out := new(MonitoringPorpsTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMonitoringPorps is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringPorps.
func (in *MonitoringPorpsTCP) DeepCopyMonitoringPorps() MonitoringPorps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]RuleMethod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Rule) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleAttributeMeta) DeepCopyInto(out *RuleAttributeMeta) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleAttributeMeta.
func (in *RuleAttributeMeta) DeepCopy() *RuleAttributeMeta {
	if in == nil {
		return nil
	}
	out := new(RuleAttributeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleList) DeepCopyInto(out *RuleList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *RuleList) DeepCopy() *RuleList {
	if in == nil {
		return nil
	}
	out := new(RuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *RuleList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethod) DeepCopyInto(out *RuleMethod) {
	*out = *in
	out.RuleAttributeMeta = in.RuleAttributeMeta
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(uint)
		**out = **in
	}
	if in.Method != nil {
		out.Method = in.Method.DeepCopyRuleMethodProps()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethod.
func (in *RuleMethod) DeepCopy() *RuleMethod {
	if in == nil {
		return nil
	}
	out := new(RuleMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *RuleMethod) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodEntryA) DeepCopyInto(out *RuleMethodEntryA) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodEntryA.
func (in *RuleMethodEntryA) DeepCopy() *RuleMethodEntryA {
	if in == nil {
		return nil
	}
	out := new(RuleMethodEntryA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodEntryA) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodEntryAAAA) DeepCopyInto(out *RuleMethodEntryAAAA) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodEntryAAAA.
func (in *RuleMethodEntryAAAA) DeepCopy() *RuleMethodEntryAAAA {
	if in == nil {
		return nil
	}
	out := new(RuleMethodEntryAAAA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodEntryAAAA) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodEntryCNAME) DeepCopyInto(out *RuleMethodEntryCNAME) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodEntryCNAME.
func (in *RuleMethodEntryCNAME) DeepCopy() *RuleMethodEntryCNAME {
	if in == nil {
		return nil
	}
	out := new(RuleMethodEntryCNAME)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodEntryCNAME) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodExitSite) DeepCopyInto(out *RuleMethodExitSite) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodExitSite.
func (in *RuleMethodExitSite) DeepCopy() *RuleMethodExitSite {
	if in == nil {
		return nil
	}
	out := new(RuleMethodExitSite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodExitSite) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodExitSorry) DeepCopyInto(out *RuleMethodExitSorry) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodExitSorry.
func (in *RuleMethodExitSorry) DeepCopy() *RuleMethodExitSorry {
	if in == nil {
		return nil
	}
	out := new(RuleMethodExitSorry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodExitSorry) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodFailover) DeepCopyInto(out *RuleMethodFailover) {
	*out = *in
	out.RuleMethodPropsCommon = in.RuleMethodPropsCommon
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodFailover.
func (in *RuleMethodFailover) DeepCopy() *RuleMethodFailover {
	if in == nil {
		return nil
	}
	out := new(RuleMethodFailover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodFailover) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodList) DeepCopyInto(out *RuleMethodList) {
	*out = *in
	out.RuleAttributeMeta = in.RuleAttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RuleMethod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodList.
func (in *RuleMethodList) DeepCopy() *RuleMethodList {
	if in == nil {
		return nil
	}
	out := new(RuleMethodList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *RuleMethodList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleMethodPropsCommon) DeepCopyInto(out *RuleMethodPropsCommon) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodPropsCommon.
func (in *RuleMethodPropsCommon) DeepCopy() *RuleMethodPropsCommon {
	if in == nil {
		return nil
	}
	out := new(RuleMethodPropsCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyRuleMethodProps is an autogenerated deepcopy function, copying the receiver, creating a new RuleMethodProps.
func (in *RuleMethodPropsCommon) DeepCopyRuleMethodProps() RuleMethodProps {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Site) DeepCopyInto(out *Site) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Site.
func (in *Site) DeepCopy() *Site {
	if in == nil {
		return nil
	}
	out := new(Site)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Site) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteAttributeMeta) DeepCopyInto(out *SiteAttributeMeta) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteAttributeMeta.
func (in *SiteAttributeMeta) DeepCopy() *SiteAttributeMeta {
	if in == nil {
		return nil
	}
	out := new(SiteAttributeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteList) DeepCopyInto(out *SiteList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Site, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteList.
func (in *SiteList) DeepCopy() *SiteList {
	if in == nil {
		return nil
	}
	out := new(SiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *SiteList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
