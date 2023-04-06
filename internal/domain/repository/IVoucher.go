package repository

import (
	"context"
	"hmdp/internal/domain/entity"
)

type IVoucherRepo interface {
	GetByShopId(ctx context.Context, shopId uint) ([]*entity.Voucher, error)
}
