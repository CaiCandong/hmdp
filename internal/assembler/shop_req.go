package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type ShopReq struct {
}

func (s *ShopReq) D2EGet(e *dto.FindShopByIdReq) *model.Shop {
	return &model.Shop{
		Model: gorm.Model{ID: *e.ID},
	}
}

func (s *ShopReq) D2EOfType(e *dto.ListShopsByTypeReq) *model.ShopType {
	return &model.ShopType{
		Model: gorm.Model{ID: e.TypeId},
	}
}
func (s *ShopReq) D2EUpdate(e *dto.UpdateShopByIdReq) *model.Shop {
	return &model.Shop{
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
