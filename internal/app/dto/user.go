package dto

type (
	UserSendCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
	}
	UserLoginByCodeReq struct {
		Phone string `json:"phone" form:"phone" binding:"required"`
		Code  string `json:"code" form:"code" binding:"required"`
	}
	UserInfoReq struct {
	}
)

type (
	UserSendCodeRsp struct {
	}
	UserLoginByCodeRsp struct {
		ID       uint   `json:"id"`
		NickName string `json:"nickName"`
		Icon     string `json:"icon"`
	}
	UserInfoRsp struct {
		ID       uint   `json:"id"`
		NickName string `json:"nickName"`
		Icon     string `json:"icon"`
	}
)
