package models

import (
	"testing"

	"github.com/leozz37/cartesian/db"
)

func setupMySQL() {
	db.ConnectMySQL("sqlite3", "testing.db")
	DeleteCoordinates()
	db.AutoMigration(Coordinate{})
}

func TestGetCoordinatesFromFileSuccess(t *testing.T) {
	// The very first coordinate from points.json
	expectedPoint := Coordinate{X: 63, Y: -72}

	coordinates, err := GetCoordinatesFromFile("../data/points.json")
	if err != nil {
		t.Error(err)
	}

	for _, v := range coordinates {
		if v.X == expectedPoint.X && v.Y == expectedPoint.Y {
			return
		}
	}

	t.Error("Couldn't find coordinates")
}

func TestCreateCoordiantesSuccess(t *testing.T) {
	setupMySQL()

	var coordinates []Coordinate
	coordinate := Coordinate{X: 1, Y: 2}
	coordinates = append(coordinates, coordinate)
	CreateCoordinates(coordinates)

	result, err := FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	if result[0].X != coordinate.X {
		t.Error("X coordiante is differente")
	} else if result[0].Y != coordinate.Y {
		t.Error("Y coordiante is differente")
	}
}

func TestCalculateManhattanDistance(t *testing.T) {
	// First test suite
	point1 := Coordinate{X: 63, Y: -72}
	point2 := Coordinate{X: -94, Y: 89}

	expectedResult := 318.0
	result := CalculateManhattanDistance(point1, point2)
	if expectedResult != result {
		t.Errorf("Result expected: %f, got: %f", expectedResult, result)
	}

	// Second test suite
	point1 = Coordinate{X: 30, Y: 50}
	point2 = Coordinate{X: -90, Y: 60}

	expectedResult = 130.0
	result = CalculateManhattanDistance(point1, point2)
	if expectedResult != result {
		t.Errorf("Result expected: %f, got: %f", expectedResult, result)
	}
}
