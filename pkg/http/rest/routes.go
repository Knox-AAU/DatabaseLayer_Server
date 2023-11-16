package rest

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

const (
	GET  = "/get"
	POST = "/post"
)

func NewRouter(s graph.Service) *gin.Engine {
	router := gin.Default()
	router.GET(GET, func(c *gin.Context) {
		getHandler(c, s)
	})

	router.POST(POST, func(c *gin.Context) {
		postHandler(c, s)
	})

	return router
}
