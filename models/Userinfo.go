package models

import (
	"IM/helpers"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User_info struct {
	Id uint `gorm:"primary_key;column:id,type:int"`
	Beautiful int `gorm:"column:beautiful,type:int,not null"`
	Fb_id string `gorm:"column:fb_id,type:varchar(255),not null`
	F string `gorm:"column:f,type:varchar(24),not null`
	Dev string `gorm:"column:dev,type:varchar(24),not null`
	Name string `gorm:"column:name,type:varchar(255),not null`
	Age string `gorm:"column:age,type:varchar(5),not null`
	Year string `gorm:"column:year,type:varchar(255),not null`
	Sex string `gorm:"column:fb_id,type:varchar(5),not null`
	Created_at timestamp.Timestamp `gorm:"column:created_at,type:timestamp,not null`
	Updated_at timestamp.Timestamp `gorm:"column:updated_at,type:timestamp,not null`
	Deleted_at timestamp.Timestamp `gorm:"column:deleted_at,type:timestamp,not null`
}

func (user User_info) GetOneUser(id int)User_info{
	var user_info User_info
	helpers.DB.Where("id = ?",id).Find(user_info)
	return user_info;
}


//func GetUserInfos()[]User_info{
//
//}









