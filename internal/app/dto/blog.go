package dto

type (
	BloGetReq struct {
		Id int `json:"id" form:"id"`
	}
	BlogHotReq struct {
		Current int `json:"current" from:"current"`
	}
)

type (
	BlogGetRsp struct {
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
)
