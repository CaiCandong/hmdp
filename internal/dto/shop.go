package dto

type (
	FindShopByIdReq struct {
		ID *uint `json:"id" uri:"id" binding:"required"`
	}

	ListShopsByTypeReq struct {
		TypeId  uint    `json:"typeId" form:"typeId" binding:"required"`
		Current int     `json:"current" form:"current"` //分页信息 单页多少?
		X       float32 `json:"x" form:"x"`
		Y       float32 `json:"y" form:"y"`
	}

	UpdateShopByIdReq struct {
		Area      string `json:"area"`      //店铺区域
		OpenHours string `json:"openHours"` //店铺营业时间
		Sold      uint32 `json:"sold"`      //店铺销量
		Address   string `json:"address"`   //店铺地址
		AvgPrice  uint64 `json:"avgPrice"`  //店铺人均价格
		Score     uint32 `json:"score"`     //店铺评分
		Name      string `json:"name"`      //店铺名称
		TypeId    uint64 `json:"typeId"`    //店铺类型ID
		ID        uint   `json:"id"`        //店铺ID
	}
	ListShopsByNameReq struct { //查询所有店铺
		Name string `json:"name" form:"name"`
	}
)

type (
	FindShopByIdRsp struct {
		Name     string `json:"name"`
		AvgPrice uint64 `json:"avgPrice"`
		Images   string `json:"images"`
		Score    uint32 `json:"score"`
		Comments uint32 `json:"comments"`
		Address  string `json:"address"`
	}
	ListShopsByTypeRsp struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Score    uint32 `json:"score"`
		Comments uint32 `json:"comments"`
		Images   string `json:"images"`
		Area     string `json:"area"`
		Distance int    `json:"distance"`
		AvgPrice uint64 `json:"avgPrice"`
		Address  string `json:"address"`
	}
	UpdateShopByIdRsp struct {
		Success bool `json:"success"` //是否成功
	}
	ListShopsByNameRsp struct {
		ShopId uint   `json:"id"`
		Name   string `json:"name"`
		Area   string `json:"area"`
	}
)

//INSERT INTO `tb_shop` VALUES (1, '103茶餐厅', 1, 'https://qcloud.dpfile.com/pc/jiclIsCKmOI2arxKN1Uf0Hx3PucIJH8q0QSz-Z8llzcN56-_QiKuOvyio1OOxsRtFoXqu0G3iT2T27qat3WhLVEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vfCF2ubeXzk49OsGrXt_KYDCngOyCwZK-s3fqawWswzk.jpg,https://qcloud.dpfile.com/pc/IOf6VX3qaBgFXFVgp75w-KKJmWZjFc8GXDU8g9bQC6YGCpAmG00QbfT4vCCBj7njuzFvxlbkWx5uwqY2qcjixFEuLYk00OmSS1IdNpm8K8sG4JN9RIm2mTKcbLtc2o2vmIU_8ZGOT1OjpJmLxG6urQ.jpg', '大关', '金华路锦昌文华苑29号', 120.149192, 30.316078, 80, 0000004215, 0000003035, 37, '10:00-22:00', '2021-12-22 18:10:39', '2022-01-13 17:32:19');
