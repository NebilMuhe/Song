package initator

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	log := InitLogger()
	log.Info(context.Background(), "logger initialized")

	log.Info(context.Background(), "initializing config")
	InitConfig("config", "config", "yaml", log)
	log.Info(context.Background(), "config initialized")

	log.Info(context.Background(), "Initializing MongoDB")
	InitDB(viper.GetString("database.url"), log)
	log.Info(context.Background(), "MongoDB initialized")

	log.Info(context.Background(), "Initializing Server")
	server := gin.New()
	gin.SetMode(gin.ReleaseMode)
	log.Info(context.Background(), "Server initialized")

	httpServer := &http.Server{
		Addr:    viper.GetString("server.host") + ":" + viper.GetString("server.port"),
		Handler: server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	signal.Notify(quit, os.Signal(syscall.SIGTERM))

	go func() {
		log.Info(context.Background(), "Starting Server",
			zap.String("host", viper.GetString("server.host")),
			zap.String("port", viper.GetString("server.port")),
		)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(context.Background(), "failed to start server", zap.Error(err))
		}
	}()

	sig := <-quit
	log.Info(context.Background(), "server stopped with signal", zap.String("signal", sig.String()))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Info(context.Background(), "shutting down server")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Warn(context.Background(), "failed to shutdown server", zap.Error(err))
	} else {
		log.Info(context.Background(), "server shutdown successfully")
	}
}
