package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leozz37/cartesian/controllers"
	"github.com/leozz37/cartesian/db"
	"github.com/leozz37/cartesian/models"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init Database and Migrate coordiantes model
	db.ConnectMySQL("mysql", os.Getenv("DATABASE_DSN"))
	db.AutoMigration(models.Coordinate{})

	// Get data from JSON file and save to DB
	controllers.Rebase()
}
