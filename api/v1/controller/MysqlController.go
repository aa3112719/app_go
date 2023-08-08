package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (User) TableName() string {
	return "user"
}

func MysqlController(c *gin.Context) {

	id := c.Query("id")
	var user User

	go updateOne(id, &user)
	go readOne(id, &user)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    user,
		"params":  id,
	})
}

func updateOne(id string, user *User) {
	mysqlDsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()
	tx.Exec("SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED")
	tx.First(&user, id)
	user.Name = "先请求"
	tx.Save(&user)
	time.Sleep(10 * time.Second)
	tx.Rollback()
}

func readOne(id string, user *User) {
	mysqlDsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// 事务级别是读未提交
	tx := db.Begin()
	tx.Exec("SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED")

	time.Sleep(5 * time.Second)
	tx.First(&user, id)

	fmt.Println(user)

	time.Sleep(5 * time.Second)
	tx.Commit()
}
