package internal

import (
	"StuScroeManageSys/internal/common"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func TotalInit() {
	viperinit()
	common.InitLogger()
}
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
