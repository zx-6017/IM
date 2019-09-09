package helpers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func init(){
	initDB()
}
func initDB(){
	//初始化 数据库
	var db *gorm.DB

	var(
		dialect = "mysql"
		host = "localhost"
		port = "3306"
		username = "root"
		passwd = "mysqlpasswd"
		database = "yome"
		charset = "utf8"
	)

	dsn := username+":"+passwd+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local"
	db,mysql_err := gorm.Open(dialect,dsn)

	if mysql_err != nil {
		//logs.Error("init DB connect failed, error: %s",mysql_err.Error())
		log.Fatalln("init DB connect failed, error: %s",mysql_err.Error())
	}
	mysql_err = db.DB().Ping()
	if mysql_err != nil{
		//logs.Error("init DB connect ping failed, error: %s",mysql_err.Error())
		log.Fatalln("init DB connect failed, error: %s",mysql_err.Error())
	}else{
		DB = db
		//logs.Info("init DB connect success")
		log.Println("init DB connect success")
	}

	db.SingularTable(true)
	//defer db.Close()



}