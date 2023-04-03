package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
	"strconv"
)

type BlogController struct {
	blogService services.IBlogService
}

func NewBlogController(blogService services.IBlogService) *BlogController {
	return &BlogController{blogService}

}

func (b *BlogController) Find(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	req := &dto.BloGetReq{Id: idx}
	rsp, err := b.blogService.Get(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

func (b *BlogController) Hot(ctx *gin.Context) {
	req := &dto.BlogHotReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := b.blogService.Hot(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}
