package dto

import (
	"time"
)

type (
	VoucherCreateReq struct {
		ActualValue int       `json:"actual_value"`        // 抵扣金额，单位是分。例如200代表2元
		BeginTime   time.Time `json:"begin_time"`          // 生效时间
		EndTime     time.Time `json:"end_time"`            // 失效时间
		ID          int64     `json:"id"`                  // 主键
		PayValue    int       `json:"pay_value"`           // 支付金额，单位是分。例如200代表2元
		Rules       string    `json:"rules,omitempty"`     // 使用规则
		ShopID      int64     `json:"shop_id,omitempty"`   // 商铺id
		Status      int64     `json:"status"`              // 1,上架; 2,下架; 3,过期
		Stock       int       `json:"stock"`               // 库存
		SubTitle    string    `json:"sub_title,omitempty"` // 副标题
		Title       string    `json:"title"`               // 代金券标题
		Type        int64     `json:"type"`                // 0,普通券；1,秒杀券
	}

	VoucherListReq struct {
		ShopID uint `json:"shopId" uri:"shopId" binding:"required"`
	}
	VoucherSecKillReq struct {
		VoucherID uint `json:"id" uri:"id" binding:"required"`
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
	VoucherSecKillRsp struct {
		OrderId uint `json:"orderId"`
	}
	VoucherCreateRsp struct {
		Status int `json:"status"`
	}
)
