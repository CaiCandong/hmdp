package domain

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/entity"
)

type UserRepository interface {
	Login(c *gin.Context, phone entity.Phone)
}
