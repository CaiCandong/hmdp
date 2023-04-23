package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/serializer"
	"hmdp/internal/service"
	"net/http"
)

type ShopTypeHandler struct {
	svc *service.ShopTypeService
}

func NewShopTypeController(shopTypeService *service.ShopTypeService) *ShopTypeHandler {
	return &ShopTypeHandler{shopTypeService}
}

func (s *ShopTypeHandler) ListShopTypes(ctx *gin.Context) {
	req := &dto.ShopTypeListReq{}
	list, err := s.svc.ListShopTypes(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(list))
}
