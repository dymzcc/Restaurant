package controller

import (
	"Areyouanxiety/service"
	"Areyouanxiety/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct{

}

func (fcc *FoodCategoryController) Router(engine *gin.Engine){
	engine.GET("/api/food_category", fcc.foodCategory)
}

func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context){
	//调用service功能获取食品种类信息
	foodCategoryService := &service.FoodCategoryService{}
	category, err := foodCategoryService.Category()
	if err != nil{
		tool.Failed(ctx, "食品种类数据获取失败")
		return
	}
	//转换格式
	//imgUrl: hello.png
	for _,category := range category{
		if category.ImageUrl!=""{
			category.ImageUrl=tool.FileServerAddr()+"/"+category.ImageUrl
		}
	}
	tool.Success(ctx, category)
}