package controllers

import (
	"IM/models"
	"github.com/gin-gonic/gin"
)



func Create(request *gin.Context){
	id, exist := request.GetQuery("id")
	if !exist {
		id = "id is not exist!"
	}
	info ,exist := request.GetQuery("info")

	if !exist{
		info = "info is not exist"
	}

	admin_im_info := models.Admin_im_info{}

	//向admin_im_info插入一条信息
	insert_adminImInfo_info := models.Admin_im_info{
		App_id:"app_id",
		Identifier:"identify",
	}
	admin_im_info.CreateAdminImInfo(&insert_adminImInfo_info)

	//从user_info 获取信息


	request.JSON(200,gin.H{
		"id":id,
		"info":info,
		"admin_info_id":insert_adminImInfo_info.Id,
	})




}
