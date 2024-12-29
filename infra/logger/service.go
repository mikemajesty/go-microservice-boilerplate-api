package infra

import (
	"context"
	"encoding/json"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type loggerOtherAdapter struct{}

var logger *log.Logger
var ctx context.Context

func (l *loggerOtherAdapter) Connect(mv *MongoWriter) {
	_logger := log.New()
	_logger.SetFormatter(&log.JSONFormatter{})

	_logger.SetOutput(io.MultiWriter(os.Stdout, mv))

	_logger.SetLevel(log.InfoLevel)
	logger = _logger
}

func (l *loggerOtherAdapter) Error(message string, attrs log.Fields) {
	attrs["traceId"] = ctx.Value("traceId")
	logger.WithContext(ctx).WithFields(attrs).Error(message)
}

func (l *loggerOtherAdapter) Info(message string, attrs log.Fields) {
	attrs["traceId"] = ctx.Value("traceId")
	logger.WithContext(ctx).WithFields(attrs).Info(message)
}

func (l *loggerOtherAdapter) Logger() *log.Logger {
	return logger
}

func (l *loggerOtherAdapter) SetContext(_ctx context.Context) {
	ctx = _ctx
}

func CreateLogger() LoggerAdapter {
	return &loggerOtherAdapter{}
}

type MongoWriter struct {
	Client *mongo.Client
}

type write struct {
	Level   string `bson:"level"`
	Msg     string `bson:"msg"`
	Time    string `bson:"time"`
	TraceId string `bson:"traceId"`
}

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	c := mw.Client.Database("go-microservice-boilerplate-api").Collection("logs")

	data := write{}
	json.Unmarshal(p, &data)

	_, err = c.InsertOne(context.TODO(), data)
	if err != nil {
		return
	}
	return len(p), nil
}
