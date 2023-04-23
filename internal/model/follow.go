package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserId   uint
	FollowId uint
}

func (f Follow) TableName() string {
	return "tb_follow"
}
