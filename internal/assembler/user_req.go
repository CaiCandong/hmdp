package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type UserReq struct {
}

func NewUserReq() *UserReq {
	return &UserReq{}
}

func (req *UserReq) D2ESendCode(d *dto.UserSendCodeReq) *model.User {
	return &model.User{Phone: d.Phone}
}

func (req *UserReq) D2ELoginByCode(d *dto.UserLoginByCodeReq) *model.User {
	return &model.User{Phone: d.Phone}
}

func (req *UserReq) D2EMe(d *dto.UserMeReq) *model.User {
	return &model.User{}
}

func (req *UserReq) D2EInfo(d *dto.UserInfoReq) *model.User {
	return &model.User{
		Model: gorm.Model{ID: d.ID},
	}
}
