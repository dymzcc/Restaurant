package model

type FoodCategory struct {
	//类别ID
	Id int64 `json:"id"`
	//食品类别标题
	Title string `json:"title"`
	//食品描述
	Description string `json:"description"`
	//食品种类图片
	ImageUrl string `json:"image_url"`
	//食品类别连接
	LinkUrl string `json:"link_url"`
	//该类别是否在服务状态
	IsInserving bool `json:"is_in_serving"`
}
