package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/sparql"
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
	// Example query: {{url}}/get?p=x&p=y&so=x&so=y
	// To query the whole graph, leave all parameters empty.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: so
	//   in: query
	//   description: Subject or Object
	//   required: false
	//   type: array
	//   items:
	//     type: string
	// - name: p
	//   in: query
	//   description: Predicate
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

	edges := c.QueryArray("p")
	nodes := c.QueryArray("so")
	_depth := c.DefaultQuery("depth", "0")

	depth, err := strconv.Atoi(_depth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(
			"depth must be an integer, got %s", _depth)})
		return
	}

	if err := validateQuery(edges, nodes, depth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := sparql.Builder(nodes, edges, depth)
	triples, err := s.Execute(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Result{
		Triples: triples,
		Query:   query,
	})
}

func validateQuery(edges, nodes []string, depth int) error {
	if depth < 0 {
		return fmt.Errorf("depth must be a positive integer, got %d", depth)
	}

	if len(edges) == 0 && len(nodes) == 0 && depth != 0 {
		return fmt.Errorf("depth must be 0 if no edge or node is specified, got %d", depth)
	}

	return nil
}
