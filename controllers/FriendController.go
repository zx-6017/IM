package controllers

import (
	"IM/models"
	"github.com/gin-gonic/gin"
)

type Res_age struct {
	Age16 int
	Age16_19 int
	Age20_24 int
	Age25_29 int
	Age30_34 int
	Age35_40 int
	Age40 int
}
type Timezone_data struct {
	Timezone string
	Un_age_num Res_age
	Age_num Res_age
}

func FriendRelation(request *gin.Context) {
	UserInfo := models.UserInfo{}

	var res []Timezone_data

	timezone := [...]string{"Asia/Ho_Chi_Minh","Asia/Saigon"}

	for _,value := range timezone{
		var un_age_num Res_age
		var age_num Res_age

		un_ages := UserInfo.GetUnFriendRelationByTimezone(value)
		for _,age :=range un_ages {
			if age.Age < 16 {
				un_age_num.Age16 ++
			}else if(age.Age <= 19){
				un_age_num.Age16_19++
			}else if(age.Age <=24){
				un_age_num.Age20_24++
			}else if(age.Age <= 29){
				un_age_num.Age25_29++
			}else if(age.Age<=34){
				un_age_num.Age30_34++
			}else if(age.Age<=40){
				un_age_num.Age35_40++
			}else if(age.Age>40){
				un_age_num.Age40++
			}
		}
		ages := UserInfo.GetFriendRelationByTimezone(value)
		for _,age :=range ages {
			if age.Age < 16 {
				age_num.Age16 ++
			}else if(age.Age <= 19){
				age_num.Age16_19++
			}else if(age.Age <=24){
				age_num.Age20_24++
			}else if(age.Age <= 29){
				age_num.Age25_29++
			}else if(age.Age<=34){
				age_num.Age30_34++
			}else if(age.Age<=40){
				age_num.Age35_40++
			}else if(age.Age>40){
				age_num.Age40++
			}
		}
		var timedata Timezone_data
		timedata.Timezone = value
		timedata.Un_age_num = un_age_num
		timedata.Age_num = age_num
		res = append(res, timedata)

	}
	request.JSON(200,res)

}