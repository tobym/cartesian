package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/leozz37/cartesian/models"
)

var MySQL *gorm.DB

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

func CreateCoordinate(coordinate models.Coordinate) error {
	err := MySQL.Create(&coordinate).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCoordinates() error {
	err := MySQL.Delete(&models.Coordinate{}).Error
	if err != nil {
		return err
	}
	return nil
}

func FindCoordinates() ([]models.Coordinate, error) {
	var coordiantes []models.Coordinate

	err := MySQL.Find(&coordiantes).Error
	if err != nil {
		return nil, err
	}
	return coordiantes, nil
}
