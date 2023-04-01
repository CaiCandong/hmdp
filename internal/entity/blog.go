package entity

type Blog struct {
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
