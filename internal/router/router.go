package router

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/di"
	"hmdp/internal/middleware"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	handlers := di.InitHandlers()

	//router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUserByToken())
	// 绑定博客相关的路由
	router.GET("/blog/of/me", handlers.BlogHandler.ListBlogsOfMe)       // 获取点赞数
	router.GET("/blog/of/user", handlers.BlogHandler.ListBlogsByUserId) // 获取用户的博客列表
	//router.GET("/blog/likes/:id", handlers.BlogHandler.GetLikes)
	router.GET("/blog/hot", handlers.BlogHandler.ListHotBlogs) // 当前热榜
	router.GET("/blog/:id", handlers.BlogHandler.FindBlogById)
	router.GET("/blog/of/follow", handlers.BlogHandler.ListBlogsBySubscription)    // 获取关注的用户的博客列表
	router.POST("/blog", handlers.BlogHandler.CreateBlog)                          // 创建博客
	router.GET("/blog/likes/:blogId", handlers.BlogHandler.ListLikedUsersByBlogId) // 获取点赞的用户列表

	//绑定消费券相关的路由
	router.GET("/voucher/list/:shopId", handlers.VoucherHandler.ListVouchersByShopId)
	router.POST("/voucher-order/seckill/:id", handlers.VoucherHandler.SecKill)
	router.POST("/voucher/create", handlers.VoucherHandler.CreateVoucher)
	// 绑定商铺类型相关的路由
	router.GET("/shop-type/list", handlers.ShopTypeController.ListShopTypes)
	// 绑定商家相关的路由
	router.GET("/shop/:id", handlers.ShopHandler.FindShopById)
	router.GET("/shop/of/type", handlers.ShopHandler.OfType)
	router.GET("/shop/of/name", handlers.ShopHandler.ListShopsByName)
	router.PUT("/shop/:id", handlers.ShopHandler.Update)
	// 绑定用户相关的路由
	router.POST("/user/code", handlers.UserHandler.UserSendCode)            // 短信登录&注册
	router.POST("/user/login", handlers.UserHandler.UserLogin)              // 注册用户登录接口
	router.GET("/user/:id", handlers.UserHandler.FindUserById)              // 根据用户id获取用户信息
	router.PUT("/follow/:id/:follow", handlers.UserHandler.FollowUser)      // 关注用户
	router.GET("/follow/or/not/:id", handlers.UserHandler.IsFollowed)       // 判断是否关注
	router.GET("/follow/common/:userId", handlers.UserHandler.CommonFollow) // 共同关注
	router.Use(middleware.AuthRequired())
	{
		router.PUT("/blog/like/:blogId", handlers.BlogHandler.LikeBlogByUserId) // 点赞
		router.POST("/upload/blog", handlers.BlogHandler.UploadBlogImg)         // 上传博客图片
		router.GET("/upload/blog/delete", handlers.BlogHandler.DeleteBlogImg)   // 删除博客图片
		router.POST("/")
		router.GET("/user/me", handlers.UserHandler.UserMe)         // 用户个人主页
		router.GET("/user/info/:id", handlers.UserHandler.UserInfo) //用户详细信息
	}
	return router
}
