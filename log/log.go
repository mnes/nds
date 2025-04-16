package log

import (
	"context"

	"github.com/qedus/nds/v2/internal"
)

func Debugf(ctx context.Context, format string, args ...any) {
	internal.Logger.Debugf(ctx, format, args...)
}

func Infof(ctx context.Context, format string, args ...any) {
	internal.Logger.Infof(ctx, format, args...)
}

func Warningf(ctx context.Context, format string, args ...any) {
	internal.Logger.Warnf(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...any) {
	internal.Logger.Errorf(ctx, format, args...)
}

func SetLogger(logger internal.Logging) {
	internal.Logger = logger
}
