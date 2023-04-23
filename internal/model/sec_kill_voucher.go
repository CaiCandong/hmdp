package model

import (
	"gorm.io/gorm"
	"time"
)

type SeckillVoucher struct {
	gorm.Model
	VoucherId uint
	Stock     int
	BeginTime time.Time
	EndTime   time.Time
}

func (s SeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
