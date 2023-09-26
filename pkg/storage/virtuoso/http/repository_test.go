package http_test

import (
	"testing"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
	"github.com/stretchr/testify/require"
)

// TestFindAll checks that the method sends a sparql query to the
// virtuoso url and returns a formatted node from the response
func TestFindAll(t *testing.T) {
	type arrange func(*graph.Node)

	mockURL := "http://localhost:3033"

	repository := http.NewVirtuosoRepository(mockURL)

	cases := map[string]arrange{
		"returns virtuoso response": func(expected *graph.Node) {
			label := "label"
			*expected = graph.Node{
				Value: "parent",
				Child: &graph.Node{
					Value: "child",
				},
				Label: &label,
			}
		},
	}
	for name, arrange := range cases {
		t.Run(name, func(t *testing.T) {
			expected := graph.Node{}

			arrange(&expected)

			// httpMockResponse := httpMockResponse(expected)

			actual, err := repository.FindAll()

			require.NoError(t, err)
			require.Equal(t, expected, actual)
		})
	}
}

func httpMockResponse(n graph.Node) string {
	// base case
	if n.Child == nil {
		return n.Value
	}

	// recursive case
	return n.Value + *n.Label + httpMockResponse(*n.Child)
}
