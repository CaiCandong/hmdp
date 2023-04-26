package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/serializer"
	"hmdp/internal/service"
	"net/http"
)

type VoucherHandler struct {
	svc *service.VoucherService
}

func NewVoucherHandler(svc *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{svc: svc}
}

// SecKill 消费券秒杀
func (v *VoucherHandler) SecKill(ctx *gin.Context) {
	req := &dto.VoucherSecKillReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.svc.VoucherSecKill(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (v *VoucherHandler) ListVouchersByShopId(ctx *gin.Context) {
	req := &dto.VoucherListReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.svc.ListVouchersByShopId(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (v *VoucherHandler) CreateVoucher(ctx *gin.Context) {
	req := &dto.VoucherCreateReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := v.svc.CreateVoucher(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
