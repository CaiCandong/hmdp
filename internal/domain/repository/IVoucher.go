package repository

import (
	"context"
	"hmdp/internal/domain/entity"
)

type IVoucherAgg interface {
	GetByShopId(ctx context.Context, shopId uint) ([]*entity.Voucher, error)
}
