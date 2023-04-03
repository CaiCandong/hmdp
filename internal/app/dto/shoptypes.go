package dto

type (
	ShopTypeListReq struct {
	}
)

type (
	ShopTypeListRsp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
		Sort uint   `json:"sort"`
	}
)

//func BuildShowTypeVo(showType entity.ShowType) ShowTypeVo {
//	return ShowTypeVo{
//		ID:   showType.ID,
//		Name: showType.Name,
//		Icon: showType.Icon,
//		Sort: showType.Sort,
//	}
//}

//func BuildShowTypeVos(showTypes []entity.ShowType) []ShowTypeVo {
//	ret := make([]ShowTypeVo, len(showTypes))
//	for i := range ret {
//		ret[i] = BuildShowTypeVo(showTypes[i])
//	}
//	return ret
//}
