package entity

import (
	"gorm.io/gorm"
	"time"
)

type VoucherOrder struct {
	gorm.Model
	UserId     uint      `gorm:"user_id"`
	VoucherId  uint      `gorm:"voucher_id"`
	PayType    int       `gorm:"pay_type"`
	Status     int       `gorm:"status"`
	PayTime    time.Time `gorm:"pay_time"`
	UseTime    time.Time `gorm:"use_time"`
	RefundTime time.Time `gorm:"refund_time"`
}

func (v VoucherOrder) TableName() string {
	return "tb_voucher_order"
}
