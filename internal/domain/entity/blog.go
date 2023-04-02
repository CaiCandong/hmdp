package entity

import "gorm.io/gorm"

//type Blog struct {
//	gorm.Model
//	Id       uint   `json:"id"`
//	ShopId   int64  `json:"shopId"`
//	UserId   uint64 `json:"userId"`
//	UserIcon string `json:"icon"`
//	UserName string `json:"name"` //发布博客的用户名
//	Title    string `json:"title"`
//	Images   string `json:"images"`
//	Content  string `json:"content"`
//	Liked    uint   `json:"liked"`
//	Comments uint   `json:"comments"`
//}

type Blog struct {
	gorm.Model
	ShopId   int64  `gorm:"column:shop_id;type:bigint(20);comment:商户id;NOT NULL" json:"shopId"`
	UserId   uint64 `gorm:"column:user_id;type:bigint(20) unsigned;comment:用户id;NOT NULL" json:"userId"`
	Title    string `gorm:"column:title;type:varchar(255);comment:标题;NOT NULL" json:"title"`
	Images   string `gorm:"column:images;type:varchar(2048)" json:"images"`
	Content  string `gorm:"column:content;type:varchar(2048);comment:探店的文字描述;NOT NULL" json:"content"`
	Liked    uint   `gorm:"column:liked;type:int(8) unsigned;default:0;comment:点赞数量" json:"liked"`
	Comments uint   `gorm:"column:comments;type:int(8) unsigned;comment:评论数量" json:"comments"`
}

func (m *Blog) TableName() string {
	return "tb_blog"
}
