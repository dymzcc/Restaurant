package dao

import (
	"Areyouanxiety/model"
	"Areyouanxiety/tool"
	"fmt"
)

type FoodCategoryDao struct {

}

//从数据库中查询所有的视频种类
func (fcd * FoodCategoryDao) QueryCategory() ([]model.FoodCategory, error){
	var category []model.FoodCategory
	err := tool.DB.Find(&category).Error
	if err != nil{
		fmt.Println(err)
	}
	return category, nil
}