package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leozz37/cartesian/handler"
)

func InitRoutes() {
	r := gin.Default()

	r.GET("/api/points", handler.FindDistance)
	r.GET("points", handler.GetPoints)
	r.NoRoute(noRoute)

	r.Run()
}

func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
}
