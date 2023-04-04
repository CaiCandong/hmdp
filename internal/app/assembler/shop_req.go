package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type ShopReq struct {
}

func (s *ShopReq) D2EGetRsp() *entity.Shop {
	return nil
}

func (s *ShopReq) D2EOfTypeRsp(e *dto.ShopOfTypeReq) *entity.ShowType {
	return &entity.ShowType{
		Model: gorm.Model{ID: e.TypeId},
	}
}
