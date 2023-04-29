package assembler

import (
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type ShopRsp struct {
}

func (s *ShopRsp) E2DListShopsByType(es []*model.Shop, distances []float64) []*dto.ListShopsByTypeRsp {
	ret := make([]*dto.ListShopsByTypeRsp, len(es))
	if distances == nil {
		distances = make([]float64, len(es))
	}
	for i, e := range es {
		ret[i] = &dto.ListShopsByTypeRsp{
			ID:       e.ID,
			Name:     e.Name,
			Score:    e.Score,
			Comments: e.Comments,
			Images:   e.Images,
			Area:     e.Area,
			Distance: distances[i],
			AvgPrice: e.AvgPrice,
			Address:  e.Address,
		}
	}
	return ret
}

func (s *ShopRsp) E2DFindShopById(shop *model.Shop) *dto.FindShopByIdRsp {
	return &dto.FindShopByIdRsp{
		Name:     shop.Name,
		AvgPrice: shop.AvgPrice,
		Images:   shop.Images,
		Score:    shop.Score,
		Comments: shop.Comments,
		Address:  shop.Address,
	}
}

func (s *ShopRsp) E2DUpdateShopById(shop *model.Shop) *dto.UpdateShopByIdRsp {
	return &dto.UpdateShopByIdRsp{Success: true}
}

func (s *ShopRsp) E2DListShopsByName(es []*model.Shop) []*dto.ListShopsByNameRsp {
	ret := make([]*dto.ListShopsByNameRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.ListShopsByNameRsp{
			Name:   e.Name,
			Area:   e.Area,
			ShopId: e.ID,
		}
	}
	return ret
}
