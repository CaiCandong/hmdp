package vo

import (
	"hmdp/internal/infrastructure/mysql"
	"hmdp/pkg/serializer"
)

type UserDTO struct {
	ID       uint   `json:"id"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
}

func BuildUserResponse(user *mysql.User) serializer.Response {
	dto := UserDTO{
		ID:       user.ID,
		NickName: user.NickName,
		Icon:     user.Icon,
	}

	return serializer.Response{
		Success: true,
		Data:    dto,
	}
}
