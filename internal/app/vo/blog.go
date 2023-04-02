package vo

import (
	"hmdp/internal/domain/entity"
	"hmdp/pkg/serializer"
)

type BlogVO struct {
	Id       uint   `json:"id"`
	ShopId   int64  `json:"shopId"`
	UserId   uint64 `json:"userId"`
	UserIcon string `json:"icon"`
	UserName string `json:"name"` //发布博客的用户名
	Title    string `json:"title"`
	Images   string `json:"images"`
	Content  string `json:"content"`
	Liked    uint   `json:"liked"`
	Comments uint   `json:"comments"`
}

func buildBlog(blog entity.Blog, user entity.User) BlogVO {
	return BlogVO{
		Id:       blog.ID,
		ShopId:   blog.ShopId,
		UserId:   blog.UserId,
		UserIcon: user.Icon,
		UserName: user.NickName,
		Title:    blog.Title,
		Images:   blog.Images,
		Content:  blog.Content,
		Liked:    blog.Liked,
		Comments: blog.Comments,
	}
}

func buildBlogs(blogs []entity.Blog, users []entity.User) (blogDTOs []BlogVO) {
	blogDTOs = make([]BlogVO, len(blogs))
	for i := range blogs {
		blogDTOs[i] = buildBlog(blogs[i], users[i])
	}
	return
}

func BuildBlogsResponse(blogs []entity.Blog, users []entity.User) serializer.Response {
	return serializer.Response{
		Success: true,
		Data:    buildBlogs(blogs, users),
	}
}
