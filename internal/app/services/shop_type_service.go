package services

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/repository"
)

type IShopTypeService interface {
	List(ctx *gin.Context, req *dto.ShopTypeListReq) ([]*dto.ShopTypeListRsp, error)
}

type ShowTypeService struct {
	ShopTypeRepo repository.IShopType
	ShopTypeReq  *assembler.ShopTypeReq
	ShopTypeRsp  *assembler.ShopTypeRsp
}

func NewShowTypeService(ShopTypeRepo repository.IShopType) IShopTypeService {
	return &ShowTypeService{
		ShopTypeRepo,
		&assembler.ShopTypeReq{},
		&assembler.ShopTypeRsp{}}
}

type ShopTypeConfiguration func(os *ShowTypeService) error

func (s *ShowTypeService) List(ctx *gin.Context, req *dto.ShopTypeListReq) ([]*dto.ShopTypeListRsp, error) {
	list, err := s.ShopTypeRepo.GetShopTypeList()
	if err != nil {
		return nil, err
	}
	return s.ShopTypeRsp.E2DShopTypeInfo(list), nil
}
