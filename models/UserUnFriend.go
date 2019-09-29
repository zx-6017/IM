package models

type UserUnFriend struct {
	Id int `gorm:"primary_key;column:id,type:int"`
	Receive_uid int `gorm:"column:receive_uid,type:int"`
	Send_uid int `gorm:"column:send_uid,type:int"`
}