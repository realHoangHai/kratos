package logger

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/realHoangHai/kratos/internal/conf"
	"github.com/realHoangHai/kratos/pkg/logger/logrus"
	"github.com/realHoangHai/kratos/pkg/logger/zap"
)

type Type string

const (
	Std    Type = "std"
	Logrus Type = "logrus"
	Zap    Type = "zap"
)

// NewLogger create a new logger
func NewLogger(cfg *conf.Logger) log.Logger {
	if cfg == nil {
		return NewStdLogger()
	}

	switch Type(cfg.Type) {
	default:
		fallthrough
	case Std:
		return NewStdLogger()
	case Zap:
		return zap.NewLogger(cfg)
	case Logrus:
		return logrus.NewLogger(cfg)
	}
}

// NewLoggerProvider create new logger provider
func NewLoggerProvider(cfg *conf.Logger, serviceInfo *conf.ServiceInfo) log.Logger {
	l := NewLogger(cfg)

	return log.With(
		l,
		"service.id", serviceInfo.ID,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// NewStdLogger create a new logger - Kratos built-in, console output
func NewStdLogger() log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return l
}
