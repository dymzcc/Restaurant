package tool
import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(cfg *Config, engine *gin.Engine){
	store, err := redis.NewStore(10, "tcp", cfg.Redis.Addr+":"+cfg.Redis.Port, "dede", []byte("secret"))
	if err != nil{
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

func Setsession(context *gin.Context, key interface{}, value interface{}) error{
	session := sessions.Default(context)
	if session == nil{
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

func Getsession(context *gin.Context, key interface{}) interface{}{
	session := sessions.Default(context)
	return session.Get(key)
}