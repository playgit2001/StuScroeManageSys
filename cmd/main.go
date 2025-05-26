package main

import (
	"StuScroeManageSys/internal"
	"StuScroeManageSys/internal/common"
	"StuScroeManageSys/internal/model"
	"StuScroeManageSys/internal/router"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	internal.TotalInit()
	common.InitDB()
	db := common.GetDB()
	if err := model.AtuoMigrateAll(db); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
