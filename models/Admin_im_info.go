package models

import (
	"IM/helpers"
	"github.com/golang/protobuf/ptypes/timestamp"
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
func (Admin_im_info *Admin_im_info)CreateAdminImInfo(info *Admin_im_info){
	helpers.DB.Create(info)
}

