package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

// BlogReq 序列化器
type BlogReq struct {
}

func NewBlogReq() *BlogReq {
	return &BlogReq{}
}

func (b *BlogReq) D2EFindBlogById(d *dto.FindBlogByIdReq) *model.Blog {
	return &model.Blog{
		Model: gorm.Model{ID: uint(d.Id)},
	}
}

func (b *BlogReq) D2EGetLike(d *dto.BlogGetLikeReq) *model.Blog {
	return &model.Blog{
		Model: gorm.Model{ID: uint(d.Id)},
	}
}
