package services

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/repository"
	"hmdp/internal/infrastructure/cache"
)

type IShopTypeService interface {
	List(ctx *gin.Context, req *dto.ShopTypeListReq) ([]*dto.ShopTypeListRsp, error)
}

type ShowTypeService struct {
	ShopTypeRepo repository.IShopTypeRepo
	ShopTypeReq  *assembler.ShopTypeReq
	ShopTypeRsp  *assembler.ShopTypeRsp
}

func NewShowTypeService(ShopTypeRepo repository.IShopTypeRepo) IShopTypeService {
	return &ShowTypeService{
		ShopTypeRepo,
		&assembler.ShopTypeReq{},
		&assembler.ShopTypeRsp{}}
}

type ShopTypeConfiguration func(os *ShowTypeService) error

func (s *ShowTypeService) List(ctx *gin.Context, req *dto.ShopTypeListReq) ([]*dto.ShopTypeListRsp, error) {
	list, err := cache.GetShopType(ctx)
	if err == nil {
		return s.ShopTypeRsp.E2DShopTypeInfo(list), nil
	}
	list, err = s.ShopTypeRepo.GetShopTypeList()
	if err != nil {
		return nil, err
	}
	err = cache.SaveShopType(ctx, list)
	return s.ShopTypeRsp.E2DShopTypeInfo(list), err
}
