package apiutils

import (
	"fmt"

	"github.com/miekg/dns"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
)

func GetZoneIdFromZonename(cl api.ClientInterface, zonename string) (string, error) {
	z, err := GetZoneFromZonename(cl, zonename)
	if err != nil {
		return "", err
	}
	return z.Id, nil
}

func GetZoneFromZonename(cl api.ClientInterface, zonename string) (*core.Zone, error) {
	zonename = dns.CanonicalName(zonename)
	keywords := &core.ZoneListSearchKeywords{
		Name: api.KeywordsString{zonename},
	}
	zoneList := &core.ZoneList{}
	if _, err := cl.ListAll(zoneList, keywords); err != nil {
		return nil, fmt.Errorf("failed to search zone: %w", err)
	}
	for _, zone := range zoneList.Items {
		if zonename == zone.Name {
			return &zone, nil
		}
	}
	return nil, fmt.Errorf("not found zone: %s", zonename)
}

func GetRecordFromZoneName(cl api.ClientInterface, zonename string, recordName string, rrtype zones.Type) (*zones.Record, error) {
	z, err := GetZoneFromZonename(cl, zonename)
	if err != nil {
		return nil, err
	}
	return GetRecordFromZoneId(cl, z.Id, recordName, rrtype)
}

func GetRecordFromZoneId(cl api.ClientInterface, zoneId string, recordName string, rrtype zones.Type) (*zones.Record, error) {
	recordName = dns.CanonicalName(recordName)
	keywords := &zones.RecordListSearchKeywords{
		Name: api.KeywordsString{recordName},
	}
	currentList := &zones.CurrentRecordList{
		AttributeMeta: zones.AttributeMeta{
			ZoneId: zoneId,
		},
	}
	if _, err := cl.ListAll(currentList, keywords); err != nil {
		return nil, fmt.Errorf("failed to search records: %w", err)
	}
	for _, record := range currentList.Items {
		if recordName == record.Name && record.RRType == rrtype {
			return &record, nil
		}
	}
	return nil, fmt.Errorf("not found record: %s", recordName)
}
