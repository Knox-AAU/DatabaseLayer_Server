package rest

import (
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query"`
}

func GETallHandler(c *gin.Context, r graph.Service) {
	node, err := r.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}
