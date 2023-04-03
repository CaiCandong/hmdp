package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
)

type ShopTypeController struct {
	shopTypeService services.IShopTypeService
}

func NewShopTypeController(shopTypeService services.IShopTypeService) *ShopTypeController {
	return &ShopTypeController{shopTypeService}
}
func (s *ShopTypeController) List(ctx *gin.Context) {
	req := &dto.ShopTypeListReq{}
	list, err := s.shopTypeService.List(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(list))
}
