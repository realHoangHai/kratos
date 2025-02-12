package logrus

import (
	"github.com/sirupsen/logrus"

	kratoslogrus "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
)

// NewLogger create new logger - Logrus
func NewLogger(cfg *conf.Logger) log.Logger {
	loggerLevel, err := logrus.ParseLevel(cfg.Logrus.Level)
	if err != nil {
		loggerLevel = logrus.InfoLevel
	}

	var loggerFormatter logrus.Formatter
	switch cfg.Logrus.Formatter {
	default:
		fallthrough
	case "text":
		loggerFormatter = &logrus.TextFormatter{
			DisableColors:    cfg.Logrus.DisableColors,
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	case "json":
		loggerFormatter = &logrus.JSONFormatter{
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	}

	logger := logrus.New()
	logger.Level = loggerLevel
	logger.Formatter = loggerFormatter

	wrapped := kratoslogrus.NewLogger(logger)
	return wrapped
}
