package tool

import (
	"Areyouanxiety/model"
	"github.com/jinzhu/gorm"
)

//声明一个全局变量
var (DB *gorm.DB)

func InitMySQL(cfg *Config)(err error){

	//拼接dsn参数
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port +
		")/" + cfg.Database.DbName + "?charset=utf8&parseTime=True&loc=Local&timeout=" + cfg.Database.Timeout
	DB, err = gorm.Open(cfg.Database.Driver, dsn)
	if err != nil{
		return
	}

	//模型绑定
	DB.AutoMigrate(&model.Member{},&model.FoodCategory{}, &model.Shop{}, &model.Service{}, &model.ShopService{}, &model.Goods{})

	//测试连通性
	err = DB.DB().Ping()
	return
}