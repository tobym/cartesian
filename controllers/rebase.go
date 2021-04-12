package controllers

import (
	"os"

	"github.com/leozz37/cartesian/models"
)

// Rebase delete as Coordinates from Database, read from the
// JSON file save them on Database
func Rebase() {
	models.DeleteCoordinates()

	filePath := os.Getenv("DATA_PATH")

	coordinates, err := models.GetCoordinatesFromFile(filePath)
	if err != nil {
		panic(err)
	}

	models.CreateCoordinates(coordinates)
}
