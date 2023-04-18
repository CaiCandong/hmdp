package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type VoucherReq struct {
}

func (v *VoucherReq) D2EVoucherList(req *dto.VoucherListReq) *entity.Voucher {
	return &entity.Voucher{}
}

func (v *VoucherReq) D2EVoucherCreate(req *dto.VoucherCreateReq) (*entity.Voucher, *entity.SeckillVoucher) {
	voucher := &entity.Voucher{
		Title:       req.Title,
		SubTitle:    req.SubTitle,
		PayValue:    req.PayValue,
		ActualValue: req.ActualValue,
	}
	if req.Type == 0 {
		return voucher, nil
	}
	seckillVoucher := &entity.SeckillVoucher{
		Stock:     req.Stock,
		BeginTime: req.BeginTime,
		EndTime:   req.EndTime,
	}
	return voucher, seckillVoucher
}
