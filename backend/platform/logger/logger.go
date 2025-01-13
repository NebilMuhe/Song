package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type log struct {
	logger *zap.Logger
}

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	extractFields(ctx context.Context) []zap.Field
	With(fields ...zap.Field) Logger
}

func NewLogger(logger *zap.Logger) Logger {
	return &log{
		logger: logger,
	}
}

func (l *log) With(fields ...zap.Field) Logger {
	lo := l.logger.With(fields...)
	return &log{
		logger: lo,
	}
}

func (l *log) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extractFields(ctx)...).Debug(msg, fields...)
}

func (l *log) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extractFields(ctx)...).Error(msg, fields...)
}

func (l *log) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extractFields(ctx)...).Fatal(msg, fields...)
}

func (l *log) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extractFields(ctx)...).Info(msg, fields...)
}

func (l *log) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extractFields(ctx)...).Warn(msg, fields...)
}

func (l *log) extractFields(ctx context.Context) []zap.Field {
	var fields []zap.Field

	fields = append(fields, zap.String("time", time.Now().Format(time.RFC3339)))
	if request_id, ok := ctx.Value("x-request-id").(string); ok {
		fields = append(fields, zap.String("request_id", request_id))
	}

	if user_id, ok := ctx.Value("x-user-id").(string); ok {
		fields = append(fields, zap.String("user_id", user_id))
	}

	if request_start_time, ok := ctx.Value("x-request-start-time").(time.Time); ok {
		fields = append(fields, zap.Time("request_start_time", request_start_time))
	}

	return fields

}
