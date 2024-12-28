package infra

import (
	"context"
	"log/slog"
)

type LoggerAdapter interface {
	Connect()
	Logger() *slog.Logger
	Error(message string, attrs ...InfoAttr)
	SetContext(ctx context.Context)
	Info(message string, attrs ...InfoAttr)
}
