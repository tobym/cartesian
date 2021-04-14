package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// getCoordinatesFromFile reads the values from the JSON file
func GetCoordinatesFromFile(filePath string) ([]Coordinate, error) {
	var coordinates []Coordinate

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
func CreateCoordinates(coordinates []Coordinate) {
	for _, coordinate := range coordinates {
		err := CreateCoordinate(coordinate)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// CalculateManhattanDistance between two coordiantes
func CalculateManhattanDistance(point1, point2 Coordinate) float64 {
	distanceXPoints := math.Abs(point1.X - point2.X)
	distanceYPoints := math.Abs(point1.Y - point2.Y)
	return distanceXPoints + distanceYPoints
}
