package types_test

import (
	"encoding/json"
	"net"

	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("types.net", func() {
	Context("ParseIPNet", func() {
		var (
			err   error
			ipnet *types.IPNet
		)
		When("failed to parse net.IPNet", func() {
			BeforeEach(func() {
				_, err = types.ParseIPNet("//")
			})
			It("returns err", func() {
				Expect(err).To(HaveOccurred())
			})
		})
		When("parse successfull", func() {
			BeforeEach(func() {
				ipnet, err = types.ParseIPNet("192.168.0.0/24")
			})
			It("returns count value", func() {
				Expect(err).To(Succeed())
				Expect(ipnet).To(Equal(&types.IPNet{
					net.IPNet{
						IP:   net.ParseIP("192.168.0.0").To4(),
						Mask: net.IPv4Mask(255, 255, 255, 0),
					},
				}))
			})
		})
	})
	Context("IPNet", func() {
		var (
			err   error
			bs    []byte
			ipnet *types.IPNet
		)
		BeforeEach(func() {
			ipnet = &types.IPNet{}
		})
		Context("MarshalJSON", func() {
			When("valid value", func() {
				BeforeEach(func() {
					ipnet, err = types.ParseIPNet("192.168.0.0/24")
					Expect(err).To(Succeed())
					bs, err = json.Marshal(ipnet)
					Expect(err).To(Succeed())
				})
				It("returns count value", func() {
					Expect(string(bs)).To(Equal(`"192.168.0.0/24"`))
				})
			})
			When("valid empty", func() {
				BeforeEach(func() {
					bs, err = json.Marshal(ipnet)
					Expect(err).To(Succeed())
				})
				It("returns count value", func() {
					Expect(string(bs)).To(Equal(`"\u003cnil\u003e"`))
				})
			})
		})
		Context("UnmarshalJSON", func() {
			When("valid format", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`"192.168.0.0/24"`), ipnet)
					Expect(err).To(Succeed())
				})
				It("returns count value", func() {
					Expect(ipnet).To(Equal(&types.IPNet{
						net.IPNet{
							IP:   net.ParseIP("192.168.0.0").To4(),
							Mask: net.IPv4Mask(255, 255, 255, 0),
						},
					}))
				})
			})
			When("input value is struct", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`{"name": "hoge"}`), ipnet)
				})
				It("returns err", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("invalid format", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`"192.168.0.0/24/24"`), ipnet)
				})
				It("returns err", func() {
					Expect(err).To(HaveOccurred())
				})
			})
			When("\u003cnil\u003e", func() {
				BeforeEach(func() {
					err = json.Unmarshal([]byte(`"\u003cnil\u003e"`), ipnet)
				})
				It("returns err", func() {
					Expect(err).To(Succeed())
					Expect(*ipnet).To(Equal(types.IPNet{}))
				})
			})
		})
		Context("DeepCopy", func() {
			BeforeEach(func() {
				ipnet, err = types.ParseIPNet("2001:db8::1/64")
				Expect(err).To(Succeed())
			})
			It("returns copy object", func() {
				copy := ipnet.DeepCopy()
				Expect(copy).To(Equal(ipnet))
				Expect(copy).ShouldNot(BeIdenticalTo(ipnet))
				Expect(copy.IPNet.IP).ShouldNot(BeIdenticalTo(ipnet.IPNet.IP))
				Expect(copy.IPNet.Mask).ShouldNot(BeIdenticalTo(ipnet.IPNet.Mask))
			})
		})
	})
})
