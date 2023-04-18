package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hmdp/docs"
	"hmdp/internal/app/middleware"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/di"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	db := mysql.GetDB()
	handlers := di.InitHandlers(db)

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUserByToken())
	// 绑定博客相关的路由
	router.GET("/blog/of/me", handlers.BlogHandler.GetBlogsByUserId) // 获取点赞数
	router.GET("/blog/likes/:id", handlers.BlogHandler.GetLikes)
	router.GET("/blog/hot", handlers.BlogHandler.Hot) // 当前热榜
	router.GET("/blog/:id", handlers.BlogHandler.Find)
	router.GET("/blog/of/follow")
	router.POST("/blog")         // 点赞
	router.PUT("/blog/like/:id") // 获取与用户相关的博客
	//绑定消费券相关的路由
	router.GET("voucher/list/:shopId", handlers.VoucherHandler.ListByShopId)
	router.POST("voucher-order/seckill/:id", handlers.VoucherHandler.SecKill)
	router.POST("voucher/create", handlers.VoucherHandler.Create)
	// 绑定商铺类型相关的路由
	router.GET("/shop-type/list", handlers.ShopTypeController.List)
	// 绑定商家相关的路由
	router.GET("/shop/:id", handlers.ShopHandler.GetShop)
	router.GET("/shop/of/type", handlers.ShopHandler.OfType)
	router.PUT("/shop/:id", handlers.ShopHandler.Update)
	// 绑定用户相关的路由
	router.POST("/user/code", handlers.UserHandler.SendCode) // 短信登录&注册
	router.POST("/user/login", handlers.UserHandler.Login)   // 注册用户登录接口
	router.Use(middleware.AuthRequired())
	{
		router.GET("/user/me", handlers.UserHandler.Me)         // 用户个人主页
		router.GET("/user/info/:id", handlers.UserHandler.Info) //用户详细信息
	}

	return router
}
