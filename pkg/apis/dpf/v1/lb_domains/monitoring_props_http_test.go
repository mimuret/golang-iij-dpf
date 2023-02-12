package lb_domains_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/lb_domains"
)

var _ = Describe("MonitoringPorpsHTTP", func() {
	var (
		bs         []byte
		c          lb_domains.MonitoringPorpsHTTP
		err        error
		s1, s2, s3 lb_domains.MonitoringPorpsHTTP
	)
	BeforeEach(func() {
		s1 = lb_domains.MonitoringPorpsHTTP{
			MonitoringPorpsCommon: lb_domains.MonitoringPorpsCommon{
				Location: lb_domains.MonitoringPropsLocationAll,
				Interval: 30,
				Holdtime: 0,
				Timeout:  1,
			},
			Port:          443,
			TLSSNI:        "example.jp",
			ResponseMatch: "successful",
			HTTPS:         true,
			Path:          "/ready",
			StatusCode:    []string{"200"},
		}
		s2 = lb_domains.MonitoringPorpsHTTP{
			MonitoringPorpsCommon: lb_domains.MonitoringPorpsCommon{
				Location: lb_domains.MonitoringPropsLocationJP,
				Interval: 600,
				Holdtime: 3600,
				Timeout:  30,
			},
			Port:          80,
			TLSSNI:        "",
			ResponseMatch: "",
			HTTPS:         false,
			Path:          "",
			StatusCode:    []string{},
		}
		s3 = lb_domains.MonitoringPorpsHTTP{
			Port: 53,
		}
	})
	Context("Read", func() {
		Context("s1", func() {
			BeforeEach(func() {
				err = api.UnmarshalRead(json.RawMessage(`{
					"location": "all",
					"interval": 30,
					"holdtime": 0,
					"timeout": 1,
					"port": 443,
					"https": true,
					"tls_sni": "example.jp",
					"response_match": "successful",
					"path": "/ready",
					"status_codes": ["200"]
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
					"location": "jp",
					"interval": 600,
					"holdtime": 3600,
					"timeout": 30,
					"port": 80,
					"https": false,
					"tls_sni": "",
					"response_match": "",
					"path": "",
					"status_codes": []
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
					"location": "all",
					"interval": 30,
					"holdtime": 0,
					"timeout": 1,
					"port": 443,
					"https": true,
					"tls_sni": "example.jp",
					"response_match": "successful",
					"path": "/ready",
					"status_codes": ["200"]
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
					"location": "",
					"interval": 0,
					"holdtime": 0,
					"timeout": 0,
					"port": 53,
					"https": false,
					"tls_sni": "",
					"response_match": "",
					"path": "",
					"status_codes": null
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
					"location": "all",
					"interval": 30,
					"holdtime": 0,
					"timeout": 1,
					"port": 443,
					"https": true,
					"tls_sni": "example.jp",
					"response_match": "successful",
					"path": "/ready",
					"status_codes": ["200"]
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
					"holdtime": 0,
					"port": 53,
					"https": false
				}`))
			})
		})
	})
})
