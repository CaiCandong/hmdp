package entity

import (
	"fmt"
	"gorm.io/gorm"
	"hmdp/pkg/utils"
)

type User struct {
	gorm.Model
	Phone    string `gorm:"column:phone;type:varchar(11);uniqueIndex:idx_phone;NOT NULL" `
	Password string `gorm:"column:password;type:varchar(128);"`
	NickName string `gorm:"column:nick_name;type:varchar(32);" json:"nick_name"`
	Icon     string `gorm:"column:icon;type:varchar(255)" json:"icon"`
}

func (u *User) TableName() string {
	return "tb_user"
}

func CreateDefaultUser(phone string) User {
	return User{
		Phone:    phone,
		NickName: fmt.Sprintf("user_%v", utils.RandStringBytes(10)),
	}
}
