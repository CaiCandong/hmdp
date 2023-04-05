package shop_voucher_agg

import "hmdp/internal/domain/entity"

type ShopVoucherAgg struct {
	Shop     *entity.Shop      //聚合根
	Vouchers []*entity.Voucher //一个用户存在多个消费券
	
}
