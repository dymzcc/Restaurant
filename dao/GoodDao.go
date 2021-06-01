package dao

import (
	"Areyouanxiety/model"
	"Areyouanxiety/tool"
)

type GoodDao struct {

}

func (gd *GoodDao) QueryFoods(shop_id int64)([]model.Goods, error){
	var goods []model.Goods
	err := tool.DB.Where("shop_id=?", shop_id).Find(&goods).Error
	if err != nil{
		return nil, err
	}
	return goods, err
}