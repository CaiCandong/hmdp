package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/pkg/logger"
)

func init() {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.InitializeLogger()
	mysql.InitDB()
}

func main() {
	r := InitRoute()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
