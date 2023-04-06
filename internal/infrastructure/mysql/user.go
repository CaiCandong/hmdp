package mysql

import (
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
	"hmdp/pkg/utils"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(DB *gorm.DB) repository.IUserRepo {
	return &UserRepo{DB}
}

// GetUser 用ID获取用户
func (repo *UserRepo) GetUser(user *entity.User) error {
	return repo.db.First(&user).Error
}

// GetUserByPhone 用Phone获取用户
func (repo *UserRepo) GetUserByPhone(phone any) (entity.User, error) {
	var user entity.User
	result := repo.db.Where("phone = ?", phone).First(&user)
	return user, result.Error
}

func (repo *UserRepo) CreateUserWithPhone(phone string) (entity.User, error) {
	var user entity.User
	user.Phone = phone
	user.NickName = "user_" + utils.RandStringBytes(10)
	repo.db.Create(&user)
	return user, nil
}

func (repo *UserRepo) GetUserOrCreate(user *entity.User) error {
	//user := entity.CreateDefaultUser(phone)
	return repo.db.Where(entity.User{Phone: user.Phone}).FirstOrCreate(&user).Error

}
