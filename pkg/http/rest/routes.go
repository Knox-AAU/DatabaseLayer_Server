package rest

import (
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

type (
	Route  string
	Method string
)

const (
	Triples Route  = "/triples"
	GET     Method = http.MethodGet
	POST    Method = http.MethodPost
)

func NewRouter(s graph.Service, ontologyGraph graph.OntologyGraphURI, knowledgeBaseGraph graph.KnowledgeBaseGraphURI, apiSecret string) *gin.Engine {
	router := gin.Default()
	router.Use(authenticate(apiSecret))
	router.Use(validateGraphParameter([]graph.TargetGraph{graph.TargetGraph(ontologyGraph), graph.TargetGraph(knowledgeBaseGraph)}))
	router.GET(string(Triples), getHandler(s))
	router.POST(string(Triples), postHandler(s))

	return router
}
