package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	MySQL *gorm.DB
)

func ConnectMySQL(databseType, dsn string) {
	database, err := gorm.Open(databseType, dsn)
	if err != nil {
		panic(err)
	}
	MySQL = database
}

func AutoMigration(model interface{}) {
	MySQL.AutoMigrate(model)
}
