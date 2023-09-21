package rest

import (
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/retrieval"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query"`
}

func QueryHandler(c *gin.Context, r retrieval.Service) {
	body := Body{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := r.Query(body.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}