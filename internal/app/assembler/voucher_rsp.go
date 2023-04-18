package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
	"time"
)

type VoucherRsp struct {
}

func (v *VoucherRsp) E2DVoucherSecKill(e *entity.VoucherOrder) *dto.VoucherSecKillRsp {
	return &dto.VoucherSecKillRsp{
		OrderId: e.ID,
	}
}

func (v *VoucherRsp) E2DVoucherList(es []*entity.Voucher) []*dto.VoucherListRsp {
	ret := make([]*dto.VoucherListRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.VoucherListRsp{
			Id:          e.ID,
			Type:        true,
			Title:       e.Title,
			SubTitle:    e.SubTitle,
			PayValue:    e.PayValue,
			ActualValue: e.ActualValue,
			Stock:       e.Stock,
			BeginTime:   time.Now(),
			// 返回两个小时之后
			EndTime: time.Now().Add(time.Hour * 2),
		}
	}
	return ret
}

func (v *VoucherRsp) E2DVoucherCreate(e *entity.Voucher, s *entity.SeckillVoucher) *dto.VoucherCreateRsp {
	return &dto.VoucherCreateRsp{
		Status: 1,
	}
}
