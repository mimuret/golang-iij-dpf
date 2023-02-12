package types_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("types.NullablePositiveInt32", func() {
	Context("UnmarshalJSON", func() {
		var (
			err error
			p32 types.NullablePositiveInt32
		)
		When("failed to parse int32", func() {
			BeforeEach(func() {
				err = p32.UnmarshalJSON([]byte("/"))
			})
			It("returns err", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("args is null", func() {
			BeforeEach(func() {
				err = p32.UnmarshalJSON([]byte("null"))
			})
			It("returns 0 value", func() {
				Expect(err).To(Succeed())
				Expect(p32).To(Equal(types.NullablePositiveInt32(0)))
			})
		})
		When("parse successfull", func() {
			BeforeEach(func() {
				err = p32.UnmarshalJSON([]byte("1"))
			})
			It("returns value", func() {
				Expect(err).To(Succeed())
				Expect(p32).To(Equal(types.NullablePositiveInt32(1)))
			})
		})
	})
	Context("MarshalJSON", func() {
		var (
			err error
			bs  []byte
		)
		When("value is zero", func() {
			BeforeEach(func() {
				bs, err = types.NullablePositiveInt32(0).MarshalJSON()
			})
			It("returns err", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(Equal([]byte("null")))
			})
		})
		When("value is not null", func() {
			BeforeEach(func() {
				bs, err = types.NullablePositiveInt32(1).MarshalJSON()
			})
			It("returns 0 value", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(Equal([]byte("1")))
			})
		})
	})
})

var _ = Describe("types.NullablePositiveInt64", func() {
	Context("UnmarshalJSON", func() {
		var (
			err error
			p64 types.NullablePositiveInt64
		)
		When("failed to parse int64", func() {
			BeforeEach(func() {
				err = p64.UnmarshalJSON([]byte("/"))
			})
			It("returns err", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("args is null", func() {
			BeforeEach(func() {
				err = p64.UnmarshalJSON([]byte("null"))
			})
			It("returns 0 value", func() {
				Expect(err).To(Succeed())
				Expect(p64).To(Equal(types.NullablePositiveInt64(0)))
			})
		})
		When("parse successfull", func() {
			BeforeEach(func() {
				err = p64.UnmarshalJSON([]byte("1"))
			})
			It("returns value", func() {
				Expect(err).To(Succeed())
				Expect(p64).To(Equal(types.NullablePositiveInt64(1)))
			})
		})
	})
	Context("MarshalJSON", func() {
		var (
			err error
			bs  []byte
		)
		When("value is zero", func() {
			BeforeEach(func() {
				bs, err = types.NullablePositiveInt64(0).MarshalJSON()
			})
			It("returns err", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(Equal([]byte("null")))
			})
		})
		When("value is not null", func() {
			BeforeEach(func() {
				bs, err = types.NullablePositiveInt64(1).MarshalJSON()
			})
			It("returns 0 value", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(Equal([]byte("1")))
			})
		})
	})
})
