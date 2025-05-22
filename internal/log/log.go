package log

import (
	"context"
	"fmt"
	"log/slog"
)

type Logging interface {
	Debugf(ctx context.Context, format string, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Warnf(ctx context.Context, format string, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
}

// default logger
type logger struct{}

func (l *logger) Debugf(ctx context.Context, format string, args ...any) {
	slog.DebugContext(ctx, fmt.Sprintf(format, args...))
}

func (l *logger) Infof(ctx context.Context, format string, args ...any) {
	slog.InfoContext(ctx, fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(ctx context.Context, format string, args ...any) {
	slog.WarnContext(ctx, fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(ctx context.Context, format string, args ...any) {
	slog.ErrorContext(ctx, fmt.Sprintf(format, args...))
}

var Logger Logging = &logger{}

func Debugf(ctx context.Context, format string, args ...any) {
	Logger.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	Logger.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...any) {
	Logger.Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	Logger.Errorf(ctx, format, args...)
}
