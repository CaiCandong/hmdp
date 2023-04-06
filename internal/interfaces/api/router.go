package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hmdp/docs"
	"hmdp/internal/app/middleware"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/controller"
	"hmdp/internal/interfaces/di"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	db := mysql.GetDB()

	userRepo := di.InitUserRepo(db)

	blogHandler := di.InitBlogHandler(db)
	shopTypeCtrl := di.InitShopTypeHandler(db)
	voucherCtrl := di.InitVoucherHandler(db)
	shopHandler := di.InitShopHandler(db)
	userCtrl := di.InitUserHandler(db)

	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUser(userRepo))
	// 绑定用户相关的路由
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	RegisterUserRoutes(router.Group("/user"), userCtrl)
	// 绑定商铺类型相关的路由
	RegisterShopTypeRoutes(router.Group("/shop-type"), shopTypeCtrl)
	// 绑定博客相关的路由
	RegisterBlogRoutes(router.Group("/blog"), blogHandler)
	// 绑定商家相关的路由
	RegisterShopRoutes(router.Group("/shop"), shopHandler)
	//绑定消费券相关的路由
	RegisterVoucherRoutes(router.Group("/voucher"), voucherCtrl)

	return router
}

// RegisterVoucherRoutes 注册用户相关的路由
func RegisterVoucherRoutes(r *gin.RouterGroup, voucherCtrl *controller.VoucherHandler) {

	r.GET("/list/:shopId", voucherCtrl.ListByShopId)
}

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(r *gin.RouterGroup, userCtrl *controller.UserHandler) {
	// 短信登录&注册
	r.POST("/code", userCtrl.SendCode)
	//// 注册用户登录接口
	r.POST("/login", userCtrl.Login)
	// 注册用户注册接口
	//r.POST("/register", userCtrl.Register)
	r.Use(middleware.AuthRequired())
	// 用户个人主页
	r.GET("/me", userCtrl.Me)
	// 注册用户注销接口
	//	r.DELETE("/logout", userCtrl.Logout)
	//
	//	// 查看用户详情
	r.GET("/info/:id", userCtrl.Info)
	//
	//	// 注册修改个人信息接口
	//	r.PUT("/:id", userCtrl.Update)
	//
	//}

}

// RegisterUploadRoutes 注册用户相关的路由
func RegisterUploadRoutes(r *gin.RouterGroup) {
	uploadCtrl := controller.NewUploadController()
	// 上传文件
	r.POST("/blog", uploadCtrl.Create)

	// 删除文件
	r.GET("/blog/delete", uploadCtrl.Delete)
}

// RegisterShopTypeRoutes 注册用户相关的路由
func RegisterShopTypeRoutes(r *gin.RouterGroup, shopTypeCtrl *controller.ShopTypeController) {
	r.GET("/list", shopTypeCtrl.List)
}

// RegisterShopRoutes 注册用商户相关的路由
func RegisterShopRoutes(r *gin.RouterGroup, shopHandler *controller.ShopHandler) {
	// 发布博客
	r.POST("")
	// 点赞
	r.PUT("/like/:id")
	// 获取与用户相关的博客
	r.GET("/of/type", shopHandler.OfType)
	// 获取点赞数
	r.GET("/like/:id")
	// 当前热榜
	r.GET("/hot")
	r.GET("/:id", shopHandler.GetShop)
	r.GET("/of/follow")

}

// RegisterFollowerRoutes 注册用户相关的路由
func RegisterFollowerRoutes(r *gin.RouterGroup) {
	followCtrl := controller.NewFollowerController()
	// 订阅
	r.PUT("/:id/:isFollow", followCtrl.Follow)
	// 查看订阅状态
	r.GET("/or/not/:id", followCtrl.IsFollow)
	// 公共
	r.GET("/common/:id", followCtrl.Common)
}

// RegisterBlogRoutes 注册用户相关的路由
func RegisterBlogRoutes(r *gin.RouterGroup, blogCtrl *controller.BlogController) {
	// 发布博客
	r.POST("")
	// 点赞
	r.PUT("/like/:id")
	// 获取与用户相关的博客
	r.GET("/of/me", blogCtrl.GetBlogsByUserId)
	// 获取点赞数
	r.GET("/likes/:id", blogCtrl.GetLikes)
	// 当前热榜
	r.GET("/hot", blogCtrl.Hot)
	r.GET("/:id", blogCtrl.Find)
	r.GET("/of/follow")
}
