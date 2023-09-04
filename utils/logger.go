package utils

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"app.log"}
	logger, err := config.Build()
	if err != nil {
		log.Fatal(err, "Failed to initialize logger")
	}

	Logger = logger.Sugar()
	defer logger.Sync()

	Logger.Info("Logger initialized")
}