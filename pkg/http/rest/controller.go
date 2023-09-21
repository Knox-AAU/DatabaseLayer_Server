package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query"`
}

func QueryHandler(c *gin.Context) {
	body := Body{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"query": body.Query,
	})
}
