package types_test

import (
	"encoding/json"
	"time"

	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("types.time", func() {
	Context("ParseTime", func() {
		var (
			err error
			t   types.Time
		)
		When("failed to parse time.Parse", func() {
			BeforeEach(func() {
				_, err = types.ParseTime(" ", "//")
			})
			It("returns err", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("parse successfull", func() {
			BeforeEach(func() {
				t, err = types.ParseTime(time.RFC3339Nano, "2021-07-17T17:27:05.999999999+09:00")
				Expect(err).To(Succeed())
			})
			It("returns count value", func() {
				Expect(t.Year()).To(Equal(2021))
				Expect(t.Month()).To(Equal(time.July))
				Expect(t.Day()).To(Equal(17))
				Expect(t.Hour()).To(Equal(17))
				Expect(t.Minute()).To(Equal(27))
			})
		})
	})
	Context("Time", func() {
		var (
			bs  []byte
			err error
			t   types.Time
		)
		Context("MarshalJSON", func() {
			When("valid value", func() {
				BeforeEach(func() {
					t, err = types.ParseTime(time.RFC3339Nano, "2021-07-17T17:27:05.01+09:00")
					Expect(err).To(Succeed())
					bs, err = json.Marshal(t)
					Expect(err).To(Succeed())
				})
				It("return string", func() {
					Expect(string(bs)).To(Equal(`"2021-07-17T17:27:05.01+09:00"`))
				})
			})
		})
		Context("UnmarshalJSON", func() {
			When("invalid format", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`"p[rke[prkew[rkwe[rkw["`), &t)
				})
				It("returns count value", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("empty value", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`""`), &t)
				})
				It("returns count value", func() {
					Expect(err).To(Succeed())
					Expect(t).To(Equal(types.Time{}))
				})
			})
			When("valid format", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`"2021-07-17T17:27:05.0+09:00"`), &t)
					Expect(err).To(Succeed())
				})
				It("returns count value", func() {
					te, err := types.ParseTime(time.RFC3339Nano, "2021-07-17T17:27:05.0+09:00")
					Expect(err).To(Succeed())
					Expect(t).To(Equal(te))
				})
			})
		})
		Context("DeepCopyInto", func() {
			BeforeEach(func() {
				err = json.Unmarshal([]byte(`"2021-07-17T17:27:05.0+09:00"`), &t)
				Expect(err).To(Succeed())
			})
			It("returns copy object", func() {
				copy := types.Time{}
				copy.DeepCopyInto(&t)
				Expect(copy).To(Equal(t))
				Expect(&copy).ShouldNot(BeIdenticalTo(&t))
			})
		})
	})
})
