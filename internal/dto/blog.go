package dto

import "time"

type (
	FindBlogByIdReq struct {
		Id int `json:"id" form:"id"`
	}
	FindBlogByIdRsp struct {
		Id         uint      `json:"id"`         // 博客id
		ShopId     uint      `json:"shopId"`     // 店铺id
		UserId     uint      `json:"userId"`     // 发布博客的用户id
		UserIcon   string    `json:"icon"`       // 发布博客的用户头像
		UserName   string    `json:"name"`       // 发布博客的用户名
		Title      string    `json:"title"`      // 博客标题
		Images     string    `json:"images"`     // 博客图片
		Content    string    `json:"content"`    // 博客内容
		Liked      uint      `json:"liked"`      // 当前博客的点赞数量
		IsLike     bool      `json:"isLike"`     // 当前用户是否点赞
		Comments   uint      `json:"comments"`   // 博客评论数
		CreateTime time.Time `json:"createTime"` // 博客发布时间
	}
)

type (
	BlogHotReq struct {
		Current int `json:"current" from:"current"`
	}
	BlogHotRsp struct {
		Id       uint   `json:"id"`
		ShopId   uint   `json:"shopId"`
		UserId   uint   `json:"userId"`
		UserIcon string `json:"icon"`
		UserName string `json:"name"` //发布博客的用户名
		Title    string `json:"title"`
		Images   string `json:"images"`
		Content  string `json:"content"`
		Liked    uint   `json:"liked"`
		IsLike   bool   `json:"isLike"`
		Comments uint   `json:"comments"`
	}
)

type (
	CreateBlogReq struct {
		ShopId  uint   `json:"shopId"`
		UserId  uint   `json:"userId"`
		Content string `json:"content" binding:"required"`
		Images  string `json:"images" binding:"required"`
		Title   string `json:"title" binding:"required"`
	}
	CreateBlogRsp struct {
	}
)

type (
	BlogGetLikeReq struct {
		Id int `json:"id" form:"id" uri:"id"`
	}
	BlogGetLikeRsp struct {
		UserId   uint   `json:"id"`
		UserIcon string `json:"icon"`
	}
)

type (
	ListBlogsByUserIdReq struct {
		UserId  uint `json:"id" form:"id"`           // 用户id
		Current int  `json:"current" form:"current"` // 页号
	}
	ListBlogsByUserIdRsp struct {
		BlogId   uint   `json:"id"`
		Images   string `json:"images"`
		Title    string `json:"title"`
		Likes    uint   `json:"likes"`
		Comments uint   `json:"comments"`
		IsLike   bool   `json:"isLike"`
	}
)

type (
	ListBlogsBySubscriptionReq struct {
		Offset   int64  `json:"offset" form:"offset"`                    // 偏移量
		UserId   uint   `json:"id" form:"id"`                            // 用户id
		LastTime string `json:"lastId" form:"lastId" binding:"required"` // 最后一条博客的id(时间戳)
	}
	ListBlogsBySubscriptionRspBlog struct {
		Images string `json:"images"`
		Title  string `json:"title"`
		Likes  uint   `json:"likes"`
	}
	ListBlogsBySubscriptionRsp struct {
		Blogs    []*ListBlogsBySubscriptionRspBlog `json:"list"`
		LastTime float64                           `json:"minTime"`
		Offset   int64                             `json:"offset"`
	}
)

type (
	UploadBlogImgReq struct {
		UserId   uint   `json:"userId"`
		Filename string `json:"filename"`
	}
	UploadBlogImgRsp struct {
		Url string `json:"url"`
	}
)

type (
	LikeBlogByUserIdReq struct {
		UserId string `json:"userId"`
		BlogId uint   `json:"blogId" uri:"blogId" binding:"required"`
	}
	LikeBlogByUserIdRsp struct {
	}
)

type (
	ListLikedUsersByBlogIdReq struct {
		BlogId uint `json:"blogId" uri:"blogId" binding:"required"`
	}
	ListLikedUsersByBlogIdRsp struct {
		Id   uint   `json:"id"`   //用户id
		Icon string `json:"icon"` //用户头像
	}
)

type (
	IsBlogLikeReq struct {
	}
	IsBlogLikeRsp struct {
	}
)

type (
	LikeBlogReq struct {
	}
	LikeBlogRsp struct {
		Id   uint `json:"id"`   //用户id
		Icon uint `json:"icon"` //用户头像
	}
)
type (
	DeleteBlogImgReq struct {
		Url string `json:"name" form:"name"`
	}
	DeleteBlogImgRsp struct {
	}
)

type (
	ListBlogLikesReq struct {
	}
)
