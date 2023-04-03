package mysql

import (
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type BlogRepo struct {
	DB *gorm.DB
}

// GetBlog 用ID获取博客
func (b *BlogRepo) GetBlog(blog *entity.Blog) error {
	return DB.First(&blog, blog.ID).Error
}

func NewBlogRepo(DB *gorm.DB) repository.IBlogRepo {
	return &BlogRepo{DB}
}

// GetBlogById 用ID获取博客
func (b *BlogRepo) GetBlogById(blog *entity.Blog) error {
	return DB.Where("id = ?", blog.ID).First(blog).Error
}

func (b *BlogRepo) GetBlogs(page, pageSize int) ([]*entity.Blog, error) {
	// 使用Gorm的Offset和Limit函数进行分页
	var blogs []*entity.Blog
	//err := DB.Model(&entity.Blog{}).Find(&blogs).Error
	err := DB.Model(entity.Blog{}).
		Preload("User").Offset((page - 1) * pageSize).
		Limit(pageSize).Find(&blogs).Error

	//err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&blogs).Error
	return blogs, err
}
