package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	go_http "net/http"
	"net/http/httptest"
	"testing"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type (
	method           string
	internalResponse struct {
		rest.Result
		ErrMessage string `json:"error"`
	}
)

const (
	GET  method = "GET"
	POST method = "POST"
)

func createRandomTriple() (*graph.GetTriple, error) {
	const _type = "uri"
	s, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	o, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	p, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	triple := graph.GetTriple{
		S: graph.BindingAttribute{
			Type:  _type,
			Value: s.String(),
		},
		O: graph.BindingAttribute{
			Type:  _type,
			Value: o.String(),
		},
		P: graph.BindingAttribute{
			Type:  _type,
			Value: p.String(),
		},
	}
	return &triple, nil
}

func toPostBody(getTriple []graph.GetTriple) map[string][][]string {
	body := make(map[string][][]string)
	body["triples"] = make([][]string, len(getTriple))
	for i, triple := range getTriple {
		body["triples"][i] = []string{
			triple.S.Value,
			triple.P.Value,
			triple.O.Value,
		}
	}

	return body
}

func TestAcceptance(t *testing.T) {
	router, testGraph := setupApp()
	triple1, err := createRandomTriple()
	require.NoError(t, err, "could not create random triple")

	triple2, err := createRandomTriple()
	require.NoError(t, err, "could not create random triple")

	triples := []graph.GetTriple{*triple1, *triple2}
	body := toPostBody(triples)

	expectedPostQuery := fmt.Sprintf("INSERT DATA { GRAPH <%s> {\n",
		testGraph)

	for _, triple := range triples {
		expectedPostQuery += fmt.Sprintf("<%s> <%s> <%s> .\n", triple.S.Value, triple.P.Value, triple.O.Value)
	}

	expectedPostQuery += "}}"
	expectedGetQuery := fmt.Sprintf("SELECT ?s ?p ?o WHERE { GRAPH <%s> { ?s ?p ?o . FILTER ((contains(str(?s), '%s') || contains(str(?s), '%s')) && (contains(str(?o), '%s') || contains(str(?o), '%s')) && (contains(str(?p), '%s') || contains(str(?p), '%s'))) . }}",
		testGraph, triple1.S.Value, triple2.S.Value, triple1.O.Value, triple2.O.Value, triple1.P.Value, triple2.P.Value)

	route := fmt.Sprintf("%s?%s=%s",
		string(rest.Triples), graph.Graph, testGraph)

	gotPOSTResponse, gotPOSTStatus := doRequest(router, route, t, method("POST"), body)

	require.Equal(t, go_http.StatusOK, gotPOSTStatus)
	require.Equal(t, expectedPostQuery, gotPOSTResponse.Query)

	gotGETResponse, statusCode := doRequest(router, fmt.Sprintf("%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s",
		route, graph.Predicate, triple1.P.Value, graph.Predicate, triple2.P.Value,
		graph.Subject, triple1.S.Value, graph.Subject, triple2.S.Value,
		graph.Object, triple1.O.Value, graph.Object, triple2.O.Value), t, method("GET"), nil)

	require.Equal(t, go_http.StatusOK, statusCode)
	require.Equal(t, expectedGetQuery, gotGETResponse.Query)
	// assert that triples have been inserted
	sliceEquals(t, triples, gotGETResponse.Triples, gotGETResponse.Query, body)
}

func sliceEquals(t *testing.T, want, got []graph.GetTriple, msg ...interface{}) {
	for _, wantTriple := range want {
		found := false
		for _, gotTriple := range got {
			if wantTriple == gotTriple {
				found = true
			}
		}
		if !found {
			require.Equal(t, want, got, msg...)
		}
	}
}

func setupApp() (*gin.Engine, config.GraphURI) {
	appRepository := config.Repository{}
	config.Load("../../../", &appRepository)
	testingOntologyURI := appRepository.TestGraphURI
	testingKnowledgebaseURI := appRepository.TestGraphURI

	virtuosoRepository := http.NewVirtuosoRepository(appRepository.VirtuosoURL, appRepository.VirtuosoUsername, appRepository.VirtuosoPassword)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service, graph.OntologyGraphURI(testingOntologyURI), graph.KnowledgeBaseGraphURI(testingKnowledgebaseURI))

	return router, appRepository.TestGraphURI
}

func doRequest(router *gin.Engine, path string, t *testing.T, _method method, body map[string][][]string) (rest.Result, int) {
	var req *go_http.Request
	var err error
	switch {
	case _method == method("GET"):
		{
			req, err = go_http.NewRequest(string(_method), path, nil)
		}
	case _method == method("POST") && body != nil:
		{
			jsonPayload, err := json.Marshal(body)
			require.NoError(t, err)

			req, err = go_http.NewRequest(string(_method), path, bytes.NewBuffer(jsonPayload))
			if err != nil {
				t.Fatalf("Error in POST request, error: %s", err)
			}
		}
	default:
		{
			t.Fatalf("Invalid method: %s", _method)
		}
	}

	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	response := rr.Result()

	defer response.Body.Close()
	actualResponse := internalResponse{}

	err = json.NewDecoder(response.Body).Decode(&actualResponse)
	require.NoError(t, err)
	require.Empty(t, actualResponse.ErrMessage)

	return rest.Result{
		Query:   actualResponse.Query,
		Triples: actualResponse.Triples,
	}, response.StatusCode
}
