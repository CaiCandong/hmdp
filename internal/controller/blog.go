package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"hmdp/internal/serializer"
	"hmdp/internal/service"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	svc *service.BlogService
}

func NewBlogHandler(svc *service.BlogService) *BlogHandler {
	return &BlogHandler{svc}
}

// FindBlogById 根据博客id获取博客
func (c *BlogHandler) FindBlogById(ctx *gin.Context) {
	req := &dto.FindBlogByIdReq{}
	id := ctx.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	req.Id = Id
	rsp, err := c.svc.FindBlogById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return

	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

// ListHotBlogs 获取博客列表
func (c *BlogHandler) ListHotBlogs(ctx *gin.Context) {
	req := &dto.BlogHotReq{}
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := c.svc.ListHotBlogs(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

// ListBlogsByUserId 根据用户id获取
func (c *BlogHandler) ListBlogsByUserId(ctx *gin.Context) {
	req := &dto.ListBlogsByUserIdReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	blogs, err := c.svc.ListBlogByUserId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(blogs))
}

// UploadBlogImg 上传博客图片
func (c *BlogHandler) UploadBlogImg(ctx *gin.Context) {
	req := &dto.UploadBlogImgReq{}
	// 获取用户id
	req.UserId = ctx.MustGet("user").(*model.User).ID
	// 获取原始文件名称
	// 生成新文件名
	// 保存文件
	// 返回结果
	rsp, err := c.svc.UploadBlogImg(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))

	//req.UserId
}

// CreateBlog 保存博客
func (c *BlogHandler) CreateBlog(ctx *gin.Context) {
	req := &dto.CreateBlogReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user := ctx.MustGet("user").(*model.User)
	req.UserId = user.ID
	rsp, err := c.svc.CreateBlog(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
}

func (c *BlogHandler) LikeBlogByUserId(ctx *gin.Context) {
	req := &dto.LikeBlogByUserIdReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user := ctx.MustGet("user").(*model.User)
	req.UserId = strconv.Itoa(int(user.ID))
	rsp, err := c.svc.LikeBlogByUserId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

// ListLikedUsersByBlogId 获取点赞用户列表
func (c *BlogHandler) ListLikedUsersByBlogId(ctx *gin.Context) {
	req := &dto.ListLikedUsersByBlogIdReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := c.svc.ListLikedUsersByBlogId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

// DeleteBlogImg 删除博客图片
func (c *BlogHandler) DeleteBlogImg(ctx *gin.Context) {
	req := &dto.DeleteBlogImgReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	rsp, err := c.svc.DeleteBlogImg(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (c *BlogHandler) ListBlogsOfMe(ctx *gin.Context) {
	req := &dto.ListBlogsByUserIdReq{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user := ctx.MustGet("user").(*model.User)
	req.UserId = user.ID
	blogs, err := c.svc.ListBlogByUserId(ctx, req)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(blogs))
}

func (c *BlogHandler) ListBlogsBySubscription(context *gin.Context) {
	req := &dto.ListBlogsBySubscriptionReq{}
	if err := context.ShouldBind(req); err != nil {
		context.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user := context.MustGet("user").(*model.User)
	req.UserId = user.ID
	rsp, err := c.svc.ListBlogsBySubscription(context, req)
	if err != nil {
		context.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	context.JSON(http.StatusOK, serializer.Success(rsp))
	return
}
