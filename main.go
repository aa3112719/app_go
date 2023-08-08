package main

import (
	"app_go/api/v1/controller"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config/config.yaml") // 设置配置文件路径
	err := viper.ReadInConfig()                 // 读取配置文件
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	port := viper.GetString("server.port") // 获取端口配置

	r := gin.Default()

	r.GET("/ping", controller.MysqlController)
	r.Run(":" + port)
}
