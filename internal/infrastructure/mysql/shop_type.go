package mysql

import "gorm.io/gorm"

type ShowType struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(32)" json:"name"`
	Icon string `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Sort uint   `gorm:"column:sort" json:"sort"`
}

func (st *ShowType) TableName() string {
	return "tb_shop_type"
}

func GetShopTypeList() ([]ShowType, error) {
	var showtypes []ShowType
	result := DB.Find(&showtypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return showtypes, nil
}
