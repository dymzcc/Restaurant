package service

import (
	"Areyouanxiety/dao"
	"Areyouanxiety/model"
)

type GoodService struct {

}
//获取商家的食品信息列表
func (gs *GoodService) GetFoods(shop_id int64) []model.Goods{
	goods, err := new(dao.GoodDao).QueryFoods(shop_id)
	if err != nil{
		return nil
	}
	return goods
}