package repository

import "hmdp/internal/domain/entity"

type IBlogRepo interface {
	GetBlog(blog *entity.Blog) error
	GetBlogById(blog *entity.Blog) error
	GetBlogs(page, pageSize int) ([]*entity.Blog, error)
	GetBlogByUserId(userId uint) ([]*entity.Blog, error)
}
