package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/vo"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/pkg/serializer"
	"net/http"
)

type BlogController interface {
	Find(c *gin.Context)
	Hot(c *gin.Context)
}

func NewBlogController() BlogController {
	return BlogControllerImp{}
}

type BlogControllerImp struct {
}

func (b BlogControllerImp) Find(ctx *gin.Context) {
	// 根据id获取博客详情
	id := ctx.Param("id")
	blog, err := mysql.GetBlog(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{Success: false})
		return
	}
	// TODO:查看当前用户是否点赞
	ctx.JSON(http.StatusOK, serializer.Response{Success: true, Data: blog})
}

func (b BlogControllerImp) Hot(ctx *gin.Context) {
	var req struct {
		Current int `json:"current" form:"current"`
	}
	// 参数绑定
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	blogs, err := mysql.GetBlogs(req.Current, 10)
	users := make([]mysql.User, len(blogs))
	for i := range blogs {
		users[i], _ = mysql.GetUser(blogs[i].UserId)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	ctx.JSON(http.StatusOK, vo.BuildBlogsResponse(blogs, users))
}
