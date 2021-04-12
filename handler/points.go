package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/models"
)

// GET /points
func GetPoints(c *gin.Context) {
	coordinates, err := models.FindCoordinates()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": coordinates})
}
