package main

import (
	"IM/config"
	"IM/helpers"
	"IM/routers"
	"github.com/astaxie/beego/logs"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"syscall"
)

func main(){


	gin.SetMode(gin.DebugMode)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	//引入路由
	routers.Router(engine)


	err := logs.SetLogger(logs.AdapterFile,`{"filename":"`+config.Conf.String("LOG_FILE")+`","daily":true,"color":true}`)

	if err != nil {
		log.Fatalln("Log init failed, err:"+err.Error())
	}
	//logs.EnableFuncCallDepth(false) //输出调用的文件名 行号

	server := endless.NewServer(config.Conf.String("APP_PORT"),engine)

	signals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	}
	for _,signal := range signals{
		_ = server.RegisterSignalHook(endless.POST_SIGNAL,signal,helpers.SignalExe)
	}

	//把pid写入文件
	server.BeforeBegin = func(add string) {
		helpers.WritePid()
	}



	//服务启动
	err = server.ListenAndServe()
	if err !=nil{
		log.Fatalf("Server start failed, error: " + err.Error())
	}

	log.Println("Server start success...")



	//退出前操作

	// 删除pid文件
	defer helpers.DelPid()

	// 删除DB Redis连接
	defer helpers.DB.Close()
	defer helpers.RedisPool.Close()
}