package models

type UserFriend struct {
	Id int `gorm:"primary_key;column:id,type:int"`
	Main_uid int `gorm:"column:main_uid,type:int"`
	Friend_uid int `gorm:"column:friend_uid,type:int"`

}