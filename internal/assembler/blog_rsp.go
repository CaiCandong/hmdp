package assembler

import (
	"context"
	"hmdp/internal/cache"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

// BlogRsp 序列化器
type BlogRsp struct {
}

func NewBlogRsp() *BlogRsp {
	return &BlogRsp{}
}

// E2DFindBlogById model转换成dto
func (s *BlogRsp) E2DFindBlogById(blog *model.Blog, user *model.User) *dto.FindBlogByIdRsp {
	ret := &dto.FindBlogByIdRsp{
		Id:         blog.ID,
		ShopId:     blog.ShopId,
		UserId:     blog.UserId,
		UserIcon:   blog.User.Icon,
		UserName:   blog.User.NickName,
		Title:      blog.Title,
		Images:     blog.Images,
		Content:    blog.Content,
		Liked:      blog.Liked,
		Comments:   blog.Comments,
		CreateTime: blog.CreatedAt,
	}
	if user != nil {
		ret.IsLike = cache.IsLike(context.Background(), blog.ID, user.ID)
	}
	return ret
}

// E2DHot model转换成dto
func (s *BlogRsp) E2DHot(blogs []*model.Blog, user *model.User) []*dto.BlogHotRsp {
	blogDTOs := make([]*dto.BlogHotRsp, len(blogs))
	for i, b := range blogs {
		blogDTOs[i] = &dto.BlogHotRsp{
			Id:       b.ID,
			ShopId:   b.ShopId,
			UserId:   b.UserId,
			UserIcon: b.User.Icon,
			UserName: b.User.NickName,
			Title:    b.Title,
			Images:   b.Images,
			Content:  b.Content,
			Liked:    b.Liked,
			Comments: b.Comments,
		}
		if user != nil {
			blogDTOs[i].IsLike = cache.IsLike(context.Background(), blogs[i].ID, user.ID)
		}
	}
	return blogDTOs
}

// E2DGetLike model转换成dto
func (s *BlogRsp) E2DGetLike(users []*model.User) []*dto.BlogGetLikeRsp {
	DTOs := make([]*dto.BlogGetLikeRsp, len(users))
	for i, b := range users {
		DTOs[i] = &dto.BlogGetLikeRsp{
			UserId:   b.ID,
			UserIcon: b.Icon,
		}
	}
	return DTOs
}

// E2DListBlogsByUserId model转换成dto
func (s *BlogRsp) E2DListBlogsByUserId(blogs []*model.Blog) []*dto.BlogGetByUseIdRsp {
	DTOs := make([]*dto.BlogGetByUseIdRsp, len(blogs))
	for i, b := range blogs {
		DTOs[i] = &dto.BlogGetByUseIdRsp{
			BlogId:   b.ID,
			Images:   b.Images,
			Title:    b.Title,
			Likes:    b.Liked,
			Comments: b.Comments,
		}
	}
	return DTOs
}

func (s *BlogRsp) E2DListBlogLikes(users []*model.User) []*dto.ListLikedUsersByBlogIdRsp {
	ret := make([]*dto.ListLikedUsersByBlogIdRsp, len(users))
	for i, u := range users {
		ret[i] = &dto.ListLikedUsersByBlogIdRsp{
			Id:   u.ID,
			Icon: u.Icon,
		}
	}
	return ret
}
