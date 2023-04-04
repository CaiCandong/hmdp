package mysql

import (
	"context"
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
)

type ShopRepo struct {
	DB *gorm.DB
}

func (s *ShopRepo) GetShopByType(ctx context.Context, shopTypeId uint, page int) ([]*entity.Shop, error) {
	var shops []*entity.Shop
	var pageSize int = 5
	err := s.DB.Model(&entity.Shop{}).
		Where("type_id = ? ", shopTypeId).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&shops).Error
	return shops, err
}
