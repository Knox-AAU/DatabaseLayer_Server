package main

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
	"github.com/gin-gonic/gin"
)

func main() {
	appRepository := config.Repository{}
	config.LoadEnv("..", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoServerURL)
	router := gin.Default()
	service := graph.NewService(virtuosoRepository)
	router.GET("/query", func(c *gin.Context) {
		rest.GETallHandler(c, service)
	})
	router.Run(":8080")
}
