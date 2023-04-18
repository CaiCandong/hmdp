package mysql

import (
	"context"
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type ShopRepo struct {
	db *gorm.DB
}

func NewShopRepo(db *gorm.DB) repository.IShopRepo {
	return &ShopRepo{db: db}
}
func (s *ShopRepo) GetShopByType(ctx context.Context, shopTypeId uint, page int) ([]*entity.Shop, error) {
	var shops []*entity.Shop
	var pageSize int = 5
	err := s.db.Model(&entity.Shop{}).
		Where("type_id = ? ", shopTypeId).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&shops).Error
	return shops, err
}

func (s *ShopRepo) GetShopById(ctx context.Context, shop *entity.Shop) error {
	return s.db.First(shop).Error
}

func (s *ShopRepo) UpdateById(ctx context.Context, shop *entity.Shop) error {
	return s.db.Updates(shop).Error
}
