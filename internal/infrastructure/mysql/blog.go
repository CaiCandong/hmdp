package mysql

import (
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type BlogRepo struct {
	//DB *gorm.DB
}

func NewBlogRepo() repository.IBlogRepo {
	return &BlogRepo{}
}

// GetBlog 用ID获取博客
func (b *BlogRepo) GetBlog(ID any) (entity.Blog, error) {
	var blog entity.Blog
	result := DB.First(&blog, ID)
	return blog, result.Error
}

func (b *BlogRepo) GetBlogs(page, pageSize int) ([]entity.Blog, error) {
	// 使用Gorm的Offset和Limit函数进行分页
	var blogs []entity.Blog
	err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&blogs).Error
	return blogs, err
}
