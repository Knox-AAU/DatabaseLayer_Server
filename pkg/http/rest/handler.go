package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

// swagger:model
type Result struct {
	Triples []graph.Triple `json:"triples"`
	Query   string         `json:"query"`
}

func getHandler(c *gin.Context, s graph.Service) {
	// swagger:operation GET /get get get
	//
	// This endpoint allows for querying with filters.
	//
	// Example query: {{url}}/get?p=x&p=y&s=x&s=y&o=x&o=y
	//
	// To query the whole graph, leave all parameters empty.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
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
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Result"

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

	query := graph.Builder(edges, subjects, objects, depth, s.GetURI())
	triples, err := s.Execute(query)
	if err != nil {
		msg := fmt.Sprintf("error executing query: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, Result{
		Triples: triples,
		Query:   query,
	})
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
