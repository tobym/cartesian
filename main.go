package main

import (
	"os"

	"github.com/leozz37/cartesian/db"
	"github.com/leozz37/cartesian/models"
)

func main() {
	// Init Database and Migrate coordiantes model
	db.ConnectMySQL("mysql", os.Getenv("DATABASE_URL"))
	db.AutoMigration(models.Coordinate{})
}
