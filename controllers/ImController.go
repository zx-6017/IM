package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"

	//"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
)

type Admin_im_info struct {
	Id uint `gorm:"primary_key"`
	App_id string `json:"app_id"`
	Identifier string `json:"identifier"`
	Usersig string `json:"usersig"`
	Overtime string `json:"overtime"`
	Created_at timestamp.Timestamp `json:"created_at"`
	Updated_at timestamp.Timestamp `json:"updated_at"`
}

func Create(request *gin.Context){
	id, exist := request.GetQuery("id")
	if !exist {
		id = "id is not exist!"
	}
	info ,exist := request.GetQuery("info")

	if !exist{
		info = "info is not exist"
	}
	//初始化 数据库
	db,err := gorm.Open("mysql","root:mysqlpasswd@/yome?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	db.SingularTable(true)
	defer db.Close()

	insert_info := Admin_im_info{App_id:"app_id",Identifier:"identify"}
	fmt.Println(insert_info)
	db.Create(&insert_info)

	request.JSON(200,gin.H{
		"id":id,
		"info":info,
		"msg":insert_info.Id,
	})




}
