package main

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
)

func main() {
	appRepository := config.Repository{}
	config.LoadEnv("..", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoServerURL)
	service := graph.NewService(virtuosoRepository)
	router := rest.MakeHandler(service)
	router.Run(":8080")
}
