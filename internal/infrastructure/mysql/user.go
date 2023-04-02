package mysql

import (
	"gorm.io/gorm"
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
	var user entity.User
	err := DB.Where("phone = ?", phone).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user = entity.CreateDefaultUser(phone)
		err = DB.Create(&user).Error
		if err != nil {
			return user, err
		}
	} else if err != nil {
		// some other error occurred
		return user, err
	}
	return user, nil
}
