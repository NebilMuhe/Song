package initator

import (
	"context"
	"song/platform/logger"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
)

func InitDB(url string, log logger.Logger) *mongo.Client{
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Warn(context.Background(), "failed to disconnect from mongodb", zap.Error(err))
		}
	}()
	if err != nil {
		log.Fatal(context.Background(), "failed to connect to mongodb", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(context.Background(), "failed to ping mongodb", zap.Error(err))
	}

	return client
}
