// Package main Database Layer Server API.
//
// This package implements a REST HTTP API to facilitate communication with the KNOX database.
//
// Terms Of Service:
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

//go:generate swagger generate spec -m -o ./swagger.yaml

import (
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
)

func main() {
	appRepository := config.Repository{}
	config.Load("..", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoURL, appRepository.VirtuosoUsername, appRepository.VirtuosoPassword)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service, graph.OntologyGraphURI(appRepository.OntologyGraphURI), graph.KnowledgeBaseGraphURI(appRepository.GraphURI))
	router.Run(":8000")
}
