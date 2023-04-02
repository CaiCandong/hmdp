package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/services"
	"hmdp/internal/infrastructure/mysql"
	"net/http"
)

type BlogController interface {
	Find(c *gin.Context)
	Hot(c *gin.Context)
}

func NewBlogController() BlogController {
	// 进行依赖注入
	blogRepo := mysql.NewBlogRepo()
	blogService := services.NewBlogService(services.WithBlogRepo(blogRepo))
	return BlogControllerImp{blogService}
}

type BlogControllerImp struct {
	blogService services.IBlogService
}

func (b BlogControllerImp) Find(ctx *gin.Context) {
	// 根据id获取博客详情
	id := ctx.Param("id")
	resp, err := b.blogService.GetBlog(id)
	if err != nil {
		resp.Success = false
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Success = true
	ctx.JSON(http.StatusOK, resp)
}

func (b BlogControllerImp) Hot(ctx *gin.Context) {
	var req struct {
		Current int `json:"current" form:"current"`
	}
	// 参数绑定
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	resp, err := b.blogService.Hot(req.Current, 10)
	if err != nil {
		resp.Success = false
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Success = true
	ctx.JSON(http.StatusOK, resp)
}
