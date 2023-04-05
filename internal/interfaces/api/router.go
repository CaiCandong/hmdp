package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/middleware"
	"hmdp/internal/app/services"
	"hmdp/internal/domain/aggregate"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/controller"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EnableCookieSession())
	router.Use(middleware.CurrentUser(mysql.NewUserRepo(mysql.DB)))

	// 绑定用户相关的路由
	RegisterUserRoutes(router.Group("/user"))
	// 绑定商铺类型相关的路由
	RegisterShopTypeRoutes(router.Group("/shop-type"))
	// 绑定博客相关的路由
	RegisterBlogRoutes(router.Group("/blog"))
	// 绑定商家相关的路由
	RegisterShopRoutes(router.Group("/shop"))
	//绑定消费券相关的路由
	RegisterVoucherRoutes(router.Group("/voucher"))

	return router
}

// RegisterVoucherRoutes 注册用户相关的路由
func RegisterVoucherRoutes(r *gin.RouterGroup) {
	// 手动依赖注入
	voucherCtrl := controller.NewVoucherHandler(
		&services.VoucherService{
			VoucherAgg: &mysql.Voucher{DB: mysql.DB},
			VoucherReq: &assembler.VoucherReq{},
			VoucherRsp: &assembler.VoucherRsp{},
		})
	////
	r.GET("/list/:shopId", voucherCtrl.ListByShopId)
}

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(r *gin.RouterGroup) {
	// 手动依赖注入
	userCtrl := controller.NewUserHandler(&services.UserService{
		UserAgg: &aggregate.UserAggregate{UserRepo: &mysql.UserRepo{DB: mysql.DB}},
		UserReq: &assembler.UserReq{},
		UserRsp: &assembler.UserRsp{},
	})
	//// 短信登录&注册
	r.POST("/code", userCtrl.SendCode)
	//
	//// 注册用户登录接口
	r.POST("/login", userCtrl.Login)
	//
	//// 注册用户注册接口
	//r.POST("/register", userCtrl.Register)
	//
	r.Use(middleware.AuthRequired())
	//{
	//	// 用户个人主页
	r.GET("/me", userCtrl.Me)
	//
	//	// 注册用户注销接口
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
func RegisterShopTypeRoutes(r *gin.RouterGroup) {

	shopTypeCtrl := controller.NewShopTypeController(&services.ShowTypeService{
		ShopTypeRepo: &mysql.ShopTypeRepo{DB: mysql.DB},
		ShopTypeReq:  &assembler.ShopTypeReq{},
		ShopTypeRsp:  &assembler.ShopTypeRsp{},
	})

	r.GET("/list", shopTypeCtrl.List)
}

// RegisterShopRoutes 注册用商户相关的路由
func RegisterShopRoutes(r *gin.RouterGroup) {
	shopHandler := controller.ShopHandler{
		ShopService: &services.ShopService{
			ShopRepo: &mysql.ShopRepo{DB: mysql.DB},
			ShopReq:  assembler.ShopReq{},
			ShopRsp:  assembler.ShopRsp{},
		},
	}
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
func RegisterBlogRoutes(r *gin.RouterGroup) {
	// 手动依赖注入
	blogCtrl := controller.NewBlogController(&services.BlogService{
		BlogRepo: &mysql.BlogRepo{DB: mysql.DB},
		BlogReq:  &assembler.BlogReq{},
		BlogRsp:  &assembler.BlogRsp{},
	})
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
