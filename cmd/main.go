// Package main Database Layer Server API.
//
// REST API for the KNOX database.
//
// Terms Of Service:
//
//	Schemes: http
//	Host: http://192.38.54.90
//	BasePath: /
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

//go:generate swagger generate spec -m -o ../swagger.yaml
//go:generate openapi-markdown -i ../swagger.yaml

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
)

func main() {
	appRepository := config.Config{}
	config.Load("..", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoURL, appRepository.VirtuosoUsername, appRepository.VirtuosoPassword)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service, graph.OntologyGraphURI(appRepository.OntologyGraphURI), graph.KnowledgeBaseGraphURI(appRepository.GraphURI), appRepository.APISecret)
	router.Run(":8000")
}
