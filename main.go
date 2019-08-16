package main

import (
	"IM/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main(){

	gin.SetMode(gin.DebugMode);

	service := gin.Default();
	service.Use(gin.Recovery());


	// Log to a file
	//gin.DisableConsoleColor()
	log_file,_ := os.Create("./storage/logs/zx.gin.com.log")
	gin.DefaultWriter = io.MultiWriter(log_file)

	//引入路由
	routers.Router(service);

	//服务启动
	service.Run(":8080");






}