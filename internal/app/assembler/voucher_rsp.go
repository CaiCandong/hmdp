package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type VoucherRsp struct {
}

func (v *VoucherRsp) E2DVoucherList(es []*entity.Voucher) []*dto.VoucherListRsp {
	ret := make([]*dto.VoucherListRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.VoucherListRsp{
			Title:       e.Title,
			SubTitle:    e.SubTitle,
			PayValue:    e.PayValue,
			ActualValue: e.ActualValue,
			Stock:       e.Stock,
		}
	}
	return ret
}
