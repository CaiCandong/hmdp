package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

type UserRsp struct {
}

func NewUserRsp() *UserRsp {
	return &UserRsp{}
}

func (req *UserRsp) E2DSendCode() *dto.UserSendCodeRsp {
	return &dto.UserSendCodeRsp{}
}

func (req *UserRsp) E2DLoginByCode(e *entity.User) *dto.UserLoginByCodeRsp {
	return &dto.UserLoginByCodeRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon}
}

func (req *UserRsp) E2DInfo(e *entity.User) *dto.UserInfoRsp {
	return &dto.UserInfoRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon}
}
