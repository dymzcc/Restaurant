package service

import (
	"Areyouanxiety/dao"
	"Areyouanxiety/model"
	"strconv"
)

type ShopService struct {

}

//查询商户列表数据
func (ss *ShopService) ShopList(long, lat string) []model.Shop{
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil{
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil{
		return nil
	}
	return new(dao.ShopDao).QueryShops(longitude, latitude, "")
}

//根据关键字搜索商铺信息
func (shopService *ShopService) SearchShops(long, lat, keyword string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	return new(dao.ShopDao).QueryShops(longitude, latitude, keyword)
}

//根据商户id获取相应服务
func (shopService *ShopService) GetService(shopId int64) []model.Service {
	return new(dao.ShopDao).QueryServiceByShopId(shopId)
}
