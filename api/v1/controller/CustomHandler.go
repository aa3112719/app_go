package controller

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"app_go/common"
)

type CustomHandler struct{}

func (c *CustomHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/custom", c.MysqlDeadLockController)
}

type TStudent struct {
	Id    int    `json:"id"`
	No    string `json:"no"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (TStudent) TableName() string {
	return "t_student"
}

func (m *CustomHandler) MysqlDeadLockController(c *gin.Context) {

	id := "3"
	var user TStudent
	var wg sync.WaitGroup

	wg.Add(1)
	err := m.readOne(id, &user, &wg)

	wg.Wait()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"data":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    user,
		"params":  id,
	})
}

func (m *CustomHandler) readOne(id string, student *TStudent, wg *sync.WaitGroup) error {
	defer wg.Done()

	mysqlConnect, err := common.MysqlConnect()

	if err != nil {
		return err
	}

	mysqlConnect.First(&student, id)

	return nil
}
