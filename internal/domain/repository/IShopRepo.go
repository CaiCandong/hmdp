package repository

import (
	"context"
	"hmdp/internal/domain/entity"
)

type IShopRepo interface {
	GetShopByType(ctx context.Context, shopTypeId uint, page int) ([]*entity.Shop, error)
	GetShopById(ctx context.Context, shop *entity.Shop) error
	UpdateById(ctx context.Context, shop *entity.Shop) error
}
