package gormlog

import (
	"context"
	"time"

	ivilog "github.com/ividernvi/ivilog"
	glog "gorm.io/gorm/logger"
)

type GormLogger struct {
	ivilog *ivilog.Logger
}

func New() *GormLogger {
	return &GormLogger{
		ivilog: ivilog.New(),
	}
}

func (gl *GormLogger) LogMode(level glog.LogLevel) glog.Interface {
	switch level {
	case glog.Info:
		gl.ivilog.SetLevel(ivilog.InfoLevel)
	case glog.Warn:
		gl.ivilog.SetLevel(ivilog.WarnLevel)
	case glog.Error:
		gl.ivilog.SetLevel(ivilog.ErrorLevel)
	case glog.Silent:
		gl.ivilog.SetLevel(ivilog.DebugLevel)
	}
	return gl
}

func (gl *GormLogger) Info(ctx context.Context, text string, args ...interface{}) {
	gl.ivilog.Infof(text, args...)
}

func (gl *GormLogger) Warn(ctx context.Context, text string, args ...interface{}) {
	gl.ivilog.Warnf(text, args...)
}

func (gl *GormLogger) Error(ctx context.Context, text string, args ...interface{}) {
	gl.ivilog.Errorf(text, args...)
}

func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		gl.ivilog.WithFields(ivilog.Fields{
			"elapsed": elapsed,
			"rows":    rows,
			"sql":     sql,
		}).WithError(err).Error("Trace Error")
	} else {
		gl.ivilog.WithFields(ivilog.Fields{
			"elapsed": elapsed,
			"rows":    rows,
			"sql":     sql,
		}).Info("Trace")
	}
}
