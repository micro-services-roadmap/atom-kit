package rich

import (
	"context"
	"fmt"
	"time"

	util "github.com/micro-services-roadmap/atom-kit/logx/util"
	"github.com/micro-services-roadmap/atom-kit/tracing"
	"github.com/micro-services-roadmap/atom-kit/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// assert that richLogger implements the Logger interface
var _ Logger = (*richLogger)(nil)

const callerDepth = 4

func NewRichLogger() Logger {
	return &richLogger{
		logger: zapx.NewZap(),
	}
}

type richLogger struct {
	ctx        context.Context
	callerSkip int
	fields     []zap.Field
	logger     *zap.Logger
}

// WithLogger returns a Logger with given caller skip.
func WithLogger(logger *zap.Logger) Logger {
	return &richLogger{
		logger: logger,
	}
}

// WithCallerSkip returns a Logger with given caller skip.
func WithCallerSkip(skip int) Logger {
	if skip <= 0 {
		return new(richLogger)
	}

	return &richLogger{
		callerSkip: skip,
	}
}

// WithContext sets ctx to log, for keeping tracing information.
func WithContext(ctx context.Context) Logger {
	return &richLogger{
		ctx: ctx,
	}
}

// WithDuration returns a Logger with given duration.
func WithDuration(d time.Duration) Logger {
	field := zap.Field{
		Key:    durationKey,
		Type:   zapcore.StringType,
		String: util.ReprOfDuration(d),
	}

	return &richLogger{
		fields: []zap.Field{field},
	}
}

func (l *richLogger) Debuga(v ...any) {
	if util.ShallLog(DebugLevel) {
		l.debug(fmt.Sprint(v...))
	}
}

func (l *richLogger) Debugf(format string, v ...any) {
	if util.ShallLog(DebugLevel) {
		l.debug(fmt.Sprintf(format, v...))
	}
}

func (l *richLogger) Debugv(msg string) {
	if util.ShallLog(DebugLevel) {
		l.debug(msg)
	}
}

func (l *richLogger) Debug(msg string, fields ...zap.Field) {
	if util.ShallLog(DebugLevel) {
		l.debug(msg, fields...)
	}
}

func (l *richLogger) Infoa(v ...any) {
	if util.ShallLog(InfoLevel) {
		l.info(fmt.Sprint(v...))
	}
}

func (l *richLogger) Infof(format string, v ...any) {
	if util.ShallLog(InfoLevel) {
		l.info(fmt.Sprintf(format, v...))
	}
}

func (l *richLogger) Infov(v string) {
	if util.ShallLog(InfoLevel) {
		l.info(v)
	}
}

func (l *richLogger) Info(msg string, fields ...zap.Field) {
	if util.ShallLog(InfoLevel) {
		l.info(msg, fields...)
	}
}

func (l *richLogger) Warna(v ...any) {
	if util.ShallLog(WarnLevel) {
		l.warn(fmt.Sprint(v...))
	}
}

func (l *richLogger) Warnf(format string, v ...any) {
	if util.ShallLog(WarnLevel) {
		l.warn(fmt.Sprintf(format, v...))
	}
}

func (l *richLogger) Warnv(v string) {
	if util.ShallLog(WarnLevel) {
		l.warn(v)
	}
}

func (l *richLogger) Warn(msg string, fields ...zap.Field) {
	if util.ShallLog(WarnLevel) {
		l.warn(msg, fields...)
	}
}

func (l *richLogger) Errora(v ...any) {
	if util.ShallLog(ErrorLevel) {
		l.err(fmt.Sprint(v...))
	}
}

func (l *richLogger) Errorf(format string, v ...any) {
	if util.ShallLog(ErrorLevel) {
		l.err(fmt.Sprintf(format, v...))
	}
}

func (l *richLogger) Errorv(v string) {
	if util.ShallLog(ErrorLevel) {
		l.err(v)
	}
}

func (l *richLogger) Error(msg string, fields ...zap.Field) {
	if util.ShallLog(ErrorLevel) {
		l.err(msg, fields...)
	}
}

func (l *richLogger) WithCallerSkip(skip int) Logger {
	if skip <= 0 {
		return l
	}

	l.callerSkip = skip
	return l
}

func (l *richLogger) WithContext(ctx context.Context) Logger {
	l.ctx = ctx
	return l
}

func (l *richLogger) WithDuration(d time.Duration) Logger {
	field := zap.Field{
		Key:    durationKey,
		Type:   zapcore.StringType,
		String: util.ReprOfDuration(d),
	}

	l.fields = append(l.fields, field)
	return l
}

func (l *richLogger) WithFields(fields ...zap.Field) Logger {
	l.fields = append(l.fields, fields...)
	return l
}

func (l *richLogger) buildFields(fields ...zap.Field) []zap.Field {
	fields = append(l.fields, fields...)
	fields = append(fields, zap.Field{
		Key:    callerKey,
		Type:   zapcore.StringType,
		String: util.GetCaller(callerDepth + l.callerSkip),
	})

	if l.ctx == nil {
		return fields
	}

	traceID := tracing.TraceIDFromContext(l.ctx)
	if len(traceID) > 0 {
		field := zap.Field{
			Key:    traceKey,
			Type:   zapcore.StringType,
			String: traceID,
		}
		fields = append(fields, field)
	}

	spanID := tracing.SpanIDFromContext(l.ctx)
	if len(spanID) > 0 {
		field := zap.Field{
			Key:    spanKey,
			Type:   zapcore.StringType,
			String: traceID,
		}
		fields = append(fields, field)
	}

	val := l.ctx.Value(fieldsContextKey)
	if val != nil {
		if arr, ok := val.([]zap.Field); ok {
			fields = append(fields, arr...)
		}
	}

	return fields
}

func (l *richLogger) debug(msg string, fields ...zap.Field) {
	if util.ShallLog(DebugLevel) {
		l.logger.Debug(msg, l.buildFields(fields...)...)
	}
}

func (l *richLogger) info(v string, fields ...zap.Field) {
	if util.ShallLog(InfoLevel) {
		l.logger.Info(v, l.buildFields(fields...)...)
	}
}

func (l *richLogger) warn(v string, fields ...zap.Field) {
	if util.ShallLog(WarnLevel) {
		l.logger.Warn(v, l.buildFields(fields...)...)
	}
}

func (l *richLogger) err(v string, fields ...zap.Field) {
	if util.ShallLog(ErrorLevel) {
		l.logger.Error(v, l.buildFields(fields...)...)
	}
}
