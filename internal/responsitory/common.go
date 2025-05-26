package responsitory

import (
	"StuScroeManageSys/internal/common"
	"StuScroeManageSys/internal/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func GetUser(login model.Loginquest) error {
	db := common.GetDB()
	var user model.Admin
	err := db.Table("admin").Where("UserName = ? and password = ?", login.UserName, login.Password).First(&user).Error
	return err
}

func AddUser(user model.User) error {
	db := common.GetDB()
	user.CreatedAt = time.Time{}
	err := db.Create(&user).Error
	return err
}

func GetUserByID(id int) (model.User, error) {
	db := common.GetDB()
	var user model.User
	err := db.Find(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, fmt.Errorf("not find User")
	}
	if err != nil {
		return user, fmt.Errorf("database error:%v", err)
	}
	return user, nil
}
func DeteleUser(id int) error {
	db := common.GetDB()
	var user model.User
	if err := db.Find(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("not find User")
		} else {
			return fmt.Errorf("数据库异常")
		}
	}
	print(id)
	if err := db.Delete(model.User{}, id).Error; err != nil {
		return fmt.Errorf("database error:%v", err)
	}
	return nil
}
