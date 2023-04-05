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
