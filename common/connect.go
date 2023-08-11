package common

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() (*gorm.DB, error) {
	mysqlDsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})

	return db, err
}
