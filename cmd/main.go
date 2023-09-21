package main

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/repository/virtuoso"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/retrieval"
	"github.com/gin-gonic/gin"
)

func main() {
	_config := config.AppConfig{}
	config.LoadEnv("..", &_config)
	repo := virtuoso.VirtuosoRepository{}
	router := gin.Default()
	r := retrieval.NewService(repo)
	router.GET("/query", func(c *gin.Context) {
		rest.QueryHandler(c, r)
	})
	router.Run(":8080")
}
