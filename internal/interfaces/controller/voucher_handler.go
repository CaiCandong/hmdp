package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
)

type VoucherHandler struct {
	VoucherService services.IVoucherService
}

func NewVoucherHandler(VoucherService services.IVoucherService) *VoucherHandler {
	return &VoucherHandler{VoucherService: VoucherService}
}

// Seckill 消费券秒杀
func (v *VoucherHandler) SecKill(ctx *gin.Context) {
	req := &dto.VoucherSecKillReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.VoucherService.VoucherSecKill(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
func (v *VoucherHandler) ListByShopId(ctx *gin.Context) {
	req := &dto.VoucherListReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.VoucherService.VoucherListByShopId(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (v *VoucherHandler) Create(ctx *gin.Context) {
	req := &dto.VoucherCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.VoucherService.VoucherCreate(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
