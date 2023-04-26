package dto

import "time"

type (
	FindUserByIdReq struct {
		ID uint `json:"id" form:"id" uri:"id"`
	}
	IsFollowedReq struct {
		CurrentUserId uint //当前登录用户id
		FollowUserId  uint `json:"id" uri:"id" binding:"required"` // 被关注的用户id
	}
	UserSendCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
	}
	LoginByCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
		Code  string `json:"code" form:"code" binding:"required"`
	}
	UserMeReq struct {
	}
	UserInfoReq struct {
		ID uint `json:"id" uri:"id" binding:"required"`
	}
	// FollowUserReq /:id/:follow
	FollowUserReq struct {
		UserId uint  `json:"userId"`
		ID     uint  `json:"id" uri:"id" binding:"required"`         // 被关注的用户id
		Follow *bool `json:"follow" uri:"follow" binding:"required"` // true:关注 false:取消关注
	}
)

type (
	FindUserByIdRsp struct {
		ID       uint   `json:"id"`
		Icon     string `json:"icon"`
		NickName string `json:"nickName"`
	}
	UserSendCodeRsp struct {
	}
	UserLoginByCodeRsp struct {
		ID       uint   `json:"id"`
		NickName string `json:"nickName"`
		Icon     string `json:"icon"`
		Token    string `json:"token"`
	}
	UserMeRsp struct {
		ID       uint   `json:"id"`
		NickName string `json:"nickName"`
		Icon     string `json:"icon"`
	}
	UserInfoRsp struct {
		ID        uint      `json:"id"`
		NickName  string    `json:"nickName"`
		Icon      string    `json:"icon"`
		City      string    `json:"city"`
		Introduce string    `json:"introduce"`
		Fans      int       `json:"fans"`
		Followee  int       `json:"followee"`
		Gender    string    `json:"gender"`
		BirthDay  time.Time `json:"birth_day"`
		Credits   int       `json:"credits"`
		Level     bool      `json:"level"`
	}
	IsFollowedRsp struct {
		Followed bool `json:"followed"` // true:已关注 false:未关注
	}
	FollowUserRsp struct {
	}
)
