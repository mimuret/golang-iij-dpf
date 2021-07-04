package api

import "fmt"

type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type StdLogger struct {
	LogLevel int
}

func (s *StdLogger) Tracef(format string, args ...interface{}) {
	if s.LogLevel <= 0 {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
func (s *StdLogger) Debugf(format string, args ...interface{}) {
	if s.LogLevel <= 1 {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
func (s *StdLogger) Infof(format string, args ...interface{}) {
	if s.LogLevel <= 2 {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
func (s *StdLogger) Errorf(format string, args ...interface{}) {
	if s.LogLevel <= 4 {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}
