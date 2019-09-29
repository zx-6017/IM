package config

import (
	"github.com/astaxie/beego/config"
	"log"
	"os"
)

var Conf config.Configer

func init(){
	env_file := ".env"

	if _,err := os.Stat(env_file);os.IsNotExist(err){
		log.Panicf("conf file [%s] not found!",env_file)
	}
	conf,err := config.NewConfig("ini",env_file)
	if err != nil{
		log.Panicf("parse config file [%s] failed, err [%s]",env_file,err.Error())
	}
	Conf = conf
	log.Println("init all config file success")
}