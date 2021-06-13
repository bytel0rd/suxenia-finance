package utils

import (
	"go.uber.org/zap"
)

var LoggerInstance *zap.SugaredLogger = nil

func init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	LoggerInstance = logger.Sugar()
}
