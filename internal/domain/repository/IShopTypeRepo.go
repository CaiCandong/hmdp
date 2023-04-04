package repository

import "hmdp/internal/domain/entity"

type IShopType interface {
	GetShopTypeList() ([]*entity.ShowType, error)
}
