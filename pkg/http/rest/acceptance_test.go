package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
	"github.com/gin-gonic/gin"
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
	GET             method = "GET"
	POST            method = "POST"
	testingGraphURI        = "http://testing"
)

func TestAcceptanceGET(t *testing.T) {
	router := setupApp()
	parameters := "?p=x&p=y&s=z&s=j&o=h&o=k"
	expectedQuery := fmt.Sprintf("SELECT ?s ?p ?o WHERE { GRAPH <%s> { ?s ?p ?o . FILTER ((contains(str(?s), 'z') || contains(str(?s), 'j')) && (contains(str(?o), 'h') || contains(str(?o), 'k')) && (contains(str(?p), 'x') || contains(str(?p), 'y'))) . }}",
		testingGraphURI)
	gotKnowledgebaseResponse, statusCode := doRequest(router, string(rest.KnowledgeBase)+parameters, t, method("GET"), nil)

	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedQuery, gotKnowledgebaseResponse.Query)

	gotOntologyResponse, statusCode := doRequest(router, string(rest.Ontology)+parameters, t, method("GET"), nil)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedQuery, gotOntologyResponse.Query)
}

func TestAcceptancePOST(t *testing.T) {
	var body graph.PostBody
	router := setupApp()
	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	err = json.NewDecoder(file).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	expectedQuery := fmt.Sprintf("INSERT DATA { GRAPH <%s> {\n<http://test1> <http://test1> <http://test1> .\n<http://test2> <http://test2> <http://test2> .\n}}",
		testingGraphURI)
	gotKnowledgeBaseResponse, gotKnowledgebaseStatus := doRequest(router, string(rest.KnowledgeBase), t, method("POST"), &body)

	require.Equal(t, http.StatusOK, gotKnowledgebaseStatus)
	require.Equal(t, expectedQuery, gotKnowledgeBaseResponse.Query)

	gotOntologyResponse, gotOntologyStatus := doRequest(router, string(rest.Ontology), t, method("POST"), &body)
	require.Equal(t, http.StatusOK, gotOntologyStatus)
	require.Equal(t, expectedQuery, gotOntologyResponse.Query)
}

func setupApp() *gin.Engine {
	appRepository := config.Repository{}
	config.Load("../../../", &appRepository)
	testingOntologyURI := appRepository.TestGraphURI
	testingKnowledgebaseURI := appRepository.TestGraphURI

	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoURL, appRepository.VirtuosoUsername, appRepository.VirtuosoPassword)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service, graph.OntologyGraphURI(testingOntologyURI), graph.KnowledgeBaseGraphURI(testingKnowledgebaseURI))

	return router
}

func doRequest(router *gin.Engine, path string, t *testing.T, _method method, body *graph.PostBody) (rest.Result, int) {
	var req *http.Request
	var err error
	switch {
	case _method == method("GET"):
		{
			req, err = http.NewRequest(string(_method), path, nil)
		}
	case _method == method("POST") && body != nil:
		{
			jsonPayload, err := json.Marshal(*body)
			require.NoError(t, err)

			req, err = http.NewRequest(string(_method), path, bytes.NewBuffer(jsonPayload))
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
