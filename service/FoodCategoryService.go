package service

import (
	"Areyouanxiety/dao"
	"Areyouanxiety/model"
)

type FoodCategoryService struct {

}

func (fcs *FoodCategoryService) Category()([]model.FoodCategory, error){
	return new(dao.FoodCategoryDao).QueryCategory()
}