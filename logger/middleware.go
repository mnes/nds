package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var TraceCtxKey = &contextKey{"trace"}

type contextKey struct {
	name string
}

type Trace struct {
	TraceID string
	SpanID  string
	Sampled bool
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		r := ctx.Request()
		header := r.Header.Get("X-Cloud-Trace-Context")
		if len(header) > 0 {
			traceID, spanID, sampled := deconstructXCloudTraceContext(header)

			t := &Trace{
				TraceID: traceID,
				SpanID:  spanID,
				Sampled: sampled,
			}
			c := loggerWithContext(r.Context(), t)
			r = r.WithContext(c)
			ctx.SetRequest(r)
		}

		return next(ctx)
	}

}

var reCloudTraceContext = regexp.MustCompile(
	// Matches on "TRACE_ID"
	`([a-f\d]+)?` +
		// Matches on "/SPAN_ID"
		`(?:/([a-f\d]+))?` +
		// Matches on ";0=TRACE_TRUE"
		`(?:;o=(\d))?`)

func deconstructXCloudTraceContext(s string) (traceID, spanID string, traceSampled bool) {
	matches := reCloudTraceContext.FindStringSubmatch(s)

	traceID, spanID, traceSampled = matches[1], matches[2], matches[3] == "1"

	if spanID == "0" {
		spanID = ""
	}

	return
}

func loggerWithContext(ctx context.Context, trace *Trace) context.Context {
	var config zap.Config
	if os.Getenv("APP_ENV") == "local" {
		// running locally we will use a human-readable output
		config = zapdriver.NewDevelopmentConfig()
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// create our uber zap configuration
		config = zapdriver.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	// creates our logger instance
	clientLogger, err := config.Build()
	if err != nil {
		log.Fatalf("zap.config.Build(): %v", err)
	}

	fields := zapdriver.TraceContext(trace.TraceID, trace.SpanID, trace.Sampled, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	fields = append(fields, zap.String("logName", fmt.Sprintf("projects/%s/logs/%s", os.Getenv("GOOGLE_CLOUD_PROJECT"), os.Getenv("GAE_SERVICE"))))
	setFields := clientLogger.With(fields...)

	return context.WithValue(ctx, TraceCtxKey, setFields.Sugar())
}

func SugarLogger() *zap.SugaredLogger {
	var config zap.Config
	if os.Getenv("APP_ENV") == "local" {
		// running locally we will use a human-readable output
		config = zapdriver.NewDevelopmentConfig()
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		// create our uber zap configuration
		config = zapdriver.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	// creates our logger instance
	clientLogger, err := config.Build()
	if err != nil {
		log.Fatalf("zap.config.Build(): %v", err)
	}

	trace := &Trace{
		TraceID: os.Getenv("GAE_INSTANCE"),
		SpanID:  fmt.Sprintf("%d", time.Now().UnixNano()),
		Sampled: true,
	}
	fields := zapdriver.TraceContext(trace.TraceID, trace.SpanID, trace.Sampled, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	fields = append(fields, zap.String("logName", fmt.Sprintf("projects/%s/logs/%s", os.Getenv("GOOGLE_CLOUD_PROJECT"), os.Getenv("GAE_SERVICE"))))
	setFields := clientLogger.With(fields...)

	return setFields.Sugar()
}
