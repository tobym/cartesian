package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/leozz37/cartesian/db"
	"github.com/leozz37/cartesian/models"
)

// getCoordinatesFromFile reads the values from the JSON file
func getCoordinatesFromFile(filePath string) ([]models.Coordinate, error) {
	var coordinates []models.Coordinate

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}

// createCoordiantes saves the coordinates to Database
func createCoordinates(coordinates []models.Coordinate) {
	for _, coordinate := range coordinates {
		err := db.CreateCoordinate(coordinate)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Rebase delete as Coordinates from Database, read from the
// JSON file save them on Database
func Rebase() {
	db.DeleteCoordinates()

	filepath := os.Getenv("DATA_PATH")
	coordinates, err := getCoordinatesFromFile(filepath)
	if err != nil {
		panic(err)
	}

	createCoordinates(coordinates)
}
