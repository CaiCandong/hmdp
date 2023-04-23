//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"hmdp/internal/controller"
	"hmdp/internal/model"
	"hmdp/internal/service"
)

func InitHandlers() *controller.Handler {
	wire.Build(
		controller.NewHandlers,
		InitShopHandler,
		InitUserHandler,
		InitBlogHandler,
		InitVoucherHandler,
		InitShopTypeHandler,
		InitDB,
	)
	return nil
}

func InitDB() *gorm.DB {
	wire.Build(model.GetDB)
	return nil
}

func InitUserHandler(db *gorm.DB) *controller.UserHandler {
	wire.Build(controller.NewUserHandler, service.NewUserService)
	return nil
}

func InitShopHandler(db *gorm.DB) *controller.ShopHandler {
	wire.Build(controller.NewShopHandler, service.NewShopService)
	return nil
}

func InitBlogHandler(db *gorm.DB) *controller.BlogHandler {
	wire.Build(controller.NewBlogHandler, service.NewBlogService)
	return nil
}

func InitVoucherHandler(db *gorm.DB) *controller.VoucherHandler {
	wire.Build(controller.NewVoucherHandler, service.NewVoucherService)
	return nil
}

func InitShopTypeHandler(db *gorm.DB) *controller.ShopTypeHandler {
	wire.Build(controller.NewShopTypeController, service.NewShowTypeService)
	return nil
}
