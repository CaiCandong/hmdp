package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
	"time"
)

type IVoucherService interface {
	VoucherListByShopId(ctx *gin.Context, req *dto.VoucherListReq) ([]*dto.VoucherListRsp, error)
	VoucherSecKill(ctx *gin.Context, req *dto.VoucherSecKillReq) (*dto.VoucherSecKillRsp, error)
	VoucherCreate(ctx *gin.Context, req *dto.VoucherCreateReq) (*dto.VoucherCreateRsp, error)
}

func NewVoucherService(VoucherAgg repository.IVoucherRepo) IVoucherService {
	return &VoucherService{
		VoucherAgg: VoucherAgg,
		VoucherReq: &assembler.VoucherReq{},
		VoucherRsp: &assembler.VoucherRsp{},
	}
}

type VoucherService struct {
	VoucherAgg repository.IVoucherRepo //调用领域聚合层|领域服务层
	VoucherReq *assembler.VoucherReq   //处理请求
	VoucherRsp *assembler.VoucherRsp   //处理响应
}

func (v *VoucherService) VoucherListByShopId(ctx *gin.Context, req *dto.VoucherListReq) ([]*dto.VoucherListRsp, error) {
	VoucherList, err := v.VoucherAgg.GetByShopId(ctx, req.ShopID)
	if err != nil {
		return nil, err
	}
	return v.VoucherRsp.E2DVoucherList(VoucherList), nil
}
func (v *VoucherService) VoucherSecKill(ctx *gin.Context, req *dto.VoucherSecKillReq) (*dto.VoucherSecKillRsp, error) {
	// TODO: Implement it
	order := &entity.VoucherOrder{
		Model:      gorm.Model{},
		UserId:     0,
		VoucherId:  0,
		PayType:    0,
		Status:     0,
		PayTime:    time.Time{},
		UseTime:    time.Time{},
		RefundTime: time.Time{},
	}
	return v.VoucherRsp.E2DVoucherSecKill(order), nil
}

func (v *VoucherService) VoucherCreate(ctx *gin.Context, req *dto.VoucherCreateReq) (*dto.VoucherCreateRsp, error) {
	voucher, SeckillVoucher := v.VoucherReq.D2EVoucherCreate(req)
	//v.VoucherAgg.Create(ctx, voucher)
	//v.VoucherAgg.Create(ctx, SeckillVoucher)
	return v.VoucherRsp.E2DVoucherCreate(voucher, SeckillVoucher), nil
}
