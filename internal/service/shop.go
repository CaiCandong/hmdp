package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type ShopService struct {
	DB  *gorm.DB
	Req *assembler.ShopReq // 输入序列化器
	Rsp *assembler.ShopRsp // 输入序列化器
}

func NewShopService(db *gorm.DB) *ShopService {
	return &ShopService{
		DB:  db,
		Req: &assembler.ShopReq{},
		Rsp: &assembler.ShopRsp{},
	}
}

// FindShopById 根据ID查找商店
func (s *ShopService) FindShopById(ctx *gin.Context, req *dto.FindShopByIdReq) (*dto.FindShopByIdRsp, error) {
	var shop model.Shop
	err := s.DB.First(&shop, req.ID).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DFindShopById(&shop), err
}

// ListShopsByType 根据类型查找商店
func (s *ShopService) ListShopsByType(ctx *gin.Context, req *dto.ListShopsByTypeReq) ([]*dto.ListShopsByTypeRsp, error) {
	var shops []*model.Shop
	page := req.Current
	//pageSize := req.PageSize
	pageSize := 5
	err := s.DB.Where("type_id = ?", req.TypeId).Offset((page - 1) * pageSize).Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DListShopsByType(shops), err
}

// ListShopsByName 根据名称查找商店
func (s *ShopService) ListShopsByName(ctx *gin.Context, req *dto.ListShopsByNameReq) ([]*dto.ListShopsByNameRsp, error) {
	var shops []*model.Shop
	err := s.DB.Where("name like ?", "%"+req.Name+"%").Find(&shops).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DListShopsByName(shops), err
}

// UpdateShopById 更新商店信息
func (s *ShopService) UpdateShopById(ctx *gin.Context, req *dto.UpdateShopByIdReq) (*dto.UpdateShopByIdRsp, error) {
	var shop model.Shop
	err := s.DB.First(&shop, req.ID).Error
	if err != nil {
		return nil, err
	}
	err = s.DB.Model(&shop).Updates(s.Req.D2EUpdate(req)).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DUpdateShopById(&shop), err
}
