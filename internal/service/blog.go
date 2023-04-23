package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type BlogService struct {
	DB  *gorm.DB
	Req *assembler.BlogReq
	Rsp *assembler.BlogRsp
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{
		DB:  db,
		Req: &assembler.BlogReq{},
		Rsp: &assembler.BlogRsp{},
	}
}
func (svc *BlogService) FindBlogById(ctx *gin.Context, req *dto.FindBlogByIdReq) (*dto.FindBlogByIdRsp, error) {
	var blog model.Blog
	err := svc.DB.First(&blog, req.Id).Error
	if err != nil {
		return nil, err

	}
	return svc.Rsp.E2DFindBlogById(&blog), err
}

func (svc *BlogService) ListBlogs(ctx *gin.Context, req *dto.BlogHotReq) ([]*model.Blog, error) {
	var blogs []*model.Blog
	//page, pageSize := req.Page, req.PageSize
	// TODO:
	page, pageSize := req.Current, 5
	err := svc.DB.Model(model.Blog{}).Preload("User").Offset((page - 1) * pageSize).Limit(pageSize).Find(&blogs).Error
	return blogs, err
}

func (svc *BlogService) ListBlogByUserId(ctx *gin.Context, req *dto.ListBlogsByUserIdReq) ([]*model.Blog, error) {
	var blogs []*model.Blog
	//var page, pageSize int
	err := svc.DB.Where("user_id = ?", req.UserId).Find(&blogs).Error
	return blogs, err
}
