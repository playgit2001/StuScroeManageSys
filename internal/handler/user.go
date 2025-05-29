package handler

import (
	"StuScroeManageSys/internal/model"
	"StuScroeManageSys/internal/responsitory"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func GetUser(c *gin.Context) {
	var tmp model.Loginquest
	if err := c.ShouldBindJSON(&tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	err := responsitory.GetUser(tmp)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "not find User"})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}
func AddUser(c *gin.Context) {
	var tmp model.User
	if err := c.ShouldBindJSON(&tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "参数格式错误"})
		return
	}
	err := responsitory.AddUser(tmp)
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}
func GetUserByID(c *gin.Context) {
	var tmp string
	var id int
	var user model.User
	var err error
	tmp = c.Param("id")
	if id, err = strconv.Atoi(tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "ID非数字格式"})
		return
	}
	if user, err = responsitory.GetUserByID(id); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success", "data": user})
}
func DeleteUser(c *gin.Context) {
	var tmp string
	var id int
	var err error
	tmp = c.Param("id")
	if id, err = strconv.Atoi(tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "ID非数字格式"})
		return
	}
	if err = responsitory.DeteleUser(id); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}
func UpdateUser(c *gin.Context) {
	var tmp model.User
	if err := c.ShouldBindJSON(&tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "参数格式错误"})
		return
	}
	if err := responsitory.UpdateUser(tmp); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
	return

}
func UpdateUserByID(c *gin.Context) {
	var tmp1, tmp2 string
	tmp1 = c.Param("newid")
	tmp2 = c.Param("oldid")
	newid, err1 := strconv.Atoi(tmp1)
	oldid, err2 := strconv.Atoi(tmp2)
	fmt.Println(tmp1, tmp2)
	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "ID非数字格式"})
		return
	}
	err := responsitory.UpdateUserId(newid, oldid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{"code": 1, "msg": "该学号不存在"})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}

func DeleteUserByIdsHandler(c *gin.Context) {
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "data error"})
		return
	}
	if err := responsitory.DeleteUserByIds(ids); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})

}

//func Regists(c *gin.Context) {
//	var tmp model.RegisterUser
//	if err := c.ShouldBindJSON(&tmp); err != nil {
//		c.JSON(400, gin.H{"code:": 1, "msg": err.Error()})
//		return
//	}
//}
