package apiutils_test

import (
	"context"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/zones"
	"github.com/mimuret/golang-iij-dpf/pkg/apiutils"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"
	"github.com/mimuret/golang-iij-dpf/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("zone", func() {
	var (
		err    error
		c      *testtool.TestClient
		zoneId string
		z      *core.Zone
		r      *zones.Record
	)
	BeforeEach(func() {
		z = &core.Zone{
			Id:               "m1",
			CommonConfigId:   1,
			ServiceCode:      "dpm0000001",
			State:            1,
			Favorite:         1,
			Name:             "example.jp.",
			Network:          "",
			Description:      "zone 1",
			ZoneProxyEnabled: types.Disabled,
		}
		r = &zones.Record{
			AttributeMeta: zones.AttributeMeta{
				ZoneId: "m1",
			},
			Id:     "r1",
			Name:   "example.jp.",
			TTL:    3600,
			RRType: zones.TypeSOA,
			RData: zones.RecordRDATASlice{
				zones.RecordRDATA{
					Value: "ns000.d-53.net. dns-managers.iij.ad.jp. 30 3600 600 604800 900",
				},
			},
			State: zones.RecordStateApplied,
		}

		c = testtool.NewTestClient("token", "http://localhost", nil)
	})
	Context("GetZoneIdFromZonename", func() {
		When("failed to read", func() {
			BeforeEach(func() {
				zoneId, err = apiutils.GetZoneIdFromZonename(context.Background(), c, "example.jp.")
			})
			It("return empty error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to search zone"))
			})
		})
		When("find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					s.AddItem(*z)
					s.AddItem(core.Zone{
						Id:   "m2",
						Name: "hoge.example.jp.",
					})
					return "", nil
				}
				zoneId, err = apiutils.GetZoneIdFromZonename(context.Background(), c, "example.jp.")
			})
			It("returns zoneId", func() {
				Expect(err).To(Succeed())
				Expect(zoneId).To(Equal("m1"))
			})
		})
		When("not find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					s.AddItem(core.Zone{
						Id:   "m2",
						Name: "hoge.example.jp.",
					})
					return "", nil
				}
				zoneId, err = apiutils.GetZoneIdFromZonename(context.Background(), c, "example.jp.")
			})
			It("returns err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found zone"))
			})
		})
	})
	Context("GetZoneFromZonename", func() {
		var z1 *core.Zone
		When("failed to read", func() {
			BeforeEach(func() {
				z1, err = apiutils.GetZoneFromZonename(context.Background(), c, "example.jp.")
			})
			It("return empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to search zone"))
				Expect(z1).To(BeNil())
			})
		})
		When("find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					s.AddItem(*z)
					s.AddItem(core.Zone{
						Id:   "m2",
						Name: "hoge.example.jp.",
					})
					return "", nil
				}
				z1, err = apiutils.GetZoneFromZonename(context.Background(), c, "example.jp.")
			})
			It("return zone", func() {
				Expect(err).To(Succeed())
				Expect(z1).To(Equal(z))
			})
		})
		When("not find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					s.AddItem(core.Zone{
						Id:   "m2",
						Name: "hoge.example.jp.",
					})
					return "", nil
				}
				z1, err = apiutils.GetZoneFromZonename(context.Background(), c, "example.jp.")
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found zone"))
			})
		})
	})
	Context("GetRecordFromZoneName", func() {
		var r1 *zones.Record
		When("failed to read", func() {
			BeforeEach(func() {
				r1, err = apiutils.GetRecordFromZoneName(context.Background(), c, "example.jp.", "example.jp.", zones.TypeSOA)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to search zone"))
			})
		})
		When("find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m1",
							Name: "example.jp.",
						})
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					case *zones.CurrentRecordList:
						v.AddItem(*r)
					}
					return "", nil
				}
				r1, err = apiutils.GetRecordFromZoneName(context.Background(), c, "example.jp.", "example.jp.", zones.TypeSOA)
			})
			It("returns record", func() {
				Expect(err).To(Succeed())
				Expect(r1).To(Equal(r))
			})
		})
		When("not find zone", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					}
					return "ok", nil
				}
				r1, err = apiutils.GetRecordFromZoneName(context.Background(), c, "example.jp.", "example.jp.", zones.TypeSOA)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found zone"))
			})
		})
		When("not find records", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m1",
							Name: "example.jp.",
						})
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					case *zones.CurrentRecordList:
						v.AddItem(zones.Record{
							AttributeMeta: zones.AttributeMeta{
								ZoneId: "m1",
							},
							Id:     "r2",
							Name:   "example.jp.",
							TTL:    3600,
							RRType: zones.TypeNS,
							RData: zones.RecordRDATASlice{
								zones.RecordRDATA{
									Value: "ns1.example.jp.",
								},
								zones.RecordRDATA{
									Value: "ns2.example.jp.",
								},
							},
							State: zones.RecordStateApplied,
						})
					}
					return "", nil
				}
				r1, err = apiutils.GetRecordFromZoneName(context.Background(), c, "example.jp.", "example.jp.", zones.TypeSOA)
			})
			It("return error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found record"))
				Expect(r1).To(BeNil())
			})
		})
	})
	Context("GetRecordFromZoneId", func() {
		var r2 *zones.Record
		When("failed to read", func() {
			BeforeEach(func() {
				r2, err = apiutils.GetRecordFromZoneId(context.Background(), c, "m1", "example.jp.", zones.TypeSOA)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to search records"))
			})
		})
		When("find", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m1",
							Name: "example.jp.",
						})
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					case *zones.CurrentRecordList:
						v.AddItem(*r)
					}
					return "", nil
				}
				r2, err = apiutils.GetRecordFromZoneId(context.Background(), c, "m1", "example.jp.", zones.TypeSOA)
			})
			It("return record", func() {
				Expect(err).To(Succeed())
				Expect(r2).To(Equal(r))
			})
		})
		When("not find zone", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					}
					return "ok", nil
				}
				r2, err = apiutils.GetRecordFromZoneId(context.Background(), c, "m1", "example.jp.", zones.TypeSOA)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found record"))
			})
		})
		When("not find records", func() {
			BeforeEach(func() {
				c.ListAllFunc = func(s api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
					switch v := s.(type) {
					case *core.ZoneList:
						v.AddItem(core.Zone{
							Id:   "m1",
							Name: "example.jp.",
						})
						v.AddItem(core.Zone{
							Id:   "m2",
							Name: "hoge.example.jp.",
						})
					case *zones.CurrentRecordList:
						v.AddItem(zones.Record{
							AttributeMeta: zones.AttributeMeta{
								ZoneId: "m1",
							},
							Id:     "r2",
							Name:   "example.jp.",
							TTL:    3600,
							RRType: zones.TypeNS,
							RData: zones.RecordRDATASlice{
								zones.RecordRDATA{
									Value: "ns1.example.jp.",
								},
								zones.RecordRDATA{
									Value: "ns2.example.jp.",
								},
							},
							State: zones.RecordStateApplied,
						})
					}
					return "", nil
				}
				r2, err = apiutils.GetRecordFromZoneId(context.Background(), c, "m1", "example.jp.", zones.TypeSOA)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not found record"))
			})
		})
	})
})
