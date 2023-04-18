package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type ShopRsp struct {
}

func (s *ShopRsp) E2DOfType(es []*entity.Shop) []*dto.ShopOfTypeRsp {
	ret := make([]*dto.ShopOfTypeRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.ShopOfTypeRsp{
			ID:       e.ID,
			Name:     e.Name,
			Score:    e.Score,
			Comments: e.Comments,
			Images:   e.Images,
			Area:     e.Area,
			Distance: 0,
			AvgPrice: e.AvgPrice,
			Address:  e.Address,
		}
	}
	return ret
}

func (s *ShopRsp) E2DGetShop(shop *entity.Shop) *dto.ShopGetRsp {
	return &dto.ShopGetRsp{
		Name:     shop.Name,
		AvgPrice: shop.AvgPrice,
		Images:   shop.Images,
		Score:    shop.Score,
		Comments: shop.Comments,
		Address:  shop.Address,
	}
}
func (s *ShopRsp) E2DUpdate(shop *entity.Shop) *dto.ShopUpdateRsp {
	return &dto.ShopUpdateRsp{Success: true}
}
