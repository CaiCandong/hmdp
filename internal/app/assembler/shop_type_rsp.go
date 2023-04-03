package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type ShopTypeRsp struct {
}

func NewShopTypeRsp() *ShopTypeRsp {
	return &ShopTypeRsp{}
}

func (rsp *ShopTypeRsp) E2DShopTypeInfo(es []*entity.ShowType) []*dto.ShopTypeListRsp {
	ret := make([]*dto.ShopTypeListRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.ShopTypeListRsp{ID: e.ID, Name: e.Name, Icon: e.Icon, Sort: e.Sort}
	}
	return ret
}
