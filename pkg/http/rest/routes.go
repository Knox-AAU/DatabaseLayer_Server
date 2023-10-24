package rest

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

func NewRouter(s graph.Service) *gin.Engine {
	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {
		getHandler(c, s)
	})
	return router
}
