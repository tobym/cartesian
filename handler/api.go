package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func FindDistances(c *gin.Context) {
	distance := c.Query("distance")
	x := c.Query("x")
	y := c.Query("y")

	err := validateQuery(x, y, distance)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s | %s | %s", distance, x, y)})
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
