package models

import (
	"IM/helpers"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type ImDirtyWord struct {
	Id int `gorm:"primary_key",json:"id"`
	Word string `json:"word"`
	Language string `json:"language"`
	Type string `json:"type"`
	Score int `json:"score"`
	Created_at timestamp.Timestamp `json:"created_at"`
	Updated_at timestamp.Timestamp `json:"updated_at"`

}

func (this *ImDirtyWord) All_Infos() ([]ImDirtyWord){
	res := []ImDirtyWord{}
	helpers.DB.Table("im_dirty_word").Find(&res)
	return res
}