package tool

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

type RedisDatabse struct {
	client *redis.Client
}

var Rdb RedisDatabse
var ctx = context.Background()

func InitRedisStore(cfg *Config) *RedisDatabse{
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	Rdb = RedisDatabse{client: client}

	base64Captcha.SetCustomStore(&Rdb)

	return &Rdb
}

//Set验证码
func (rd *RedisDatabse) Set(id string, value string)  {
	err := rd.client.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil{
		log.Println(err)
	}
}

//Get验证码
func (rd *RedisDatabse) Get(id string, clear bool) string{
	val, err := rd.client.Get(ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	return val
}

