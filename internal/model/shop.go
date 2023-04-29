package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name      string  `gorm:"column:name;NOT NULL;comment:'商铺名称'"`
	TypeId    uint    `gorm:"column:type_id;NOT NULL;comment:'商铺类型的id'"`
	Images    string  `gorm:"column:images;NOT NULL;comment:'商铺图片，多个图片以\,\隔开'"`
	Area      string  `gorm:"column:area;default:NULL;comment:'商圈，例如陆家嘴'"`
	Address   string  `gorm:"column:address;NOT NULL;comment:'地址'"`
	X         float64 `gorm:"column:x;NOT NULL;comment:'经度'"`
	Y         float64 `gorm:"column:y;NOT NULL;comment:'维度'"`
	AvgPrice  uint64  `gorm:"column:avg_price;default:NULL;comment:'均价，取整数'"`
	Sold      uint32  `gorm:"column:sold;NOT NULL;comment:'销量'"`
	Comments  uint32  `gorm:"column:comments;NOT NULL;comment:'评论数量'"`
	Score     uint32  `gorm:"column:score;NOT NULL;comment:'评分，1~5分，乘10保存，避免小数'"`
	OpenHours string  `gorm:"column:open_hours;default:NULL;comment:'营业时间，例如 10:00-22:00'"`
}

func (t *Shop) TableName() string {
	return "tb_shop"
}

// ShopGeoKey 商铺的地理位置Redis的key
func ShopGeoKey(TypeId uint) string {
	return fmt.Sprintf("shop:geo:%d", TypeId)
}
