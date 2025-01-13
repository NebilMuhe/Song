package initator

import (
	"context"
	"song/platform/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func InitConfig(path,name,config_type string, log logger.Logger) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(config_type)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(context.Background(),"failed to read config file",zap.Error(err))
	}
}