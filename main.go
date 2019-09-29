package main

import (
	"IM/config"
	"IM/routers"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){


	gin.SetMode(gin.DebugMode);

	service := gin.Default();
	service.Use(gin.Recovery());

	err := logs.SetLogger(logs.AdapterFile,`{"filename":"./storage/logs/im.log","daily":true,"color":true}`)

	if err != nil {
		log.Fatalln("Log init failed, err:"+err.Error())
	}
	//logs.EnableFuncCallDepth(false) //输出调用的文件名 行号

	//引入路由
	routers.Router(service)
	//服务启动
	service.Run(config.Conf.String("APP_PORT"))








}