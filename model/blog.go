package model

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ShopId   int64  `gorm:"column:shop_id;type:bigint(20);comment:商户id;NOT NULL" json:"shop_id"`
	UserId   uint64 `gorm:"column:user_id;type:bigint(20) unsigned;comment:用户id;NOT NULL" json:"user_id"`
	Title    string `gorm:"column:title;type:varchar(255);comment:标题;NOT NULL" json:"title"`
	Images   string `gorm:"column:images;type:varchar(2048)"`
	Content  string `gorm:"column:content;type:varchar(2048);comment:探店的文字描述;NOT NULL" json:"content"`
	Liked    uint   `gorm:"column:liked;type:int(8) unsigned;default:0;comment:点赞数量" json:"liked"`
	Comments uint   `gorm:"column:comments;type:int(8) unsigned;comment:评论数量" json:"comments"`
}

func (m *Blog) TableName() string {
	return "tb_blog"
}
