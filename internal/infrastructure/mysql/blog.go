package mysql

import (
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type BlogRepo struct {
	db *gorm.DB
}

// GetBlog 用ID获取博客
func (repo *BlogRepo) GetBlog(blog *entity.Blog) error {
	return repo.db.First(&blog, blog.ID).Error
}

func NewBlogRepo(DB *gorm.DB) repository.IBlogRepo {
	return &BlogRepo{DB}
}

// GetBlogById 用ID获取博客
func (repo *BlogRepo) GetBlogById(blog *entity.Blog) error {
	return repo.db.Where("id = ?", blog.ID).First(blog).Error
}

func (repo *BlogRepo) GetBlogs(page, pageSize int) ([]*entity.Blog, error) {
	// 使用Gorm的Offset和Limit函数进行分页
	var blogs []*entity.Blog
	//err := DB.Model(&entity.Blog{}).Find(&blogs).Error
	err := repo.db.Model(entity.Blog{}).
		Preload("User").Offset((page - 1) * pageSize).
		Limit(pageSize).Find(&blogs).Error

	//err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&blogs).Error
	return blogs, err
}

func (repo *BlogRepo) GetBlogByUserId(userId uint) ([]*entity.Blog, error) {
	var blogs []*entity.Blog
	err := repo.db.Where("user_id = ?", userId).Find(&blogs).Error
	return blogs, err
}
