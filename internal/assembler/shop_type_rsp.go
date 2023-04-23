package assembler

import (
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type ShopTypeRsp struct {
}

func NewShopTypeRsp() *ShopTypeRsp {
	return &ShopTypeRsp{}
}

func (rsp *ShopTypeRsp) E2DListShopTypes(es []*model.ShopType) []*dto.ShopTypeListRsp {
	ret := make([]*dto.ShopTypeListRsp, len(es))
	for i, e := range es {
		ret[i] = &dto.ShopTypeListRsp{ID: e.ID, Name: e.Name, Icon: e.Icon, Sort: e.Sort}
	}
	return ret
}
