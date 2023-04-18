package entity

import (
	"fmt"
	"gorm.io/gorm"
	"hmdp/pkg/logger"
	"hmdp/pkg/utils"
	"regexp"
)

type User struct {
	gorm.Model
	Phone    string `gorm:"column:phone;type:varchar(11);uniqueIndex:idx_phone;NOT NULL" `
	Password string `gorm:"column:password;type:varchar(128);"`
	NickName string `gorm:"column:nick_name;type:varchar(32);" json:"nick_name"`
	Icon     string `gorm:"column:icon;type:varchar(255)" json:"icon"`
	Token    string `gorm:"-"`
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

func (u *User) VerifyMobileFormat() bool {
	// 校验手机号码格式
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(string(u.Phone))
}

func (u *User) SendCode(code string) error {
	//TODO:调用运营商SDK下发验证码
	logger.Logger.Info(fmt.Sprintf("phone:%v;code:%v", u.Phone, code))
	return nil
}

func (u *User) GenToken() string {
	u.Token = utils.UUID()
	return u.Token
}
