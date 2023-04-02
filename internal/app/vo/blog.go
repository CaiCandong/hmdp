package vo

import (
	"hmdp/internal/domain/entity"
	"hmdp/pkg/serializer"
)

type BlogVO struct {
	Id       uint   `json:"id"`
	ShopId   int64  `json:"shopId"`
	UserId   uint   `json:"userId"`
	UserIcon string `json:"icon"`
	UserName string `json:"name"` //发布博客的用户名
	Title    string `json:"title"`
	Images   string `json:"images"`
	Content  string `json:"content"`
	Liked    uint   `json:"liked"`
	Comments uint   `json:"comments"`
}

func buildBlog(blog entity.Blog) BlogVO {
	return BlogVO{
		Id: blog.ID,
		//ShopId:   blog.ShopId,
		UserId:   blog.UserId,
		UserIcon: blog.User.Icon,
		UserName: blog.User.NickName,
		Title:    blog.Title,
		Images:   blog.Images,
		Content:  blog.Content,
		Liked:    blog.Liked,
		Comments: blog.Comments,
	}
}

func buildBlogs(blogs []entity.Blog) (blogDTOs []BlogVO) {
	blogDTOs = make([]BlogVO, len(blogs))
	for i := range blogs {
		blogDTOs[i] = buildBlog(blogs[i])
	}
	return
}

func BuildBlogsResponse(blogs []entity.Blog) serializer.Response {
	return serializer.Response{
		Success: true,
		Data:    buildBlogs(blogs),
	}
}
