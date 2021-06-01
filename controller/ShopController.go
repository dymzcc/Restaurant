package controller

import (
	"Areyouanxiety/service"
	"Areyouanxiety/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopController struct {

}

func (sc *ShopController) Router(engine *gin.Engine){
	engine.GET("/api/shops",sc.GetShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}

//根据搜索查询商户列表
func (sc *ShopController) SearchShop(context *gin.Context){
	//获取经纬度,如不存在就赋默认值
	longgitude := context.Query("longgitude")
	latitude := context.Query("latitude")
	keyword := context.DefaultQuery("keyword", "xxx")
	fmt.Println(keyword, longgitude, latitude)

	if keyword == ""{
		tool.Failed(context, "查询错误，请重新输入商铺名称")
		return
	}

	if longgitude == "" || longgitude == "undefined" || latitude == "" || latitude == "undefined"{
		longgitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.SearchShops(longgitude, latitude, keyword)
	if len(shops) != 0{
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "未搜索到相关商户")
}

//获取商户列表
func (sc *ShopController) GetShopList(context *gin.Context){
	//获取经纬度,如不存在就赋默认值
	longgitude := context.Query("longgitude")
	latitude := context.Query("latitude")

	if longgitude == "" || longgitude == "undefined" || latitude == "" || latitude == "undefined"{
		longgitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longgitude, latitude)
	if len(shops) == 0{
		tool.Failed(context, "未搜索到相关商户")
		return
	}
	//返回之前先用shopid查询其具有的service
	for _, shop := range shops {
		shopServices := shopService.GetService(shop.Id)
		if len(shopServices) == 0 {
			shop.Supports = nil
		} else {
			shop.Supports = shopServices
		}
	}
	tool.Success(context, shops)
}