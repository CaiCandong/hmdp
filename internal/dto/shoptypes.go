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
