package initator

import (
	"log"
	"song/platform/logger"

	"go.uber.org/zap"
)

func InitLogger() logger.Logger {
	logg,err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	return logger.NewLogger(logg)
}