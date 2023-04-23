package model

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	ShopId      uint
	Title       string
	SubTitle    string
	Rules       string
	PayValue    int
	ActualValue int
	Type        int
	Status      int
	Stock       int
	//BeginTime time.Time
	//EndTime time.Time
}

func (v Voucher) TableName() string {
	return "tb_voucher"
}
