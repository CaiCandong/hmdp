package main

import (
	"fmt"
	"github.com/spf13/viper"
	"hmdp/model"
	"hmdp/routes"
	"hmdp/utils"
)

func init() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	utils.InitializeLogger()
	model.InitDB()
}

func main() {
	r := routes.InitRoute()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
