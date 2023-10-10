package http_test

import (
	"net/http"
	"testing"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	virtuoso_http "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// TestFindAll checks that the method sends a sparql query to the
// virtuoso url and returns a formatted node from the response
func TestFindAll(t *testing.T) {
	type arrange func(*graph.Node)

	mockURL := "http://localhost:3033"

	repository := virtuoso_http.NewVirtuosoRepository(mockURL)

	cases := map[string]arrange{
		"returns empty virtuoso response": func(expected *graph.Node) {
		},
	}
	for name, arrange := range cases {
		t.Run(name, func(t *testing.T) {
			expected := graph.Node{}

			arrange(&expected)

			//httpMockResponse := httpMockResponse(expected)
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", mockURL,
				httpmock.NewStringResponder(http.StatusOK, "httpMockResponse"))

			actual, err := repository.FindAll()

			assert.NoError(t, err)
			assert.Equal(t, &expected, actual)
		})
	}
}

//func httpMockResponse(n graph.Node) string {
// base case
//if n.Child == nil {
//	return n.Value
//}

// recursive case
//return n.Value + *n.Label + httpMockResponse(*n.Child)
//}
