package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
)

type VoucherHandler struct {
	VoucherService *services.VoucherService
}

func NewVoucherHandler(service *services.VoucherService) *VoucherHandler {
	return &VoucherHandler{VoucherService: service}
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
