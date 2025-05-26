package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Loginquest struct {
	gorm.Model
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type Admin struct {
	ID        int    `form:"ID" json:"id" gorm:"primary_key"`
	UserName  string `form:"UserName" json:"UserName" binding:"required" gorm:"column:UserName;not null"`
	Password  string `form:"password" json:"password" binding:"required" gorm:"not null"`
	Level     int    `form:"level" json:"level" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type User struct {
	gorm.Model
	ID         int    `form:"ID" json:"id" binding:"required" gorm:"primary_key"`
	UserName   string `form:"UserName" json:"Username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Age        int    `form:"age" json:"age" binding:"gte=0"`
	Sex        int    `form:"sex" json:"sex" binding:"gte=0"`
	Department string `form:"department" json:"department" binding:"required"`
	Class      string `form:"class" json:"class" binding:"required"`
}
