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
			//ShopId: b.ShopId,
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

//
//func (s *BlogReq) buildBlogs(blogs []entity.Blog) (blogDTOs []dto.BlogVO) {
//	blogDTOs = make([]dto.BlogVO, len(blogs))
//	for i := range blogs {
//		blogDTOs[i] = s.buildBlog(blogs[i])
//	}
//	return
//}
//
//func (s BlogReq) BuildBlogResponse(blog entity.Blog) Response {
//	return Response{
//		Success: true,
//		Data:    s.buildBlog(blog),
//	}
//}
//
//func (s BlogReq) BuildBlogsResponse(blogs []entity.Blog) Response {
//	return Response{
//		Success: true,
//		Data:    s.buildBlogs(blogs),
//	}
//}
