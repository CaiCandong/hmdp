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
func (s *ShopReq) D2EUpdate(e *dto.ShopUpdateReq) *entity.Shop {
	return &entity.Shop{
		Model:     gorm.Model{ID: e.ID},
		Area:      e.Area,
		OpenHours: e.OpenHours,
		Sold:      e.Sold,
		Address:   e.Address,
		AvgPrice:  e.AvgPrice,
		Score:     e.Score,
		Name:      e.Name,
		TypeId:    e.TypeId,
	}
}
