package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
)

type ShopHandler struct {
	ShopService services.IShopService
}

func NewShopHandler(ShopService services.IShopService) *ShopHandler {
	//return
	return &ShopHandler{ShopService}
}

func (s *ShopHandler) OfType(ctx *gin.Context) {
	req := &dto.ShopOfTypeReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := s.ShopService.OfType(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
