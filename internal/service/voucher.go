package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"hmdp/pkg/utils"
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

// VoucherSecKill 秒杀代金券
func (s *VoucherService) VoucherSecKill(ctx *gin.Context, req *dto.VoucherSecKillReq) (*dto.VoucherSecKillRsp, error) {
	// 拦截器保证用户已经登录并存放在context中
	user, _ := ctx.Get("user")
	//id := req.VoucherID
	// 开启数据库事务
	tx := s.DB.Begin()
	defer tx.Commit()
	// 1. 查询优惠券 & 判断库存
	voucher := model.SeckillVoucher{}
	if err := tx.Select("stock", "begin_time", "end_time").Where("id = ?", req.VoucherID).First(&voucher).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("优惠券不存在")
	}
	if !utils.IsInTime(voucher.BeginTime, voucher.EndTime) {
		tx.Rollback()
		return nil, errors.New("不在活动时间内")
	}
	if voucher.Stock < 1 {
		tx.Rollback()
		return nil, errors.New("库存不足")
	}
	// 2. 扣除库存
	if err := tx.Model(&voucher).Update("stock", voucher.Stock-1).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("库存不足")
	}
	// 3. 创建订单
	order := model.VoucherOrder{
		UserId:    user.(*model.User).ID,
		VoucherId: utils.SnowFlakeID(),
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("创建订单失败")
	}
	return s.Rsp.E2DVoucherSecKill(&order), nil
}

func (s *VoucherService) ListVouchersByShopId(ctx *gin.Context, req *dto.VoucherListReq) (*dto.VoucherListRsp, error) {
	return nil, nil
}

func (s *VoucherService) CreateVoucher(ctx *gin.Context, req *dto.VoucherCreateReq) (*dto.VoucherCreateRsp, error) {
	return nil, nil
}
