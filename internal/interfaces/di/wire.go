//go:build wireinject
// +build wireinject

package di

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"hmdp/internal/app/services"
	"hmdp/internal/domain/aggregate"
	"hmdp/internal/domain/repository"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/internal/interfaces/controller"
)

func InitUserRepo(db *gorm.DB) repository.IUserRepo {
	wire.Build(mysql.NewUserRepo)
	return nil
}

func InitShopHandler(db *gorm.DB) *controller.ShopHandler {
	wire.Build(controller.NewShopHandler, services.NewShopService, mysql.NewShopRepo)
	return nil
}
func InitUserHandler(db *gorm.DB) *controller.UserHandler {
	wire.Build(controller.NewUserHandler, services.NewUserService, aggregate.NewUserAggregate, mysql.NewUserRepo)
	return nil
}
func InitBlogHandler(db *gorm.DB) *controller.BlogController {
	wire.Build(controller.NewBlogController, services.NewBlogService, mysql.NewBlogRepo)
	return nil
}

func InitVoucherHandler(db *gorm.DB) *controller.VoucherHandler {
	wire.Build(controller.NewVoucherHandler, services.NewVoucherService, mysql.NewVoucherRepo)
	return nil
}

func InitShopTypeHandler(db *gorm.DB) *controller.ShopTypeController {
	wire.Build(controller.NewShopTypeController, services.NewShowTypeService, mysql.NewShopTypeRepo)
	return nil
}
