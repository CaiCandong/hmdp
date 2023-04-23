package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/service"
	"hmdp/pkg/serializer"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	svc *service.BlogService
}

func NewBlogHandler(svc *service.BlogService) *BlogHandler {
	return &BlogHandler{svc}
}

func (c *BlogHandler) FindBlogById(ctx *gin.Context) {
	req := &dto.FindBlogByIdReq{}
	var svc service.BlogService
	id := ctx.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	req.Id = Id
	rsp, err := svc.FindBlogById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

func (c *BlogHandler) BlogHot(ctx *gin.Context) {
	var svc service.BlogService
	req := &dto.BlogHotReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := svc.ListBlogs(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

// ListBlogsByUserId 根据用户id获取
func (c *BlogHandler) ListBlogsByUserId(ctx *gin.Context) {
	req := &dto.ListBlogsByUserIdReq{}
	session := sessions.Default(ctx)
	req.UserId = session.Get("user_id").(uint)
	blogs, err := c.svc.ListBlogByUserId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(blogs))
}
