package mysql

import (
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type ShopTypeRepo struct {
	db *gorm.DB
}

func (repo *ShopTypeRepo) GetShopTypeList() ([]*entity.ShowType, error) {
	var showtypes []*entity.ShowType
	result := repo.db.Find(&showtypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return showtypes, nil
}

func NewShopTypeRepo(DB *gorm.DB) repository.IShopTypeRepo {
	return &ShopTypeRepo{DB}
}
