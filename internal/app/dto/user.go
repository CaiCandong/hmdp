package dto

import "time"

type (
	UserSendCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
	}
	UserLoginByCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
		Code  string `json:"code" form:"code" binding:"required"`
	}
	UserMeReq struct {
	}
	UserInfoReq struct {
		ID uint `json:"id" uri:"id" binding:"required"`
	}
)

type (
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
)
