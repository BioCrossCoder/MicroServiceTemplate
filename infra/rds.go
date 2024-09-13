package infra

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	rds     *gorm.DB
	rdsOnce sync.Once
)

func NewMySQLClient() *gorm.DB {
	rdsOnce.Do(func() {
		dsn := "root:root@tcp(localhost:3306)/app?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		rds, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return rds
}
