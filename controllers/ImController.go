package controllers

import (
	"IM/config"
	"IM/controllers/Struct"
	"IM/helpers"
	"IM/models"
	"github.com/gomodule/redigo/redis"
	"strings"

	_"IM/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Create(request *gin.Context){
	id, exist := request.GetQuery("id")
	if !exist {
		id = "id is not exist!"
		return
	}
	info ,exist := request.GetQuery("info")

	if !exist{
		info = "info is not exist"
		return
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

	//CallbackCommand,exist := request.GetQuery("CallbackCommand")

	data,data_err := request.GetRawData()
	if data_err != nil{
		request.JSON(200,gin.H{
			"ActionStatus":"FAIL",
			"ErrorCode":1,
			"ErrorInfo":"Get Data Body error" + data_err.Error(),
		})
	}

	var data_info  Struct.Data_struct

	json_err := json.Unmarshal(data,&data_info)
	if json_err != nil{
		fmt.Println("json post data err",json_err.Error())
		return
	}
	//sender := helpers.GetUidByIdentify(data_info.From_Account)
	//receiver := helpers.GetUidByIdentify(data_info.To_Account)

	dirty_words := []models.ImDirtyWord{}
	cache_imDirtyWord,err := helpers.RedisPool.Get().Do("get","im_dirty_word")

	if err != nil{
		fmt.Println("reids cache dirty_words get err " , err.Error())
	}

	if cache_imDirtyWord == nil { //没有缓存值
		model_imdirtyword := models.ImDirtyWord{}
		dirty_words = model_imdirtyword.All_Infos()
		cache_imDirtyWord, _ = json.Marshal(dirty_words)
		_,save_redis_err :=helpers.RedisPool.Get().Do("set","im_dirty_word",cache_imDirtyWord)
		if save_redis_err  != nil{
			fmt.Println("save redis cache im_dirty_word err",save_redis_err.Error())
		}
	}
	cache_imDirtyWord_byte ,_:= redis.Bytes(cache_imDirtyWord,nil)
	json_err  = json.Unmarshal(cache_imDirtyWord_byte,&dirty_words)
	if json_err != nil{
		fmt.Println("json decode err",json_err.Error())
	}


	msg_data := Struct.Msg_data{}

	for _,msgbody := range data_info.MsgBody {
		 if msgbody.MsgType != "TIMCustomElem"{
		 	continue
		 }
		 err =json.Unmarshal([]byte(msgbody.MsgContent.Data),&msg_data)
		 if err != nil{
		 	fmt.Println("json decode err",err.Error())
		 	continue
		 }
		 text := msg_data.PrimaryText
		 for _,word_info := range dirty_words{
			 if strings.IndexAny(text,word_info.Word) != -1 {
			 	request.JSON(200,gin.H{
					 "ActionStatus":"FAIL",
					 "ErrorCode":1,
					 "ErrorInfo":"dirty word " + word_info.Word ,
				 })
			 	return
			 }
		 }

	}



}
