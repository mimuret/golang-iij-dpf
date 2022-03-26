package core_test

import (
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("log", func() {
	var (
		err    error
		s1     core.Log
		atTime types.Time
	)
	BeforeEach(func() {
		atTime, err = types.ParseTime(time.RFC3339Nano, "2021-06-20T13:21:12.881Z")
		Expect(err).To(Succeed())
		s1 = core.Log{
			Time:      atTime,
			LogType:   "service",
			Operator:  "user1",
			Operation: "add_cc_primary",
			Target:    "1",
			Status:    core.LogStatusStart,
		}
	})
	Describe("Log", func() {
		Context("DeepCopy", func() {
			var (
				copy    *core.Log
				nilMeta *core.Log
			)
			When("Log is not nil", func() {
				BeforeEach(func() {
					copy = s1.DeepCopy()
				})
				It("returns copy ", func() {
					Expect(copy).To(Equal(&s1))
				})
			})
			When("Log is nil", func() {
				BeforeEach(func() {
					copy = nilMeta.DeepCopy()
				})
				It("returns copy ", func() {
					Expect(copy).To(BeNil())
				})
			})
		})
	})
	Describe("LogStatus", func() {
		Context("String", func() {
			It("returns string", func() {
				Expect(core.LogStatusStart.String()).To(Equal("start"))
			})
		})
	})
	Describe("SearchLogsOffset", func() {
		Context("Validate", func() {
			When("value < 0", func() {
				It("returns false", func() {
					Expect(core.SearchLogsOffset(-1).Validate()).To(BeFalse())
				})
			})
			When("normal", func() {
				It("returns true", func() {
					Expect(core.SearchLogsOffset(0).Validate()).To(BeTrue())
					Expect(core.SearchLogsOffset(9900).Validate()).To(BeTrue())
				})
			})
			When("value > 9900", func() {
				It("returns false", func() {
					Expect(core.SearchLogsOffset(9901).Validate()).To(BeFalse())
				})
			})
		})
	})
	Describe("SearchLogsLimit", func() {
		Context("Validate", func() {
			When("value < 0", func() {
				It("returns false", func() {
					Expect(core.SearchLogsLimit(0).Validate()).To(BeFalse())
					Expect(core.SearchLogsLimit(-1).Validate()).To(BeFalse())
				})
			})
			When("normal", func() {
				It("returns true", func() {
					Expect(core.SearchLogsLimit(1).Validate()).To(BeTrue())
					Expect(core.SearchLogsLimit(100).Validate()).To(BeTrue())
				})
			})
			When("value > 9900", func() {
				It("returns false", func() {
					Expect(core.SearchLogsLimit(101).Validate()).To(BeFalse())
				})
			})
		})
	})
	Describe("TestLogListSearchKeywords", func() {
		Context("GetValues", func() {
			testcase := []struct {
				keyword core.LogListSearchKeywords
				values  url.Values
			}{
				{
					core.LogListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeAND,
							Offset: int32(10),
							Limit:  int32(100),
						},
					},
					url.Values{
						"type":   []string{"AND"},
						"offset": []string{"10"},
						"limit":  []string{"100"},
					},
				},
				{
					core.LogListSearchKeywords{
						CommonSearchParams: api.CommonSearchParams{
							Type:   api.SearchTypeOR,
							Offset: int32(10),
							Limit:  int32(100),
						},
						FullText:  api.KeywordsString{"hogehoge", "🐰"},
						LogType:   api.KeywordsString{"service", "record", "dnssec"},
						Operator:  api.KeywordsString{"rabbit@example.jp", "SA0000000"},
						Operation: api.KeywordsString{"updating_default_ttl", "dismiss_zone_edits"},
						Target:    api.KeywordsString{"hoge", "fuga"},
						Detail:    api.KeywordsString{"🐇", "🍺"},
						RequestID: api.KeywordsString{"f02fe1e1404140cab93c8f7af26081b7", "f098f808cde249d48a885490f7622df9"},
						Status:    core.KeywordsLogStatus{core.LogStatusStart, core.LogStatusSuccess, core.LogStatusFailure, core.LogStatusRetry},
					},
					/*
						_keywords_full_text[]
						_keywords_log_type[]
						_keywords_operator[]
						_keywords_operation[]
						_keywords_target[]
						_keywords_detail[]
						_keywords_request_id[]
						_keywords_status[]

					*/
					url.Values{
						"type":                   []string{"OR"},
						"offset":                 []string{"10"},
						"limit":                  []string{"100"},
						"_keywords_full_text[]":  []string{"hogehoge", "🐰"},
						"_keywords_log_type[]":   []string{"service", "record", "dnssec"},
						"_keywords_operator[]":   []string{"rabbit@example.jp", "SA0000000"},
						"_keywords_operation[]":  []string{"updating_default_ttl", "dismiss_zone_edits"},
						"_keywords_target[]":     []string{"hoge", "fuga"},
						"_keywords_detail[]":     []string{"🐇", "🍺"},
						"_keywords_request_id[]": []string{"f02fe1e1404140cab93c8f7af26081b7", "f098f808cde249d48a885490f7622df9"},
						"_keywords_status[]":     []string{"start", "success", "failure", "retry"},
					},
				},
			}
			It("can convert url.Value", func() {
				for _, tc := range testcase {
					s, err := tc.keyword.GetValues()
					Expect(err).To(Succeed())
					Expect(s).To(Equal(tc.values))
				}
			})
		})
	})
})
