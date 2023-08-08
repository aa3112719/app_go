package controller

import (
	"fmt"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup

	wg.Add(2)
	go updateOne(id, &user, &wg)
	go readOne(id, &wg)

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    user,
		"params":  id,
	})
}

func updateOne(id string, user *User, wg *sync.WaitGroup) {
	defer wg.Done()
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
	d := 10 * time.Second
	time.Sleep(d)
	tx.Rollback()
}

func readOne(id string, wg *sync.WaitGroup) {
	defer wg.Done()

	var user User
	mysqlDsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 事务级别是读未提交
	tx := db.Begin()
	tx.Exec("SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED")

	count := 100
	for i := 0; i < count; i++ {

		tx.First(&user, id)
		if user.Name == "先请求" {
			fmt.Println("读到了")
		}
	}

	tx.Commit()
}
