package controller

import (
	"Areyouanxiety/service"
	"Areyouanxiety/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GoodController struct {

}

func (gc *GoodController) Router(engine *gin.Engine){
	engine.GET("/api/foods", gc.getGoods)
}

//获取某个商户下面所包含的食品
func (gc *GoodController) getGoods (context *gin.Context){
	shopId, exist := context.GetQuery("shop_id")
	if !exist{
		tool.Failed(context, "请求参数错误，请重试")
		return
	}

	id, err := strconv.Atoi(shopId)
	if err != nil{
		tool.Failed(context, "请求参数错误，请重试")
		return
	}
	goods := new(service.GoodService).GetFoods(int64(id))
	if len(goods) == 0{
		tool.Failed(context, "未查询到相关数据")
	}
	tool.Success(context, goods)
}