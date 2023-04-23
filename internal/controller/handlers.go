package controller

type Handler struct {
	UserHandler        *UserHandler
	ShopHandler        *ShopHandler
	ShopTypeController *ShopTypeHandler
	BlogHandler        *BlogHandler
	VoucherHandler     *VoucherHandler
}

func NewHandlers(User *UserHandler, Shop *ShopHandler, ShopType *ShopTypeHandler, Blog *BlogHandler, Voucher *VoucherHandler) *Handler {
	return &Handler{
		UserHandler:        User,
		ShopHandler:        Shop,
		ShopTypeController: ShopType,
		BlogHandler:        Blog,
		VoucherHandler:     Voucher,
	}
}
