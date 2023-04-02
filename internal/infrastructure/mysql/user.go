package mysql

import (
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
	"hmdp/pkg/utils"
)

type UserRepo struct {
	//DB *gorm.DB
}

func NewUserRepo() repository.IUserRepo {
	return &UserRepo{}
}

// GetUser 用ID获取用户
func (repo *UserRepo) GetUser(ID any) (entity.User, error) {
	var user entity.User
	result := DB.First(&user, ID)
	return user, result.Error
}

// GetUserByPhone 用Phone获取用户
func (repo *UserRepo) GetUserByPhone(phone any) (entity.User, error) {
	var user entity.User
	result := DB.Where("phone = ?", phone).First(&user)
	return user, result.Error
}

func (repo *UserRepo) CreateUserWithPhone(phone string) (entity.User, error) {
	var user entity.User
	user.Phone = phone
	user.NickName = "user_" + utils.RandStringBytes(10)
	DB.Create(&user)
	return user, nil
}

func (repo *UserRepo) GetUserOrCreate(phone string) (entity.User, error) {
	user := entity.CreateDefaultUser(phone)
	err := DB.Where(entity.User{Phone: phone}).FirstOrCreate(&user).Error
	return user, err
}
