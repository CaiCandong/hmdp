package repository

import "hmdp/internal/domain/entity"

type IShowType interface {
	GetShopTypeList() ([]*entity.ShowType, error)
}
