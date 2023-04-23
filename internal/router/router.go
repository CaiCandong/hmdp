package router

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/di"
	middleware2 "hmdp/internal/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	handlers := di.InitHandlers()

	router.Use(middleware2.EnableCookieSession())
	router.Use(middleware2.CurrentUserByToken())
	// 绑定博客相关的路由
	router.GET("/blog/of/me", handlers.BlogHandler.ListBlogsByUserId) // 获取点赞数
	//router.GET("/blog/likes/:id", handlers.BlogHandler.GetLikes)
	router.GET("/blog/hot", handlers.BlogHandler.BlogHot) // 当前热榜
	router.GET("/blog/:id", handlers.BlogHandler.FindBlogById)
	router.GET("/blog/of/follow")
	router.POST("/blog")         // 点赞
	router.PUT("/blog/like/:id") // 获取与用户相关的博客
	//绑定消费券相关的路由
	router.GET("/voucher/list/:shopId", handlers.VoucherHandler.ListVouchersByShopId)
	router.POST("/voucher-order/seckill/:id", handlers.VoucherHandler.SecKill)
	router.POST("/voucher/create", handlers.VoucherHandler.CreateVoucher)
	// 绑定商铺类型相关的路由
	router.GET("/shop-type/list", handlers.ShopTypeController.ListShopTypes)
	// 绑定商家相关的路由
	router.GET("/shop/:id", handlers.ShopHandler.FindShopById)
	router.GET("/shop/of/type", handlers.ShopHandler.OfType)
	router.PUT("/shop/:id", handlers.ShopHandler.Update)
	// 绑定用户相关的路由
	router.POST("/user/code", handlers.UserHandler.UserSendCode) // 短信登录&注册
	router.POST("/user/login", handlers.UserHandler.UserLogin)   // 注册用户登录接口
	router.Use(middleware2.AuthRequired())
	{
		router.GET("/user/me", handlers.UserHandler.UserMe)         // 用户个人主页
		router.GET("/user/info/:id", handlers.UserHandler.UserInfo) //用户详细信息
	}

	return router
}
