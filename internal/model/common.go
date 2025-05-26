package model

import "github.com/jinzhu/gorm"

func AtuoMigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&Admin{}, &User{}).Error
}
