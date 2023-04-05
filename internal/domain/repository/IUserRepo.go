package repository

import "hmdp/internal/domain/entity"

type IUserRepo interface {
	GetUser(*entity.User) error                            //根据主键获取用户
	GetUserByPhone(phone any) (entity.User, error)         //根据手机号获取用户
	GetUserOrCreate(*entity.User) error                    //根据手机号码获取/创建用户
	CreateUserWithPhone(phone string) (entity.User, error) //根据手机号码创建用户
	//IsCached(user *entity.User) bool                       //判断用户信息是否缓存
	//Cache(user *entity.User) bool                          //缓存用户信息
}
