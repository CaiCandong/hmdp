package services

import (
	"context"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/repository"
)

type IShopService interface {
	OfType(ctx context.Context, req *dto.ShopOfTypeReq) (rsp []*dto.ShopOfTypeRsp, err error)
	GetById(ctx context.Context, req *dto.ShopGetReq) (rsp *dto.ShopGetRsp, err error)
}

type ShopService struct {
	ShopRepo repository.IShopRepo
	ShopReq  assembler.ShopReq
	ShopRsp  assembler.ShopRsp
}

func (s *ShopService) OfType(ctx context.Context, req *dto.ShopOfTypeReq) (rsp []*dto.ShopOfTypeRsp, err error) {
	shops, err := s.ShopRepo.GetShopByType(ctx, req.TypeId, req.Current)
	if err != nil {
		return nil, err
	}
	return s.ShopRsp.E2DOfType(shops), nil
}

func (s *ShopService) GetById(ctx context.Context, req *dto.ShopGetReq) (rsp *dto.ShopGetRsp, err error) {
	shop := s.ShopReq.D2EGet(req)
	err = s.ShopRepo.GetShopById(ctx, shop)
	if err != nil {
		return nil, err
	}
	return s.ShopRsp.E2DGetShop(shop), nil
}
