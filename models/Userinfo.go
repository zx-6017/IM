package models

import (
	"IM/helpers"
	"time"
)

type UserInfo struct {
	Id uint `gorm:"primary_key;column:id"`
	Beautiful int `gorm:"column:beautiful,type:int,not null"`
	Fb_id string `gorm:"column:fb_id,type:varchar(255),not null`
	F string `gorm:"column:f,type:varchar(24),not null`
	Dev string `gorm:"column:dev,type:varchar(24),not null`
	Name string `gorm:"column:name,type:varchar(255),not null`
	Age string `gorm:"column:age,type:varchar(5),not null`
	Year string `gorm:"column:year,type:varchar(255),not null`
	Sex string `gorm:"column:sex,type:varchar(5),not null`
	Timezone string `gorm:"column:timezone,type:varchar(255),not null`
	Created_at time.Time `gorm:"column:created_at,type:timestamp,not null`
	Updated_at time.Time `gorm:"column:updated_at,type:timestamp,not null`
	//Deleted_at time.Time `gorm:"column:deleted_at,type:timestamp`
}
type RelationAge struct {
		Age int
}

func (user UserInfo) GetOneUser(id int)UserInfo{
	var user_info UserInfo
	helpers.DB.Where("id = ?",id).Find(&user_info)
	return user_info;
}

func (UserInfo) GetUnFriendRelationByTimezone(timezone string) []RelationAge{
	var ages []RelationAge
	helpers.DB.Table("user_un_friend as un_f").
		Joins("join user_info as u_a on un_f.send_uid=u_a.id").
		Joins("join user_info as u_b on un_f.receive_uid=u_b.id").
		Where("u_a.timezone='"+timezone+"'").
		Where("u_a.created_at >= '2019-08-01'").
		Where("u_a.sex = 2").
		Where("u_a.is_normal = 1").
		Select("u_b.age").
		Scan(&ages)
	return ages
}

func (UserInfo) GetFriendRelationByTimezone(timezone string)[]RelationAge{
	var main_ages []RelationAge
	var friend_ages []RelationAge
	helpers.DB.Table("user_friend as u_f").
		Joins("join user_info as u_a on u_f.main_uid=u_a.id").
		Joins("join user_info as u_b on u_f.friend_uid=u_b.id").
		Where("u_a.timezone='"+timezone+"'").
		Where("u_a.created_at >= '2019-08-01'").
		Where("u_a.sex = 2").
		Where("u_a.is_normal = 1").
		Select("u_b.age").
		Scan(&main_ages)
	//fmt.Println(len(main_ages))
	helpers.DB.Table("user_friend as u_f").
		Joins("join user_info as u_a on u_f.main_uid=u_a.id").
		Joins("join user_info as u_b on u_f.friend_uid=u_b.id").
		Where("u_b.timezone='"+timezone+"'").
		Where("u_b.created_at >= '2019-08-01'").
		Where("u_b.sex = 2").
		Where("u_b.is_normal = 1").
		Select("u_a.age").
		Scan(&friend_ages)
	//fmt.Println(len(friend_ages))

	for _,value := range friend_ages{
			main_ages = append(main_ages,value)
		}
	//fmt.Println(len(main_ages))
	return main_ages

}









