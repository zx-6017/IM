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

	admin_im_info := models.AdminImInfo{}

	//向admin_im_info插入一条信息
	insert_adminImInfo_info := models.AdminImInfo{
		App_id:"app_id",
		Identifier:"identify",
	}
	admin_im_info.CreateAdminImInfo(&insert_adminImInfo_info)

	//从user_info 获取信息
	userInfoModel := models.UserInfo{}
	user_info := userInfoModel.GetOneUser(6778)

	request.JSON(200,gin.H{
		"id":id,
		"info":info,
		"admin_info_id":insert_adminImInfo_info.Id,
		"user_info_6778_fb_id":user_info.Fb_id,
	})




}
