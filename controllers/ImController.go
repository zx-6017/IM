package controllers

import (
	"IM/config"
	"IM/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImAction struct {
	State_StateChange string
	C2C_CallbackBeforeSendMsg string
}

func (this ImAction) CallbackBeforeSendMsg(){

}


func init(){

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



func ImCallBack(request *gin.Context){
	sdk_appid ,exist := request.GetQuery("SdkAppid")

	if !exist || sdk_appid != config.Conf.String("SDK_APP_ID"){
		request.JSON(200,gin.H{
			"ActionStatus":"FAIL",
			"ErrorCode":1,
			"ErrorInfo":"SDK APP ID ERROR",
		})
	}

	CallbackCommand,exist := request.GetQuery("CallbackCommand")
	fmt.Println(CallbackCommand)
	//if !exist || ImAction.CallbackCommand != {
	//	request.JSON(200,gin.H{
	//		"ActionStatus":"FAIL",
	//		"ErrorCode":1,
	//		"ErrorInfo":"IM ACTION ERROR",
	//	})
	//}

}

func reflect(b interface{},action string){

}
