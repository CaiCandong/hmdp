package assembler

import (
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/entity"
)

// BlogRsp 序列化器
type BlogRsp struct {
}

func NewBlogRsp() *BlogRsp {
	return &BlogRsp{}
}

// E2DGet entity转换成dto
func (s *BlogRsp) E2DGet(blog *entity.Blog) *dto.BlogGetRsp {
	return &dto.BlogGetRsp{
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

// E2DHot entity转换成dto
func (s *BlogRsp) E2DHot(blogs []*entity.Blog) []*dto.BlogHotRsp {
	blogDTOs := make([]*dto.BlogHotRsp, len(blogs))
	for i, b := range blogs {
		blogDTOs[i] = &dto.BlogHotRsp{
			Id: b.ID,
			//ShopId:   b.ShopId,
			UserId:   b.UserId,
			UserIcon: b.User.Icon,
			UserName: b.User.NickName,
			Title:    b.Title,
			Images:   b.Images,
			Content:  b.Content,
			Liked:    b.Liked,
			Comments: b.Comments,
		}
	}
	return blogDTOs
}

// E2DGetLike entity转换成dto
func (s *BlogRsp) E2DGetLike(users []*entity.User) []*dto.BlogGetLikeRsp {
	DTOs := make([]*dto.BlogGetLikeRsp, len(users))
	for i, b := range users {
		DTOs[i] = &dto.BlogGetLikeRsp{
			UserId:   b.ID,
			UserIcon: b.Icon,
		}
	}
	return DTOs
}

// E2DGetBlogById entity转换成dto
func (s *BlogRsp) E2DGetBlogById(blogs []*entity.Blog) []*dto.BlogGetByUseIdRsp {
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
