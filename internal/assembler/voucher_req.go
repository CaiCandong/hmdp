package assembler

import (
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type VoucherReq struct {
}

func (v *VoucherReq) D2EVoucherList(req *dto.VoucherListReq) *model.Voucher {
	return &model.Voucher{}
}

func (v *VoucherReq) D2EVoucherCreate(req *dto.VoucherCreateReq) (*model.Voucher, *model.SeckillVoucher) {
	voucher := &model.Voucher{
		Title:       req.Title,
		SubTitle:    req.SubTitle,
		PayValue:    req.PayValue,
		ActualValue: req.ActualValue,
	}
	if req.Type == 0 {
		return voucher, nil
	}
	seckillVoucher := &model.SeckillVoucher{
		Stock:     req.Stock,
		BeginTime: req.BeginTime,
		EndTime:   req.EndTime,
	}
	return voucher, seckillVoucher
}
