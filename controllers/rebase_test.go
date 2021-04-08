package controllers

import (
	"os"
	"testing"

	"github.com/leozz37/cartesian/db"
	"github.com/leozz37/cartesian/models"
)

func setupMySQL() {
	db.ConnectMySQL("sqlite3", "testing.db")
	db.DeleteCoordinates()
	db.AutoMigration(models.Coordinate{})
}

func TestGetCoordinatesFromFileSuccess(t *testing.T) {
	// The very first coordinate from points.json
	expectedPoint := models.Coordinate{X: 63, Y: -72}

	coordinates, err := getCoordinatesFromFile("../data/points.json")
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

	var coordinates []models.Coordinate
	coordinate := models.Coordinate{X: 1, Y: 2}
	coordinates = append(coordinates, coordinate)
	createCoordinates(coordinates)

	result, err := db.FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	if result[0].X != coordinate.X {
		t.Error("X coordiante is differente")
	} else if result[0].Y != coordinate.Y {
		t.Error("Y coordiante is differente")
	}
}

func TestRebaseSuccess(t *testing.T) {
	setupMySQL()
	os.Setenv("DATA_PATH", "../data/points.json")

	Rebase()
	result, err := db.FindCoordinates()
	if err != nil {
		t.Error(err)
	}

	expectedCoordiante := models.Coordinate{X: 63, Y: -72}

	if result[0].X != expectedCoordiante.X {
		t.Error("X coordiante is differente")
	} else if result[0].Y != expectedCoordiante.Y {
		t.Error("Y coordiante is differente")
	}
}

func TestRebaseFailFilePath(t *testing.T) {
	setupMySQL()
	os.Setenv("DATA_PATH", "random/path/unknow.json")

	// Checking if Rebase() will panic!
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	Rebase()
}
