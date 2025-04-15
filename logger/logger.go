package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

func ExtractLogger(ctx context.Context, level int, format string, args ...interface{}) {
	logger, ok := ctx.Value(TraceCtxKey).(*zap.SugaredLogger)
	if !ok {
		logger = SugarLogger()
	}

	switch level {
	case DEBUG:
		logger.Debugf(format, args...)
	case INFO:
		logger.Infof(format, args...)
	case WARN:
		logger.Warnf(format, args...)
	case ERROR:
		logger.Errorf(format, args...)
	default:
		logger.Fatalf(format, args...)
	}
}
