package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/services"
	"hmdp/internal/infrastructure/mysql"
	"net/http"
)

type ShopTypeController interface {
	List(c *gin.Context)
}

func NewShowController() ShopTypeController {
	// 依赖注入
	showTypeRepo := mysql.NewShopTypeRepo()
	userService := services.NewShopTypeService(services.WithShopTypeRepo(showTypeRepo))
	return &ShopTypeControllerImp{userService}
}

type ShopTypeControllerImp struct {
	shopTypeService services.IShopTypeService
}

func (s *ShopTypeControllerImp) List(ctx *gin.Context) {
	resp, err := s.shopTypeService.List()
	if err != nil {
		resp.Success = false
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Success = true
	ctx.JSON(http.StatusOK, resp)
}
