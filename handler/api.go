package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/models"
)

func FindDistances(c *gin.Context) {
	distance := c.Query("distance")
	x := c.Query("x")
	y := c.Query("y")

	err := validateQuery(x, y, distance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Converting string values to float
	distanceFloat, _ := strconv.ParseFloat(distance, 64)
	xFloat, _ := strconv.ParseFloat(x, 64)
	yFloat, _ := strconv.ParseFloat(y, 64)

	// Getting machinting distances
	coordiante := models.Coordinate{X: xFloat, Y: yFloat}
	matches, err := getWithinDistances(coordiante, distanceFloat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": matches})
}

func getWithinDistances(coordinate models.Coordinate, distance float64) ([]models.Coordinate, error) {
	var matchingDistances []models.Coordinate

	// Getting coordinates from file
	coordinates, err := models.FindCoordinates()
	if err != nil {
		return nil, err
	}

	// Comparing the manhattan distance with the coordinates
	// from the file
	for _, point := range coordinates {
		manhattanDistance := models.CalculateManhattanDistance(coordinate, point)
		if manhattanDistance == distance {
			matchingDistances = append(matchingDistances, point)
		}
	}

	return matchingDistances, nil
}

func validateQueryValue(value string) error {
	match, _ := regexp.MatchString("[+-]?([0-9]*[.])?[0-9]+", value)
	if !match {
		return errors.New(value + " is an invalid value")
	}
	return nil
}

func validateQuery(x, y, distance string) error {
	if err := validateQueryValue(distance); err != nil {
		return err
	}
	if err := validateQueryValue(x); err != nil {
		return err
	}
	if err := validateQueryValue(y); err != nil {
		return err
	}
	return nil
}
