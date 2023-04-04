package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	gorm.Model
	UserID    uint
	City      string
	Introduce string
	Fans      int
	Followee  int
	Gender    bool
	BirthDay  time.Time
	Credits   int
	Level     bool
}
