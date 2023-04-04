package repository

import (
	"context"
	"hmdp/internal/domain/entity"
)

type IShopRepo interface {
	GetShopByType(ctx context.Context, shopTypeId uint, page int) ([]*entity.Shop, error)
}
