package routers

import (
	"IM/controllers"
	"IM/helpers"
	"IM/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Router(router *gin.Engine){

	router.GET("favicon.ico", func(request *gin.Context) {
		request.Status(200)
	})

	api := router.Group("/api").Use(middlewares.Request())
	{
		api.GET("/create",controllers.Create)
		api.GET("/friendrelation",controllers.FriendRelation)
		api.GET("/loginfo",controllers.GetLogInfo)
		api.GET("/test", func(request *gin.Context) {

			redisCon := helpers.RedisPool.Get()

			_,err := redisCon.Do("set","name","zx")
			if err != nil{
				fmt.Println("redis set err...",err.Error())
				return
			}
			name,err := redis.String(redisCon.Do("get","name"))

			if err != nil{
				fmt.Println("redis get err...",err.Error())
				return
			}
			request.JSON(200,gin.H{
				"name":name,
			})
		})
	}


}
