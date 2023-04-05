package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/infrastructure/redis"
	"hmdp/internal/interfaces/api"
	"hmdp/pkg/logger"
)

func init() {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.InitializeLogger()
	mysql.InitDB()
	redis.InitRedisStore()
}

func main() {
	r := api.InitRoute()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
