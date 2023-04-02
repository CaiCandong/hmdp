package repository

import "hmdp/internal/domain/entity"

type IBlogRepo interface {
	GetBlog(ID any) (entity.Blog, error)
	GetBlogs(page, pageSize int) ([]entity.Blog, error)
}
