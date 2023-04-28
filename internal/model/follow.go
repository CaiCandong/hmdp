package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserId   uint
	FollowId uint
}

func (f *Follow) TableName() string {
	return "tb_follow"
}

func FollowRedisKey(userId uint) string {
	return fmt.Sprintf("follow:%v", userId)
}
