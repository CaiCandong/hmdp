package controller

import (
	"github.com/gin-contrib/sessions"
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

// GetLikes 获取博客的点赞数量
func (b *BlogController) GetLikes(ctx *gin.Context) {

}

// GetBlogsByUserId 根据用户id获取
func (b *BlogController) GetBlogsByUserId(ctx *gin.Context) {
	req := &dto.BlogGetByUseIdReq{}
	session := sessions.Default(ctx)
	req.UserId = session.Get("user_id").(uint)
	blogs, err := b.blogService.GetBlogByUserId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(blogs))
}
