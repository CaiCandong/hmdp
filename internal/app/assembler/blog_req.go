package assembler

import (
	"gorm.io/gorm"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

// BlogReq 序列化器
type BlogReq struct {
}

func NewBlogReq() *BlogReq {
	return &BlogReq{}
}

func (b *BlogReq) D2EGet(d *dto.BloGetReq) *entity.Blog {
	return &entity.Blog{
		Model: gorm.Model{ID: uint(d.Id)},
	}
}
