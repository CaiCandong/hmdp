package mysql

import (
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type ShopTypeRepo struct {
}

func NewShopTypeRepo() repository.IShowType {
	return &ShopTypeRepo{}
}

func (shop *ShopTypeRepo) GetShopTypeList() ([]entity.ShowType, error) {
	var showtypes []entity.ShowType
	result := DB.Find(&showtypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return showtypes, nil
}
