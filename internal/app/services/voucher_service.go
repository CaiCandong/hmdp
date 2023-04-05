package services

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/repository"
)

type VoucherService struct {
	VoucherAgg repository.IVoucherAgg //调用领域聚合层|领域服务层
	VoucherReq *assembler.VoucherReq  //处理请求
	VoucherRsp *assembler.VoucherRsp  //处理响应
}

func (v *VoucherService) VoucherListByShopId(ctx *gin.Context, req *dto.VoucherListReq) ([]*dto.VoucherListRsp, error) {
	VoucherList, err := v.VoucherAgg.GetByShopId(ctx, req.ShopID)
	if err != nil {
		return nil, err
	}
	return v.VoucherRsp.E2DVoucherList(VoucherList), nil
}
