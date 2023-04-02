package vo

import (
	"hmdp/internal/domain/entity"
	"hmdp/pkg/serializer"
)

type UserDTO struct {
	ID       uint   `json:"id"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
}

func BuildUserResponse(user entity.User) serializer.Response {
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
