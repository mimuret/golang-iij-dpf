package lb_domains_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
)

var _ = Describe("MonitoringPorpsStatic", func() {
	var (
		bs         []byte
		c          lb_domains.MonitoringPorpsStatic
		err        error
		s1, s2, s3 lb_domains.MonitoringPorpsStatic
	)
	BeforeEach(func() {
		s1 = lb_domains.MonitoringPorpsStatic{
			Result: lb_domains.MonitoringPorpsStaticStatusUp,
		}
		s2 = lb_domains.MonitoringPorpsStatic{
			Result: lb_domains.MonitoringPorpsStaticStatusDown,
		}
		s3 = lb_domains.MonitoringPorpsStatic{
			Result: lb_domains.MonitoringPorpsStaticStatusUnkown,
		}
	})
	Context("Read", func() {
		Context("s1", func() {
			BeforeEach(func() {
				err = api.UnmarshalRead(json.RawMessage(`{
					"result": "up"
				}`), &c)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(c).To(Equal(s1))
			})
		})
		Context("s2", func() {
			BeforeEach(func() {
				err = api.UnmarshalRead(json.RawMessage(`{
					"result": "down"
				}`), &c)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(c).To(Equal(s2))
			})
		})
	})
	Context("Update", func() {
		Context("s1", func() {
			BeforeEach(func() {
				bs, err = api.MarshalUpdate(s1)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(MatchJSON(`{
					"result": "up"
				}`))
			})
		})
		Context("s3", func() {
			BeforeEach(func() {
				bs, err = api.MarshalUpdate(s3)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(MatchJSON(`{
					"result": "unknown"
				}`))
			})
		})
	})
	Context("Create", func() {
		Context("s1", func() {
			BeforeEach(func() {
				bs, err = api.MarshalCreate(s1)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(MatchJSON(`{
					"result": "up"
				}`))
			})
		})
		Context("s3", func() {
			BeforeEach(func() {
				bs, err = api.MarshalCreate(s3)
			})
			It("succeed", func() {
				Expect(err).To(Succeed())
				Expect(bs).To(MatchJSON(`{
					"result": "unknown"
				}`))
			})
		})
	})
})
