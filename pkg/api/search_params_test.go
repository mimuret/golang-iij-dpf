package api_test

import (
	"net/url"

	"github.com/google/go-querystring/query"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/types"
)

var _ = Describe("search_params", func() {
	Context("RowSearchParams", func() {
		var (
			err    error
			params *api.RowSearchParams
		)
		Context("NewRowSearchParams", func() {
			When("query string parse error", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("%1")
				})
				It("returns error", func() {
					Expect(err).To(HaveOccurred())
					Expect(params).To(BeNil())
				})
			})
			When("valid query string", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("hogehoge=1&offset=0&limit=100")
				})
				It("returns RowSearchParams", func() {
					Expect(err).To(Succeed())
					Expect(params).NotTo(BeNil())
				})
			})
		})
		Context("GetOffset", func() {
			When("limit not set", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("")
					Expect(err).To(Succeed())
				})
				It("returns 0", func() {
					Expect(params.GetOffset()).To(Equal(int32(0)))
				})
			})
			When("offset is set, not integer value", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("offset=AAA")
					Expect(err).To(Succeed())
				})
				It("returns 0", func() {
					Expect(params.GetOffset()).To(Equal(int32(0)))
				})
			})
			When("limit is set, integer value", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("offset=10")
					Expect(err).To(Succeed())
				})
				It("returns int value", func() {
					Expect(params.GetOffset()).To(Equal(int32(10)))
				})
			})
		})
		Context("SetOffset", func() {
			When("offset set to 1", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("")
					Expect(err).To(Succeed())
					params.SetOffset(1)
				})
				It("returns 1", func() {
					Expect(params.GetOffset()).To(Equal(int32(1)))
				})
			})
		})
		Context("GetLimit", func() {
			When("limit not set", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("")
					Expect(err).To(Succeed())
				})
				It("returns 100", func() {
					Expect(params.GetLimit()).To(Equal(int32(100)))
				})
			})
			When("limit is set, but not integer", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("limit=ABC")
					Expect(err).To(Succeed())
				})
				It("returns 100", func() {
					Expect(params.GetLimit()).To(Equal(int32(100)))
				})
			})
			When("limit is set, integer value", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("limit=10")
					Expect(err).To(Succeed())
				})
				It("returns int value", func() {
					Expect(params.GetLimit()).To(Equal(int32(10)))
				})
			})
		})
		Context("SetLimit", func() {
			When("limit set to 50", func() {
				BeforeEach(func() {
					params, err = api.NewRowSearchParams("")
					Expect(err).To(Succeed())
					params.SetLimit(50)
				})
				It("returns 50", func() {
					Expect(params.GetLimit()).To(Equal(int32(50)))
				})
			})
		})
		Context("GetValues", func() {
			BeforeEach(func() {
				params, err = api.NewRowSearchParams("limit=1&offset=3")
				Expect(err).To(Succeed())
			})
			It("returns url.Values", func() {
				v := url.Values{
					"limit":  []string{"1"},
					"offset": []string{"3"},
				}
				Expect(params.GetValues()).To(Equal(v))
			})
		})
	})
	Context("CommonSearchParams", func() {
		var params *api.CommonSearchParams
		BeforeEach(func() {
			params = &api.CommonSearchParams{
				Type:   api.SearchTypeAND,
				Offset: 1,
				Limit:  50,
			}
		})
		Context("GetType", func() {
			It("returns type", func() {
				Expect(params.GetType()).To(Equal(api.SearchTypeAND))
			})
		})
		Context("SetType", func() {
			It("can set type", func() {
				params.SetType(api.SearchTypeOR)
				Expect(params.GetType()).To(Equal(api.SearchTypeOR))
			})
		})
		Context("GetOffset", func() {
			It("returns offset", func() {
				Expect(params.GetOffset()).To(Equal(int32(1)))
			})
		})
		Context("SetOffset", func() {
			It("can set offset", func() {
				params.SetOffset(2)
				Expect(params.GetOffset()).To(Equal(int32(2)))
			})
		})
		Context("GetLimit", func() {
			When("if limit not set", func() {
				It("returns 100", func() {
					params.SetLimit(0)
					Expect(params.GetLimit()).To(Equal(int32(100)))
				})
			})
			It("returns limit", func() {
				Expect(params.GetLimit()).To(Equal(int32(50)))
			})
		})
		Context("SetLimit", func() {
			It("can set offset", func() {
				params.SetLimit(100)
				Expect(params.GetLimit()).To(Equal(int32(100)))
			})
		})
	})
	Context("SearchType.Validate", func() {
		When("SearchType is not `AND` or `OR`", func() {
			It("returns false", func() {
				Expect(api.SearchType("and").Validate()).To(BeFalse())
				Expect(api.SearchType("or").Validate()).To(BeFalse())
				Expect(api.SearchType("NOT").Validate()).To(BeFalse())
			})
		})
		When("SearchType is `AND` or `OR`", func() {
			It("returns true", func() {
				Expect(api.SearchType("AND").Validate()).To(BeTrue())
				Expect(api.SearchType("OR").Validate()).To(BeTrue())
			})
		})
	})
	Context("SearchOrder.Validate", func() {
		When("SearchOrder is not `ASC` or `DESC`", func() {
			It("returns false", func() {
				Expect(api.SearchOrder("asc").Validate()).To(BeFalse())
				Expect(api.SearchOrder("desc").Validate()).To(BeFalse())
				Expect(api.SearchOrder("hoge").Validate()).To(BeFalse())
			})
		})
		When("SearchOrder is `ASC` or `DESC`", func() {
			It("returns true", func() {
				Expect(api.SearchOrder("ASC").Validate()).To(BeTrue())
				Expect(api.SearchOrder("DESC").Validate()).To(BeTrue())
			})
		})
	})
	Context("SearchOffset.Validate", func() {
		When("SearchOffset is not in range 0 to 10000000", func() {
			It("returns false", func() {
				Expect(api.SearchOffset(-2).Validate()).To(BeFalse())
				Expect(api.SearchOffset(-1).Validate()).To(BeFalse())
				Expect(api.SearchOffset(10000001).Validate()).To(BeFalse())
				Expect(api.SearchOffset(10000002).Validate()).To(BeFalse())
			})
		})
		When("SearchOffset is in range 0 to 10000000", func() {
			It("returns true", func() {
				Expect(api.SearchOffset(0).Validate()).To(BeTrue())
				Expect(api.SearchOffset(1).Validate()).To(BeTrue())
				Expect(api.SearchOffset(10000000).Validate()).To(BeTrue())
				Expect(api.SearchOffset(9999999).Validate()).To(BeTrue())
			})
		})
	})
	Context("SearchLimit.Validate", func() {
		When("SearchLimit is not in range 1 to 10000", func() {
			It("returns false", func() {
				Expect(api.SearchLimit(-1).Validate()).To(BeFalse())
				Expect(api.SearchLimit(0).Validate()).To(BeFalse())
				Expect(api.SearchLimit(10001).Validate()).To(BeFalse())
				Expect(api.SearchLimit(10002).Validate()).To(BeFalse())
			})
		})
		When("SearchLimit is in range 1 to 10000", func() {
			It("returns true", func() {
				Expect(api.SearchLimit(1).Validate()).To(BeTrue())
				Expect(api.SearchLimit(2).Validate()).To(BeTrue())
				Expect(api.SearchLimit(10000).Validate()).To(BeTrue())
				Expect(api.SearchLimit(9999).Validate()).To(BeTrue())
			})
		})
	})
	Context("KeywordsString.Validate", func() {
		When("KeywordsString length grather than 255", func() {
			It("returns false", func() {
				s := ""
				for i := 0; i < 256; i++ {
					s += "A"
				}
				Expect(api.KeywordsString{s}.Validate()).To(BeFalse())
				Expect(api.KeywordsString{"", s}.Validate()).To(BeFalse())
				Expect(api.KeywordsString{"", "b", s}.Validate()).To(BeFalse())
			})
		})
		When("SearchLimit is not include length grather than 255", func() {
			It("returns true", func() {
				s := ""
				for i := 0; i < 255; i++ {
					s += "A"
				}
				Expect(api.KeywordsString{"", "b", s}.Validate()).To(BeTrue())
				Expect(api.KeywordsString{""}.Validate()).To(BeTrue())
			})
		})
	})
	Context("KeywordsID.Validate", func() {
		When("KeywordsID is include minus value", func() {
			It("returns false", func() {
				Expect(api.KeywordsID{-1}.Validate()).To(BeFalse())
				Expect(api.KeywordsID{1, -1}.Validate()).To(BeFalse())
			})
		})
		When("KeywordsID is not include minus value", func() {
			It("returns false", func() {
				Expect(api.KeywordsID{0}.Validate()).To(BeTrue())
				Expect(api.KeywordsID{0, 1}.Validate()).To(BeTrue())
				Expect(api.KeywordsID{0, 1, 2}.Validate()).To(BeTrue())
			})
		})
	})
	Context("KeywordsBoolean.EncodeValues", func() {
		var (
			value    url.Values
			err      error
			keywords struct {
				Tests api.KeywordsBoolean `url:"tests"`
			}
		)
		BeforeEach(func() {
			keywords.Tests = api.KeywordsBoolean{
				types.Enabled, types.Disabled,
			}
			value, err = query.Values(keywords)
		})
		It("can encode url.Vale", func() {
			Expect(err).To(Succeed())
			v := url.Values{
				"tests": []string{"1", "0"},
			}
			Expect(value).To(Equal(v))
		})
	})
	Context("KeywordsState.EncodeValues", func() {
		var (
			value    url.Values
			err      error
			keywords struct {
				Tests api.KeywordsState `url:"tests"`
			}
		)
		BeforeEach(func() {
			keywords.Tests = api.KeywordsState{
				types.StateBeforeStart, types.StateRunning,
			}
			value, err = query.Values(keywords)
		})
		It("can encode url.Vale", func() {
			Expect(err).To(Succeed())
			v := url.Values{
				"tests": []string{"1", "2"},
			}
			Expect(value).To(Equal(v))
		})
	})
	Context("KeywordsState.KeywordsFavorite", func() {
		var (
			value    url.Values
			err      error
			keywords struct {
				Tests api.KeywordsFavorite `url:"tests"`
			}
		)
		BeforeEach(func() {
			keywords.Tests = api.KeywordsFavorite{
				types.FavoriteHighPriority, types.FavoriteLowPriority,
			}
			value, err = query.Values(keywords)
		})
		It("can encode url.Vale", func() {
			Expect(err).To(Succeed())
			v := url.Values{
				"tests": []string{"1", "2"},
			}
			Expect(value).To(Equal(v))
		})
	})
})
