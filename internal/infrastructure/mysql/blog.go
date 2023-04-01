package mysql

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ShopId   int64  `gorm:"column:shop_id;type:bigint(20);comment:商户id;NOT NULL" json:"shopId"`
	UserId   uint64 `gorm:"column:user_id;type:bigint(20) unsigned;comment:用户id;NOT NULL" json:"userId"`
	Title    string `gorm:"column:title;type:varchar(255);comment:标题;NOT NULL" json:"title"`
	Images   string `gorm:"column:images;type:varchar(2048)" json:"images"`
	Content  string `gorm:"column:content;type:varchar(2048);comment:探店的文字描述;NOT NULL" json:"content"`
	Liked    uint   `gorm:"column:liked;type:int(8) unsigned;default:0;comment:点赞数量" json:"liked"`
	Comments uint   `gorm:"column:comments;type:int(8) unsigned;comment:评论数量" json:"comments"`
}

func (m *Blog) TableName() string {
	return "tb_blog"
}

// GetBlog 用ID获取博客
func GetBlog(ID any) (Blog, error) {
	var blog Blog
	result := DB.First(&blog, ID)
	return blog, result.Error
}

func GetBlogs(page, pageSize int) ([]Blog, error) {
	// 使用Gorm的Offset和Limit函数进行分页
	var blogs []Blog
	err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&blogs).Error

	return blogs, err
}
