package user_blog_agg

import (
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

// 用户博客聚合

// UserBlogAgg 聚合根：User
type UserBlogAgg struct {
	User     *entity.User
	UserInfo *entity.UserInfo // 一个用户只存在一个详细信息
	Blogs    []*entity.Blog   //一个用户存在多个博文
	Follow   []*entity.Follow // 一个用户存在多个粉丝
	UserRepo repository.IUserRepo
}

// UesrBasicInfo   返回用户的基本信息
//func (u *UserBlogAgg) UesrBasicInfo() (*entity.User, error) {
//	// 用户信息是否存在缓存中
//	if u.UserRepo.IsCached(u.User) {
//		return u.User, nil
//	}
//	// 从数据库中读取用户信息
//	if err := u.UserRepo.GetUser(u.User); err != nil {
//		// 将用户信息保存到cache中
//		u.UserRepo.Cache(u.User)
//	}
//	// 查看用户的信息
//	return u.User, nil
//}
//
//// UesrInfo  返回用户的详细
//func (u *UserBlogAgg) UesrInfo() (*entity.User, error) {
//	// 用户信息是否存在缓存中
//	if u.UserRepo.IsCached(u.User) {
//		return u.User, nil
//	}
//	// 从数据库中读取用户信息
//	if err := u.UserRepo.GetUser(u.User); err != nil {
//		// 将用户信息保存到cache中
//		u.UserRepo.Cache(u.User)
//	}
//	// 查看用户的信息
//	return u.User, nil
//}
