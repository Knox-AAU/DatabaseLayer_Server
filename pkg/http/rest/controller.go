package rest

import (
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query"`
}

func MakeHandler(service graph.Service) *gin.Engine {
	router := gin.Default()
	router.GET("/get-all", func(c *gin.Context) {
		GETallHandler(c, service)
	})
	return router
}

func GETallHandler(c *gin.Context, s graph.Service) {
	virtuosoObject, err := s.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	node, err := virtuosoObject.Read()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource does not exist"})
		return
	}

	c.JSON(http.StatusOK, virtuosoObject)
}
