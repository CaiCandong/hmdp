package dto

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
		UserId uint `json:"userId"`
	}
)

type (
	FindBlogByIdRsp struct {
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
	}
)
