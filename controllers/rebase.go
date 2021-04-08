package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/leozz37/cartesian/models"
)

func getCoordinatesFromFile(filePath string) []models.Coordinate {
	var coordinates []models.Coordinate

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Error reading JSON file.")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &coordinates)
	if err != nil {
		log.Println(err)
	}

	return coordinates
}
