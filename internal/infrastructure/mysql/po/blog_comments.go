package entity

import "gorm.io/gorm"

type BlogComment struct {
	gorm.Model
	UserId   uint   `gorm:"user_id"`
	BlogId   uint   `gorm:"blog_id"`
	ParentId uint   `gorm:"parent_id"`
	AnswerId uint   `gorm:"answer_id"`
	Content  string `gorm:"content"`
	Status   int    `gorm:"status"`
}

func (b BlogComment) TableName() string {
	return "tb_blog_comment"
}
