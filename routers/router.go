package routers

import (
	"IM/controllers"
	"IM/helpers"
	"IM/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Router(router *gin.Engine) {

	//web 页面 ico
	router.GET("favicon.ico", func(request *gin.Context) {
		request.Status(200)
	})

	//平常测试接口
	test := router.Group("/test").Use(middlewares.Request())
	{
		test.GET("/redis", func(request *gin.Context) {

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
		test.GET("/mail", func(request *gin.Context) {

			mailTo := []string{
				"zhangxiao@wukongtv.com",
			}
			fmt.Println(mailTo[2])
			//subject := "Gin Im err..."
			//body := "err"
			//helpers.SendMailSmtp(mailTo,subject,body)
		})
		test.GET("/create",controllers.Create)
		test.GET("/friendrelation",controllers.FriendRelation)
		test.GET("/loginfo",controllers.GetLogInfo)
	}



	//// 测试 聊天室
	//chat := router.Group("/chat").Use(middlewares.Request())
	//{
	//	chat.GET("/server",server.Server)
	//	chat.GET("/client",client.Client)
	//}


	//yome 测试
	yome := router.Group("/api").Use(middlewares.Request())
	{
		yome.POST("/tool/imcallback",controllers.ImCallBack)

	}

}
