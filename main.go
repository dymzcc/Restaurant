package main

import (
	"Areyouanxiety/controller"
	"Areyouanxiety/middleware"
	"Areyouanxiety/tool"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	//解析配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil{
		panic(err.Error())
	}

	//实例化数据库
	tool.InitMySQL(cfg)
	//初始化Redis
	tool.InitRedisStore(cfg)

	engine := gin.Default()

	//使用中间件
	registerMiddleWare(engine)
	//集成session
	tool.InitSession(cfg, engine)
	//注册路由
	registerRouter(engine)

	engine.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(engine *gin.Engine){
	new(controller.HelloController).Router(engine)
	new(controller.MemberController).Router(engine)
	new(controller.FoodCategoryController).Router(engine)
	new(controller.ShopController).Router(engine)
	new(controller.GoodController).Router(engine)
}

//使用中间件
func registerMiddleWare(engine *gin.Engine){
	//设置全局跨域访问
	engine.Use(middleware.Cors())
}



