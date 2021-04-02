package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
