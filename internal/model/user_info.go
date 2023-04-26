package model

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	UserID    uint      `gorm:"column:user_id;type:bigint(20) unsigned;comment:用户id;NOT NULL"`
	City      string    `gorm:"column:city;type:varchar(255);comment:城市;NOT NULL"`
	Introduce string    `gorm:"column:introduce;type:varchar(255);comment:个人介绍;NOT NULL"`
	Fans      int       `gorm:"column:fans;type:int(8) unsigned;default:0;comment:粉丝数量"`
	Followee  int       `gorm:"column:followee;type:int(8) unsigned;default:0;comment:关注数量"`
	Gender    bool      `gorm:"column:gender;type:tinyint(1);default:0;comment:性别 0:未知 1:男 2:女"`
	BirthDay  time.Time `gorm:"column:birth_day;type:datetime;comment:生日"`
	Credits   int       `gorm:"column:credits;type:int(8) unsigned;default:0;comment:积分"`
	Level     int       `gorm:"column:level;type:tinyint(1);default:0;comment:等级"`
	User      *User     `gorm:"foreignKey:UserID;references:id"`
}
