package main

import (
	"app_go/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", controller.MysqlController)
	r.Run()
}
