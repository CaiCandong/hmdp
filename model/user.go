package model

import (
	"gorm.io/gorm"
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

// GetUser 用ID获取用户
func GetUser(ID any) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// GetUserByPhone 用Phone获取用户
func GetUserByPhone(phone any) (User, error) {
	var user User
	result := DB.Where("phone = ?", phone).First(&user)
	return user, result.Error
}
