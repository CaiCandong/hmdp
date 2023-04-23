package assembler

import (
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"time"
)

type UserRsp struct {
}

func NewUserRsp() *UserRsp {
	return &UserRsp{}
}

func (req *UserRsp) E2DSendCode() *dto.UserSendCodeRsp {
	return &dto.UserSendCodeRsp{}
}

func (req *UserRsp) E2DLoginByCode(e *model.User) *dto.UserLoginByCodeRsp {
	return &dto.UserLoginByCodeRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon, Token: e.Token}
}

func (req *UserRsp) E2DMe(e *model.User) *dto.UserMeRsp {
	return &dto.UserMeRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon}
}

func (req *UserRsp) E2DInfo(e *model.User) *dto.UserInfoRsp {
	// TODO: Not Implement
	return &dto.UserInfoRsp{
		ID:        e.ID,
		NickName:  e.NickName,
		Icon:      e.Icon,
		City:      "北京",
		Introduce: "我爱学习",
		Fans:      100,
		Followee:  100,
		Gender:    "男",
		BirthDay:  time.Now(),
		Credits:   1,
		Level:     true,
	}
}