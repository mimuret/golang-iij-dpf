package apiutils

import (
	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

func ZonenameToSystemId(cl api.ClientInterface, zonename string) string {
	zonename = dns.CanonicalName(zonename)
	keywords := &core.ZoneListSearchKeywords{
		Name: api.KeywordsString{zonename},
	}
	zoneList := &core.ZoneList{}
	if _, err := cl.ListALL(zoneList, keywords); err != nil {
		return ""
	}
	for _, zone := range zoneList.Items {
		if zonename == zone.Name {
			return zone.Id
		}
	}
	return ""
}
