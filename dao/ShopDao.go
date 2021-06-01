package dao

import (
	"Areyouanxiety/model"
	"Areyouanxiety/tool"
)

type ShopDao struct {

}

//查询经度,维度相差DEFAULT_RANGE之内的的商家
const DEFAULT_RANGE = 5
func (shopDao *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	if keyword == "" {
		err := tool.DB.Where("longitude> ? and longitude < ? and latitude > ? and latitude < ?",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops).Error
		if err != nil {
			return nil
		}
	} else {
		err := tool.DB.Where("longitude> ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status=1",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, "%"+keyword+"%").Find(&shops).Error
		if err != nil {
			return nil
		}
	}
	return shops
}

//根据商户id查询对应的服务
func (shopDao *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {
	var service []model.Service
	err := tool.DB.Joins("INNER JOIN service.id=shop_service.service_id and shop_service.shop_id=?", shopId).Find(&service).Error
	if err != nil {
		return nil
	}
	return service
}