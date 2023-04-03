package services

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/repository"
)

type IBlogService interface {
	Get(ctx *gin.Context, req *dto.BloGetReq) (*dto.BlogGetRsp, error)
	Hot(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error)
}

func NewBlogService(blogRepo repository.IBlogRepo) IBlogService {
	return &BlogService{blogRepo, &assembler.BlogReq{}, &assembler.BlogRsp{}}
}

type BlogService struct {
	BlogRepo repository.IBlogRepo
	BlogReq  *assembler.BlogReq
	BlogRsp  *assembler.BlogRsp
}

func (b *BlogService) Get(ctx *gin.Context, req *dto.BloGetReq) (*dto.BlogGetRsp, error) {
	blog := b.BlogReq.D2EGet(req)
	err := b.BlogRepo.GetBlog(blog)
	if err != nil {
		return nil, err
	}
	return b.BlogRsp.E2DGet(blog), nil
}

func (b *BlogService) Hot(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error) {
	//blog := b.BlogReq.D2EGet(req)
	blogs, err := b.BlogRepo.GetBlogs(req.Current, 10)
	if err != nil {
		return nil, err
	}
	return b.BlogRsp.E2DHot(blogs), nil
}
