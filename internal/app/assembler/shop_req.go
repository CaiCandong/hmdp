package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type ShopReq struct {
}

func (s *ShopReq) D2EGet(e *dto.ShopGetReq) *entity.Shop {
	return &entity.Shop{
		Model: gorm.Model{ID: *e.ID},
	}
}

func (s *ShopReq) D2EOfType(e *dto.ShopOfTypeReq) *entity.ShowType {
	return &entity.ShowType{
		Model: gorm.Model{ID: e.TypeId},
	}
}
