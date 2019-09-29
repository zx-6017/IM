package helpers

import (
	."IM/config"
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
		dialect = Conf.String("DB_DIALECT")
		host = Conf.String("DB_HOST")
		port = Conf.String("DB_PORT")
		username = Conf.String("DB_USERNAME")
		passwd = Conf.String("DB_PASSWORD")
		database = Conf.String("DB_DATABASE")
		charset = Conf.String("DB_CHARSET")
		parseTime = Conf.String("DB_PARSETIME")
		loc = Conf.String("DB_LOC")
	)
	dsn := username+":"+passwd+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc
	db,mysql_err := gorm.Open(dialect,dsn)

	if mysql_err != nil {
		log.Fatalln("init DB connect failed, error: %s",mysql_err.Error())
	}
	mysql_err = db.DB().Ping()
	if mysql_err != nil{
		log.Fatalln("init DB connect failed, error: %s",mysql_err.Error())
	}else{
		DB = db
		log.Println("init DB connect success")
	}
	db.LogMode(true)
	db.SingularTable(true)



}