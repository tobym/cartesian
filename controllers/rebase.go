package controllers

import (
	"os"

	"github.com/leozz37/cartesian/models"
)

// Rebase delete as Coordinates from Database, read from the
// JSON file save them on Database
func Rebase() {
	models.DeleteCoordinates()

	filepath := os.Getenv("DATA_PATH")

	coordinates, err := models.GetCoordinatesFromFile(filepath)
	if err != nil {
		panic(err)
	}

	models.CreateCoordinates(coordinates)
}
