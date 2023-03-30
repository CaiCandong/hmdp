package services

import "github.com/gin-gonic/gin"

type User interface {
	SendCode(c *gin.Context)
	Login(c *gin.Context)
}
