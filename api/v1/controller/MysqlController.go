package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (User) TableName() string {
	return "user"
}

func MysqlController(c *gin.Context) {
	mysqlDsn := "root:root@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	id := c.Query("id")
	var user User

	// 事务
	tx := db.Begin()

	tx.First(&user, id)
	tx.Update("name", "test2").Where("id = ?", id).Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    user,
		"params":  id,
	})
}
