package model

import "gorm.io/gorm"

type ShopType struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(32)" json:"name"`
	Icon string `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Sort uint   `gorm:"column:sort" json:"sort"`
}

func (st *ShopType) TableName() string {
	return "tb_shop_type"
}
