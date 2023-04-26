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

func (rsp *UserRsp) E2DSendCode() *dto.UserSendCodeRsp {
	return &dto.UserSendCodeRsp{}
}

func (rsp *UserRsp) E2DLoginByCode(e *model.User) *dto.UserLoginByCodeRsp {
	return &dto.UserLoginByCodeRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon, Token: e.Token}
}

func (rsp *UserRsp) E2DFindUserById(e *model.User) *dto.FindUserByIdRsp {
	return &dto.FindUserByIdRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon}
}

func (rsp *UserRsp) E2DIsFollowed(e *model.Follow) *dto.IsFollowedRsp {
	if e == nil {
		return &dto.IsFollowedRsp{Followed: false}
	}
	return &dto.IsFollowedRsp{Followed: true}
}

func (rsp *UserRsp) E2DMe(e *model.User) *dto.UserMeRsp {
	return &dto.UserMeRsp{ID: e.ID, NickName: e.NickName, Icon: e.Icon}
}

func (rsp *UserRsp) E2DInfo(e *model.UserInfo) *dto.UserInfoRsp {
	// TODO: Not Implement
	return &dto.UserInfoRsp{
		ID:        e.ID,
		NickName:  e.User.NickName,
		Icon:      e.User.Icon,
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
