package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/cache"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"strings"
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
	// 计算分页参数
	page, pageSize := req.Current, viper.Get("page_size").(int)
	from, end := (page-1)*pageSize, page*pageSize
	// 判断是否需要根据坐标查询
	if req.X == 0 && req.Y == 0 {
		err := s.DB.Where("type_id = ?", req.TypeId).Offset((page - 1) * pageSize).Find(&shops).Error
		if err != nil {
			return nil, err
		}
		return s.Rsp.E2DListShopsByType(shops, nil), err
	}
	// 查询redis缓存 按照距离排序、分页
	key := model.ShopGeoKey(req.TypeId)
	// 获取附近的商店id
	result, err := cache.RedisStore.GeoRadius(ctx, key, req.X, req.Y, &redis.GeoRadiusQuery{
		Radius:    5 * 1000, // 5km
		Unit:      "m",
		WithCoord: true,
		WithDist:  true,
		Count:     end,
		Sort:      "ASC",
	}).Result()
	if err != nil || from >= len(result) {
		return nil, err
	}

	// 解析出id
	var ids []string
	var distances []float64
	result = result[from:]
	for _, item := range result {
		ids = append(ids, item.Name)
		distances = append(distances, item.Dist)
	}
	order := fmt.Sprintf("FIELD(id,%s)", strings.Join(ids, ","))
	err = s.DB.Where("id IN ?", ids).Order(order).Find(&shops).Error
	// 查询mysql
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DListShopsByType(shops, distances), err
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
