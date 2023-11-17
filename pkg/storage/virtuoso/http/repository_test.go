package http_test

// TestFindAll checks that the method sends a sparql query to the
// virtuoso url and returns a formatted node from the response
// func TestFindAll(t *testing.T) {
// 	type arrange func(*[]graph.Triple)

// 	mockURL := "http://localhost:3033"

// 	repository := virtuoso_http.NewVirtuosoRepository(mockURL)

// 	cases := map[string]arrange{
// 		"returns empty virtuoso response": func(expected *[]graph.Triple) {
// 		},
// 	}
// 	for name, arrange := range cases {
// 		t.Run(name, func(t *testing.T) {
// 			expected := []graph.Triple{}

// 			arrange(&expected)

// 			//httpMockResponse := httpMockResponse(expected)
// 			httpmock.Activate()
// 			defer httpmock.DeactivateAndReset()
// 			httpmock.RegisterResponder("GET", mockURL,
// 				httpmock.NewStringResponder(http.StatusOK, "httpMockResponse"))

// 			actual, err := repository.FindAll()

// 			assert.NoError(t, err)
// 			assert.True(t, reflect.DeepEqual(expected, *actual))
// 		})
// 	}
// }

//func httpMockResponse(n graph.Node) string {
// base case
//if n.Child == nil {
//	return n.Value
//}

// recursive case
//return n.Value + *n.Label + httpMockResponse(*n.Child)
//}
