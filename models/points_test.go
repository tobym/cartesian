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
