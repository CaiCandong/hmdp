package controller

type Handlers struct {
	UserHandler        *UserHandler
	ShopHandler        *ShopHandler
	BlogHandler        *BlogController
	VoucherHandler     *VoucherHandler
	ShopTypeController *ShopTypeController
}

func NewHandlers(u *UserHandler, s *ShopHandler, b *BlogController, v *VoucherHandler, st *ShopTypeController) *Handlers {
	return &Handlers{u, s, b, v, st}
}
