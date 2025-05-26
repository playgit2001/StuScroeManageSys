package handler

import (
	"StuScroeManageSys/internal/model"
	"StuScroeManageSys/internal/responsitory"
	"github.com/gin-gonic/gin"
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

}

//func Regists(c *gin.Context) {
//	var tmp model.RegisterUser
//	if err := c.ShouldBindJSON(&tmp); err != nil {
//		c.JSON(400, gin.H{"code:": 1, "msg": err.Error()})
//		return
//	}
//}
