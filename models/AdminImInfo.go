package models

import (
	"IM/helpers"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type AdminImInfo struct {
	Id uint `gorm:"primary_key"`
	App_id string `json:"app_id"`
	Identifier string `json:"identifier"`
	Usersig string `json:"usersig"`
	Overtime string `json:"overtime"`
	Created_at timestamp.Timestamp `json:"created_at"`
	Updated_at timestamp.Timestamp `json:"updŒated_at"`
}
func (AdminImInfo *AdminImInfo)CreateAdminImInfo(info *AdminImInfo){
	helpers.DB.Create(info)
}

