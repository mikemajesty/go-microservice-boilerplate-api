package infra

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type LogAttrInput = log.Fields

type LoggerAdapter interface {
	Connect(database *MongoWriter)
	Logger() *log.Logger
	Error(message string, attrs LogAttrInput)
	SetContext(ctx context.Context)
	Info(message string, attrs LogAttrInput)
}
