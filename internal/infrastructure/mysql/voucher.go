package mysql

import (
	"context"
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

type Voucher struct {
	db *gorm.DB
}

func NewVoucherRepo(db *gorm.DB) repository.IVoucherRepo {
	return &Voucher{db: db}
}
func (v *Voucher) GetByShopId(ctx context.Context, shopId uint) ([]*entity.Voucher, error) {
	var ret []*entity.Voucher
	err := v.db.Model(&entity.Voucher{}).Where("shop_id = ?", shopId).Find(&ret).Error
	return ret, err
}
