package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/dto"
	"hmdp/internal/model"
)

type ShopTypeService struct {
	DB  *gorm.DB
	Req *assembler.ShopTypeReq // 输入序列化器
	Rsp *assembler.ShopTypeRsp // 输入序列化器
}

func NewShowTypeService(db *gorm.DB) *ShopTypeService {
	return &ShopTypeService{
		DB:  db,
		Req: &assembler.ShopTypeReq{},
		Rsp: &assembler.ShopTypeRsp{},
	}
}

// ListShopTypes 列出所有商店类型
func (s *ShopTypeService) ListShopTypes(ctx *gin.Context, req *dto.ShopTypeListReq) (interface{}, error) {
	var shopTypes []*model.ShopType
	err := s.DB.Find(&shopTypes).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DListShopTypes(shopTypes), err
}
