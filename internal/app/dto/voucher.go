package dto

type (
	VoucherListReq struct {
		ShopID uint `json:"shopId" uri:"shopId" binding:"required"`
	}
)

type (
	VoucherListRsp struct {
		Title       string `json:"title"`
		SubTitle    string `json:"subTitle"`
		PayValue    int    `json:"payValue"`
		ActualValue int    `json:"actualValue"`
		Stock       int    `json:"stock"`
	}
)
