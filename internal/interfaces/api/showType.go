package api

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/services"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/controller"
)

// RegisterShopTypeRoutes 注册用户相关的路由
func RegisterShopTypeRoutes(r *gin.RouterGroup) {

	shopTypeCtrl := controller.NewShopTypeController(&services.ShowTypeService{
		ShopTypeRepo: &mysql.ShopTypeRepo{DB: mysql.DB},
		ShopTypeReq:  &assembler.ShopTypeReq{},
		ShopTypeRsp:  &assembler.ShopTypeRsp{},
	})

	r.GET("/list", shopTypeCtrl.List)
}
