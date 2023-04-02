package services

import (
	"hmdp/internal/app/vo"
	"hmdp/internal/domain/repository"
	"hmdp/pkg/serializer"
)

type IShopTypeService interface {
	List() (serializer.Response, error)
}

type ShowTypeService struct {
	repo repository.IShowType
}

type ShopTypeConfiguration func(os *ShowTypeService) error

func NewShopTypeService(cfgs ...ShopTypeConfiguration) IShopTypeService {
	// Create the user service
	os := &ShowTypeService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil
		}
	}
	return os
}

func WithShopTypeRepo(repo repository.IShowType) ShopTypeConfiguration {
	return func(os *ShowTypeService) error {
		os.repo = repo
		return nil
	}
}

func (s ShowTypeService) List() (serializer.Response, error) {
	shoptypes, err := s.repo.GetShopTypeList()
	if err != nil {
		return serializer.Response{}, nil
	}
	return serializer.Response{Data: vo.BuildShowTypeVos(shoptypes)}, nil
}
