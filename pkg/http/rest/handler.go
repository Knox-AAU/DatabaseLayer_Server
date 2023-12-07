package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

// swagger:model
type Result struct {
	Triples []graph.GetTriple `json:"triples"`
	Query   string            `json:"query"`
}

func getHandler(s graph.Service) func(c *gin.Context) {
	// swagger:operation GET /triples getTriples
	//
	// Query graph applying filters.
	//
	// To query the whole graph, leave parameters empty.
	// Example: /triples?g=http://knox_database&s=subjekt1&s=subjekt2&o=object1&p=predicate1
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: g
	//   in: query
	//   description: Target graph of the query. Currently (http://knox_ontology, http://knox_database) are valid graphs, but this could change in the future. Error responses will always contain the valid graphs, so you can just send an empty request to get the valid graphs.
	//   required: true
	//   type: string
	//   example: http://knox_database
	// - name: s
	//   in: query
	//   description: Subjects
	//   required: false
	//   type: array
	//   items:
	//     type: string
	// - name: o
	//   in: query
	//   description: Objects
	//   required: false
	//   type: array
	//   items:
	//     type: string
	// - name: p
	//   in: query
	//   description: Predicates
	//   required: false
	//   type: array
	//   items:
	//     type: string
	// responses:
	//   '200':
	//     description: filtered triples response
	//     schema:
	//       "$ref": "#/definitions/Result"
	return func(c *gin.Context) {
		targetGraph := graph.TargetGraph(c.Query(string(graph.Graph)))
		if err := targetGraph.Validate(); err != nil {
			msg := fmt.Sprintf("error validating target graph: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		edges := c.QueryArray(string(graph.Predicate))
		subjects := c.QueryArray(string(graph.Subject))
		objects := c.QueryArray(string(graph.Object))
		_depth := c.DefaultQuery("depth", "0")

		depth, err := strconv.Atoi(_depth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(
				"depth must be an integer, got %s", _depth)})
			return
		}

		if err := validateQuery(edges, subjects, objects, depth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := s.GETBuilder(edges, subjects, objects, depth, targetGraph)
		triples, err := s.ExecuteGET(query)
		if err != nil {
			msg := fmt.Sprintf("repository: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, Result{
			Triples: triples,
			Query:   query,
		})
	}
}
func postHandler(s graph.Service) func(c *gin.Context) {
	// swagger:operation POST /triples UpsertTriples
	//
	// This endpoint upserts triples.
	//
	// If a new predicate is sent with an existing subject, will the existing subject be updated with the new predicate.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: g
	//   in: query
	//   description: Target graph of the query. Only valid graphs will be accepted. If the graph does not exist, the error response will contain the valid graphs.
	//   required: true
	//   example: http://knox_database
	//   type: string
	// - name: triples
	//   in: body
	//   description: Triples to upsert. Each triple must contain a subject, predicate and object, in that order.
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/PostBody"
	//   example:
	//     triples:
	//     - subject: http://knox_database/subject1
	//       predicate: http://knox_database/predicate1
	//       object: http://knox_database/object1
	//     - subject: http://knox_database/subject2
	//       predicate: http://knox_database/predicate2
	//       object: http://knox_database/object2
	// responses:
	//   '200':
	//     description: response with produced query and null value for triples
	//     schema:
	//       "$ref": "#/definitions/Result"
	return func(c *gin.Context) {
		var triples graph.PostBody
		graphParameter := graph.TargetGraph(c.Query(string(graph.Graph)))

		decoder := json.NewDecoder(c.Request.Body)
		if err := decoder.Decode(&triples); err != nil {
			msg := fmt.Sprintf("error parsing json body: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		if err := triples.Validate(); err != nil {
			msg := fmt.Sprintf("error validating json body: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		query := s.POSTBuilder(triples, graphParameter)
		if err := s.ExeutePOST(query); err != nil {
			msg := fmt.Sprintf("error executing query: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, Result{
			Triples: nil,
			Query:   query,
		})
	}
}

func validateQuery(edges, subject, object []string, depth int) error {
	if depth < 0 {
		return fmt.Errorf("depth must be a positive integer, got %d", depth)
	}

	if len(edges) == 0 && len(subject) == 0 && len(object) == 0 && depth != 0 {
		return fmt.Errorf("depth must be 0 if no edge or node is specified, got %d", depth)
	}

	return nil
}
