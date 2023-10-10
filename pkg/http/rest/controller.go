package rest

import (
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/gin-gonic/gin"
)

type Body struct {
	Query string `json:"query"`
}

func GETallHandler(c *gin.Context, r graph.Service) {
	// swagger:operation GET /get-all getAllTriples
	//
	// Returns all triples that exist on the database
	//
	// The endpoint is only intended to work temporarily, until more defined use cases are implemented.
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: all triples response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/Triple"
	triples, err := r.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if triples == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "resource does not exist"})
		return
	}

	c.JSON(http.StatusOK, triples)
}
