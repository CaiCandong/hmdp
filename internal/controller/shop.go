package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/serializer"
	"hmdp/internal/service"
	"net/http"
)

type ShopHandler struct {
	svc *service.ShopService
}

func NewShopHandler(svc *service.ShopService) *ShopHandler {
	return &ShopHandler{svc}
}

func (handler *ShopHandler) ListShopsByType(ctx *gin.Context) {
	req := &dto.ListShopsByTypeReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := handler.svc.ListShopsByType(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
func (handler *ShopHandler) ListShopsByName(ctx *gin.Context) {
	req := &dto.ListShopsByNameReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := handler.svc.ListShopsByName(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
func (handler *ShopHandler) FindShopById(ctx *gin.Context) {
	req := &dto.FindShopByIdReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := handler.svc.FindShopById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (handler *ShopHandler) Update(ctx *gin.Context) {
	req := &dto.UpdateShopByIdReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := handler.svc.UpdateShopById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
