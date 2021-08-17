package api_test

import (
	"bytes"
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
)

var _ = Describe("logger", func() {
	Context("StdLogger", func() {
		var (
			logger *api.StdLogger
			buf    *bytes.Buffer
		)
		BeforeEach(func() {
			buf = bytes.NewBuffer(nil)
			logger = api.NewStdLogger(buf, "test", 0, 0)
		})
		Context("NewStdLogger", func() {
			When("default value", func() {
				BeforeEach(func() {
					logger = api.NewStdLogger(nil, "", 0, 0)
				})
				It("io.Writer is os.Stderr", func() {
					Expect(logger.Logger).To(Equal(log.New(os.Stderr, "dpf-client", 0)))
				})
			})
		})
		Context("Tracef", func() {
			When("level 0", func() {
				It("is out message", func() {
					logger.LogLevel = 0
					logger.Tracef("trace message")
					Expect(string(buf.String())).To(MatchRegexp("trace message"))
				})
			})
			When("level >= 1", func() {
				It("is not out message", func() {
					logger.LogLevel = 1
					logger.Tracef("trace message")
					Expect(string(buf.String())).NotTo(MatchRegexp("trace message"))
				})
				It("is not out message", func() {
					logger.LogLevel = 2
					logger.Tracef("trace message")
					Expect(string(buf.String())).NotTo(MatchRegexp("trace message"))
				})
			})
		})
		Context("Debugf", func() {
			When("level <= 1", func() {
				It("is out message (level 0)", func() {
					logger.LogLevel = 0
					logger.Debugf("debug message")
					Expect(string(buf.String())).To(MatchRegexp("debug message"))
				})
				It("is out message (level 1)", func() {
					logger.LogLevel = 1
					logger.Debugf("debug message")
					Expect(string(buf.String())).To(MatchRegexp("debug message"))
				})
			})
			When("level > 1", func() {
				It("is not out message", func() {
					logger.LogLevel = 2
					logger.Debugf("debug message")
					Expect(string(buf.String())).NotTo(MatchRegexp("debug message"))
				})
				It("is not out message", func() {
					logger.LogLevel = 3
					logger.Debugf("debug message")
					Expect(string(buf.String())).NotTo(MatchRegexp("debug message"))
				})
			})
		})
		Context("Infof", func() {
			When("level <= 2", func() {
				It("is out message (level 0)", func() {
					logger.LogLevel = 0
					logger.Infof("info message")
					Expect(string(buf.String())).To(MatchRegexp("info message"))
				})
				It("is out message (level 1)", func() {
					logger.LogLevel = 1
					logger.Infof("info message")
					Expect(string(buf.String())).To(MatchRegexp("info message"))
				})
				It("is out message (level 2)", func() {
					logger.LogLevel = 2
					logger.Infof("info message")
					Expect(string(buf.String())).To(MatchRegexp("info message"))
				})
			})
			When("level > 2", func() {
				It("is not out message", func() {
					logger.LogLevel = 3
					logger.Infof("info message")
					Expect(string(buf.String())).NotTo(MatchRegexp("info message"))
				})
			})
		})
		Context("Errorf", func() {
			When("level <= 4", func() {
				It("is out message (level 0)", func() {
					logger.LogLevel = 0
					logger.Errorf("error message")
					Expect(string(buf.String())).To(MatchRegexp("error message"))
				})
				It("is out message (level 1)", func() {
					logger.LogLevel = 1
					logger.Errorf("error message")
					Expect(string(buf.String())).To(MatchRegexp("error message"))
				})
				It("is out message (level 2)", func() {
					logger.LogLevel = 2
					logger.Errorf("error message")
					Expect(string(buf.String())).To(MatchRegexp("error message"))
				})
				It("is out message (level 3)", func() {
					logger.LogLevel = 3
					logger.Errorf("error message")
					Expect(string(buf.String())).To(MatchRegexp("error message"))
				})
				It("is out message (level 4)", func() {
					logger.LogLevel = 4
					logger.Errorf("error message")
					Expect(string(buf.String())).To(MatchRegexp("error message"))
				})
			})
		})
	})
})
