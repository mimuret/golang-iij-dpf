package types

import (
	"encoding/json"
	"net"
)

func ParseIPNet(str string) (*IPNet, error) {
	_, ipnet, err := net.ParseCIDR(str)
	if err != nil {
		return nil, err
	}
	return &IPNet{*ipnet}, nil
}

// for testing
func MustParseIPNet(str string) *IPNet {
	n, err := ParseIPNet(str)
	if err != nil {
		panic(err)
	}
	return n
}

// +k8s:deepcopy-gen=false
type IPNet struct {
	net.IPNet
}

func (i IPNet) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *IPNet) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return err
	}
	i.IPNet = *ipnet
	return nil
}

func (i *IPNet) DeepCopyInto(in *IPNet) {
	copy(i.IPNet.IP, in.IPNet.IP)
	copy(i.IPNet.Mask, in.IPNet.Mask)
}

func (i *IPNet) DeepCopy() *IPNet {
	res := &IPNet{}
	res.DeepCopyInto(i)
	return res
}
