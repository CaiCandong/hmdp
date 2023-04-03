package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hmdp/internal/app/middleware"
	"hmdp/internal/app/services"
	"hmdp/internal/domain/aggregate"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/infrastructure/redis"
	"hmdp/internal/interfaces/api"
	"hmdp/internal/interfaces/controller"
	"hmdp/pkg/logger"
)

func init() {
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	logger.InitializeLogger()
	DB := mysql.InitDB()
	redis.InitRedisStore()
	// 初始化repo
	showTypeRepo := mysql.NewShopTypeRepo(DB)
	blogRepo := mysql.NewBlogRepo(DB)
	userRepo := mysql.NewUserRepo(DB)
	// 初始化aggre
	userAggr := aggregate.NewUserAggregate(userRepo)
	// 初始化app/service
	blogSvc := services.NewBlogService(blogRepo)
	shopTypeSvc := services.NewShowTypeService(showTypeRepo)
	userSvc := services.NewUserService(userAggr)
	//初始化ctrl
	controller.NewBlogController(blogSvc)
	controller.NewShopTypeController(shopTypeSvc)
	controller.NewUserController(userSvc)
}

// DependenceInject 完成依赖注入
// func DependenceInject() {
//
// }
func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUser(mysql.NewUserRepo()))

	// 绑定用户相关的路由
	api.RegisterUserRoutes(router.Group("/user"))
	// 绑定商铺类型相关的路由
	//api.RegisterShopTypeRoutes(router.Group("/shop-type"))
	// 绑定博客相关的路由
	api.RegisterBlogRoutes(router.Group("/blog"))

	return router
}
func main() {
	r := InitRoute()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
