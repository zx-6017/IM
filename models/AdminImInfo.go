package models

import (
	"IM/helpers"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type AdminImInfo struct {
	Id uint `gorm:"primary_key",json:"id"`
	App_id string `json:"app_id"`
	Identifier string `json:"identifier"`
	Usersig string `json:"usersig"`
	Overtime string `json:"overtime"`
	Created_at timestamp.Timestamp `json:"created_at"`
	Updated_at timestamp.Timestamp `json:"updated_at"`
}
func (this *AdminImInfo)CreateAdminImInfo(info *AdminImInfo){
	helpers.DB.Create(info)
}

