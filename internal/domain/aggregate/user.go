package aggregate

import (
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
	"hmdp/pkg/utils"
)

type IUserAggregate interface {
	SendCode(ctx context.Context, session sessions.Session, user *entity.User) error
	LoginByCode(ctx context.Context, session sessions.Session, user *entity.User, code string) error
	//Info(ctx context.Context, session sessions.Session, user *entity.User) error
}

func NewUserAggregate(userRepo repository.IUserRepo) IUserAggregate {
	return &UserAggregate{userRepo}
}

type UserAggregate struct {
	UserRepo repository.IUserRepo
}

func (agg *UserAggregate) SendCode(ctx context.Context, session sessions.Session, user *entity.User) error {
	// 校验手机号码
	if !user.VerifyMobileFormat() {
		return fmt.Errorf("%v:手机号码格式错误", user.Phone)
	}
	// 生成验证码
	code := utils.GenValidateCode(6)
	// 下发验证码
	if err := user.SendCode(code); err != nil {
		return err
	}
	// 保存验证码
	session.Set("code", code)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}

func (agg *UserAggregate) LoginByCode(ctx context.Context, session sessions.Session, user *entity.User, code string) error {
	//校验手机号码
	if !user.VerifyMobileFormat() {
		return fmt.Errorf("%v:手机号码格式错误", user.Phone)
	}
	// 校验验证码是否一致
	originalCode := session.Get("code")
	if code, ok := originalCode.(string); !ok || code != originalCode {
		return fmt.Errorf("%v:验证码错误", user.Phone)
	}
	// 数据库查询
	err := agg.UserRepo.GetUserOrCreate(user)
	if err != nil {
		return err
	}
	// 会话保持
	session.Set("user_id", user.ID)
	err = session.Save()
	if err != nil {
		return err
	}
	return nil
}

//func (agg *UserAggregate) Info(ctx context.Context, session sessions.Session, user *entity.User) error {
//	//TODO implement me
//	panic("implement me")
//}
