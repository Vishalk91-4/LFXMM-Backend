package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	router := gin.Default()

	router.GET("/search", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "The Search is running perfectly.",
		})
	})

	// Register all the routes and there corresponding handlers
	router.GET("/search/projects", getOrg)

	router.Run("0.0.0.0:8080")
}
