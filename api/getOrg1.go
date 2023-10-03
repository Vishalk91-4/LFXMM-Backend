package api

import (
	"eshaanagg/lfx/database/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOrg1(c *gin.Context) {
	client := handlers.New()
	defer client.Close()

	id := c.Param("id")
	org := client.GetOrganizationByID(id)

	if org == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "There are 2 organization with this id",
		},
		)
		return
	}

	projects := client.GetProjectsByParentOrgID(id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":           org.ID,
		"logo":         org.Logo,
		"org":          org.Name,
		"projectCount": len(projects),
		"search":       projects,
	})
}
