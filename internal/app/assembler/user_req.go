package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type UserReq struct {
}

func NewUserReq() *UserReq {
	return &UserReq{}
}

func (req *UserReq) D2ESendCode(d *dto.UserSendCodeReq) *entity.User {
	return &entity.User{Phone: d.Phone}
}

func (req *UserReq) D2ELoginByCode(d *dto.UserLoginByCodeReq) *entity.User {
	return &entity.User{Phone: d.Phone}
}

func (req *UserReq) D2EInfo(d *dto.UserInfoReq) *entity.User {
	return &entity.User{}
}
