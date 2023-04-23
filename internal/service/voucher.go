package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/dto"
)

type VoucherService struct {
	DB  *gorm.DB
	Req *assembler.VoucherReq
	Rsp *assembler.VoucherRsp
}

func NewVoucherService(db *gorm.DB) *VoucherService {
	return &VoucherService{
		DB:  db,
		Req: &assembler.VoucherReq{},
		Rsp: &assembler.VoucherRsp{},
	}
}

func (s VoucherService) VoucherSecKill(ctx *gin.Context, req *dto.VoucherSecKillReq) (*dto.VoucherSecKillRsp, error) {
	return nil, nil
}

func (s VoucherService) ListVouchersByShopId(ctx *gin.Context, req *dto.VoucherListReq) (*dto.VoucherListRsp, error) {
	return nil, nil
}

func (s VoucherService) CreateVoucher(ctx *gin.Context, req *dto.VoucherCreateReq) (*dto.VoucherCreateRsp, error) {
	return nil, nil
}
