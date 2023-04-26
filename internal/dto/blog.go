package dto

import "time"

type (
	FindBlogByIdReq struct {
		Id int `json:"id" form:"id"`
	}
	BlogHotReq struct {
		Current int `json:"current" from:"current"`
	}
	BlogGetLikeReq struct {
		Id int `json:"id" form:"id" uri:"id"`
	}
	ListBlogsByUserIdReq struct {
		UserId  uint `json:"id" form:"id"`           // 用户id
		Current int  `json:"current" form:"current"` // 页号
	}
	UploadBlogImgReq struct {
		UserId   uint   `json:"userId"`
		Filename string `json:"filename"`
	}
	CreateBlogReq struct {
	}
	LikeBlogByUserIdReq struct {
		UserId uint `json:"userId"`
		BlogId uint `json:"blogId" uri:"blogId" binding:"required"`
	}
	ListLikedUsersByBlogIdReq struct {
		BlogId uint `json:"blogId" uri:"blogId" binding:"required"`
	}
	IsBlogLikeReq struct {
	}
	LikeBlogReq struct {
	}
	ListBlogLikesReq struct {
	}
)

type (
	FindBlogByIdRsp struct {
		Id         uint      `json:"id"`         // 博客id
		ShopId     int64     `json:"shopId"`     // 店铺id
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
	BlogHotRsp struct {
		Id       uint   `json:"id"`
		ShopId   int64  `json:"shopId"`
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
	BlogGetLikeRsp struct {
		UserId   uint   `json:"id"`
		UserIcon string `json:"icon"`
	}
	BlogGetByUseIdRsp struct {
		BlogId   uint   `json:"id"`
		Images   string `json:"images"`
		Title    string `json:"title"`
		Likes    uint   `json:"likes"`
		Comments uint   `json:"comments"`
		IsLike   bool   `json:"isLike"`
	}
	CreateBlogRsp struct {
	}
	IsBlogLikeRsp struct {
	}
	LikeBlogRsp struct {
		Id   uint `json:"id"`   //用户id
		Icon uint `json:"icon"` //用户头像
	}

	ListLikedUsersByBlogIdRsp struct {
		Id   uint   `json:"id"`   //用户id
		Icon string `json:"icon"` //用户头像
	}
	LikeBlogByUserIdRsp struct {
	}
)
