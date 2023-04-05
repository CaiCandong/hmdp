package mysql

import (
	"context"
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
)

type Voucher struct {
	DB *gorm.DB
}

func (v *Voucher) GetByShopId(ctx context.Context, shopId uint) ([]*entity.Voucher, error) {
	var ret []*entity.Voucher
	err := v.DB.Model(&entity.Voucher{}).Where("shop_id = ?", shopId).Find(&ret).Error
	return ret, err
}
