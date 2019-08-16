package controllers

import "github.com/gin-gonic/gin"

func Create(request *gin.Context){
	id, exist := request.GetQuery("id")
	if !exist {
		id = "id is not exist!"
	}
	info ,exist := request.GetQuery("info")

	if !exist{
		info = "info is not exist"
	}

	request.JSON(200,gin.H{
		"id":id,
		"info":info,
	})




}
