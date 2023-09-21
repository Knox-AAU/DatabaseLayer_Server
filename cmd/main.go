package main

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/query", func(c *gin.Context) {
		rest.QueryHandler(c)
	})
	router.Run(":8080")
}
