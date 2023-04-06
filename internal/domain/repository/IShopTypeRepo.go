package repository

import "hmdp/internal/domain/entity"

type IShopTypeRepo interface {
	GetShopTypeList() ([]*entity.ShowType, error)
}
