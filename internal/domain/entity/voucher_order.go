package entity

import (
	"gorm.io/gorm"
	"time"
)

type VoucherOrder struct {
	gorm.Model
	UserId     uint
	VoucherId  uint
	PayType    int
	Status     int
	PayTime    time.Time
	UseTime    time.Time
	RefundTime time.Time
}

func (v VoucherOrder) TableName() string {
	return "tb_voucher_order"
}
