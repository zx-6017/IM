package helpers

import (
	"IM/config"
	"github.com/gomodule/redigo/redis"
	"log"
)

var RedisPool *redis.Pool

func init(){
	initRedis()
}

func initRedis(){
	host := config.Conf.String("REDIS_HOST")
	port := config.Conf.String("REDIS_PORT")
	pass := config.Conf.String("REDIS_PASSWORD")


	pool := &redis.Pool{
		MaxIdle:         16, //最大空闲连接数
		MaxActive:       0, //和数据库的最大连接数(最大并发) 0：无限制
		IdleTimeout:     120, //最大空闲时间
		Wait:            true,
		Dial: func() (conn redis.Conn, e error) {
			connect ,err := redis.Dial("tcp",host+":"+port,redis.DialPassword(pass))
			if err != nil{
				log.Fatalln("redis poll init failed...",err.Error())
			}
			return connect,err
		},
	}
	RedisPool = pool

}


