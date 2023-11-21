package rest

import (
	"log"
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

type (
	Route  string
	Method string
)

const (
	KnowledgeBase Route = "/knowledge-base"
	Ontology      Route = "/ontology"
)
const (
	GET  Method = http.MethodGet
	POST Method = http.MethodPost
)

func NewRouter(s graph.Service, ontologyGraph graph.OntologyGraphURI, knowledgeBaseGraph graph.KnowledgeBaseGraphURI) *gin.Engine {
	router := gin.Default()
	selectGraph := func(route Route) graph.TargetGraph {
		switch route {
		case KnowledgeBase:
			return graph.TargetGraph(knowledgeBaseGraph)
		case Ontology:
			return graph.TargetGraph(ontologyGraph)
		default:
			log.Fatalf("unknown route: %s", route)
			return ""
		}
	}

	setRoute := func(route Route, method Method) {
		switch method {
		case GET:
			router.GET(string(route), func(c *gin.Context) {
				getHandler(c, s, selectGraph(route))
			})
		case POST:
			router.POST(string(route), func(c *gin.Context) {
				postHandler(c, s, selectGraph(route))
			})
		default:
			log.Fatalf("unknown method: %s", method)
		}
	}

	setRoute(KnowledgeBase, GET)
	setRoute(Ontology, GET)
	setRoute(KnowledgeBase, POST)
	setRoute(Ontology, POST)

	return router
}
