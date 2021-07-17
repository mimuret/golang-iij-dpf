package apiutils

import (
	"fmt"

	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

func ZonenameToSystemId(cl api.ClientInterface, zonename string) (string, error) {
	zonename = dns.CanonicalName(zonename)
	keywords := &core.ZoneListSearchKeywords{
		Name: api.KeywordsString{zonename},
	}
	zoneList := &core.ZoneList{}
	if _, err := cl.ListALL(zoneList, keywords); err != nil {
		return "", fmt.Errorf("failed to search zone: %w", err)
	}
	for _, zone := range zoneList.Items {
		if zonename == zone.Name {
			return zone.Id, nil
		}
	}
	return "", nil
}
