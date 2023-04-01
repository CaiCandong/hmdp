package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/pkg/serializer"
	"net/http"
)

type ShopTypeController interface {
	List(c *gin.Context)
}

func NewShowController() ShopTypeController {
	return &ShopTypeControllerImp{}
}

type ShopTypeControllerImp struct {
}

func (s *ShopTypeControllerImp) List(ctx *gin.Context) {
	list, err := mysql.GetShopTypeList()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, serializer.Response{
		Success: true,
		Data:    list,
	})
}
