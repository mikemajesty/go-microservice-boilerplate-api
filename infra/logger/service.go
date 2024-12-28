package infra

import (
	"context"
	"log/slog"
)

type loggerAdapter struct{}

var logger *slog.Logger
var ctx context.Context

func CreateLogger() LoggerAdapter {
	return &loggerAdapter{}
}

type InfoAttr struct {
	Value any
	Key   string
}

func (l *loggerAdapter) Info(message string, attrs ...InfoAttr) {
	var attributes []any
	for _, value := range attrs {
		attributes = append(attributes, slog.Attr{Key: value.Key, Value: slog.AnyValue(value.Value)})
	}
	attributes = append(attributes, slog.Attr{Key: "traceID", Value: slog.AnyValue(ctx.Value("traceID"))})
	logger.InfoContext(ctx, message, attributes...)
}

func (l *loggerAdapter) Error(message string, attrs ...InfoAttr) {
	var attributes []any
	for _, value := range attrs {
		attributes = append(attributes, slog.Attr{Key: value.Key, Value: slog.AnyValue(value.Value)})
	}
	attributes = append(attributes, slog.Attr{Key: "traceID", Value: slog.AnyValue(ctx.Value("traceID"))})
	logger.ErrorContext(ctx, message, attributes...)
}

func (l *loggerAdapter) Connect() {
	logger = slog.Default()
}

func (l *loggerAdapter) Logger() *slog.Logger {
	return logger
}

func (l *loggerAdapter) SetContext(_ctx context.Context) {
	ctx = _ctx
}
