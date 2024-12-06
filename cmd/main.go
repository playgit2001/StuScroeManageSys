package main

import (
	"StuScroeManageSys/internal/common"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func viperinit() {
	dir, _ := os.Getwd()
	dir = strings.TrimSuffix(dir, "/cmd")
	viper.SetConfigName("config")
	viper.AddConfigPath(dir + "/config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {

		panic("")
	}
}

func main() {
	viperinit()
	DB := common.InitDB()
	DB.Close()

}
