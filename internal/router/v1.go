package router

import (
	"StuScroeManageSys/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		//登录
		v1.POST("/login", handler.GetUser)
		v1.POST("/addUser", handler.AddUser)
		v1.GET("/getUsers/:id", handler.GetUserByID)
		v1.DELETE("/deleteUser/:id", handler.DeleteUser)
		v1.POST("/updateUser", handler.UpdateUser)
		v1.GET("/UpdateUserID/:newid/:oldid", handler.UpdateUserByID)
		v1.DELETE("/deleteUserByIds", handler.DeleteUserByIdsHandler)
	}
}
