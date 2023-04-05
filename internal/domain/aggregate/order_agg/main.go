package order_agg

import "hmdp/internal/domain/entity"

type OrderAgg struct {
	VoucherOrder *entity.VoucherOrder
	User         *entity.User
	Voucher      *entity.Voucher
}
