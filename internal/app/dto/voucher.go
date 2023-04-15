package dto

import "time"

type (
	VoucherListReq struct {
		ShopID uint `json:"shopId" uri:"shopId" binding:"required"`
	}
)

type (
	VoucherListRsp struct {
		Id          uint      `json:"id"`
		Type        bool      `json:"type"`
		Title       string    `json:"title"`
		SubTitle    string    `json:"subTitle"`
		PayValue    int       `json:"payValue"`
		ActualValue int       `json:"actualValue"`
		Stock       int       `json:"stock"`
		BeginTime   time.Time `json:"beginTime"`
		EndTime     time.Time `json:"endTime"`
	}
)
