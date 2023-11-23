package rest

import (
	"fmt"
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

func validateGraphParameter(validGraphs []graph.TargetGraph) gin.HandlerFunc {
	return func(c *gin.Context) {
		graphParameter := c.Query(string(graph.Graph))
		for _, graph := range validGraphs {
			if string(graph) == graphParameter {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("got graph parameter: %s, want one of: %v", graphParameter, validGraphs)})
		c.Abort()
	}
}
