package mysql

import (
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type ShopTypeRepo struct {
	DB *gorm.DB
}

func (shop *ShopTypeRepo) GetShopTypeList() ([]*entity.ShowType, error) {
	var showtypes []*entity.ShowType
	result := shop.DB.Find(&showtypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return showtypes, nil
}

func NewShopTypeRepo(DB *gorm.DB) repository.IShopType {
	return &ShopTypeRepo{DB}
}
