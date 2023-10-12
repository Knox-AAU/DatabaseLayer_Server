// Package main Database Layer Server API.
//
// This repository implements a server to facilitate communication on the KNOX pipeline.
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
	"log"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
)

func main() {
	log.Printf("I am an AAU project that is deployed using Watchtower.\nhello world\nhello world\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nhello world :)\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nhello world!\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nhello there... o.o\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	appRepository := config.Repository{}
	config.LoadEnv("..", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoServerURL)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service)
	router.Run(":8080")
}
