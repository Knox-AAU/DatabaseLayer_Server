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

type method string

const (
	GET  method = "GET"
	POST method = "POST"
)

func TestAcceptanceGET(t *testing.T) {
	router := setupApp()
	input := rest.GET + "?p=x&p=y&s=z&s=j&o=h&o=k"
	expectedQuery := "SELECT ?s ?p ?o WHERE { GRAPH <http://testing/> { ?s ?p ?o . FILTER ((contains(str(?s), 'z') || contains(str(?s), 'j')) && (contains(str(?o), 'h') || contains(str(?o), 'k')) && (contains(str(?p), 'x') || contains(str(?p), 'y'))) . }}"
	actualResponse, statusCode := doRequest(router, input, t, method("GET"), nil)

	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedQuery, actualResponse.Query)
}

func TestAcceptancePOST(t *testing.T) {
	var body []graph.Triple
	router := setupApp()

	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	err = json.NewDecoder(file).Decode(&body)
	if err != nil {
		log.Fatal(err)
	}

	expectedQuery := "INSERT DATA { GRAPH <http://testing/> {<http://testing/Barack_Obama> <http://dbpedia.org/ontology/spouse> <http://testing/Michelle_Obama>.}}"
	actualResponse, statusCode := doRequest(router, rest.POST, t, method("POST"), body)

	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedQuery, actualResponse.Query)
}

func setupApp() *gin.Engine {
	appRepository := config.Repository{}
	config.Load("../../../", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoURL, appRepository.TestGraphURI)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service)

	return router
}

func doRequest(router *gin.Engine, path string, t *testing.T, _method method, body []graph.Triple) (rest.Result, int) {
	var req *http.Request
	var err error
	switch {
	case _method != method("GET") && body != nil:
		{
			jsonPayload, err := json.Marshal(body)
			require.NoError(t, err)

			req, err = http.NewRequest(string(_method), path, bytes.NewBuffer(jsonPayload))
		}
	case _method != method("POST") && body != nil:
		{
			jsonPayload, err := json.Marshal(body)
			require.NoError(t, err)

			req, err = http.NewRequest(string(_method), path, bytes.NewBuffer(jsonPayload))
		}
	default:
		{
			req, err = http.NewRequest(string(_method), path, nil)
		}
	}

	require.NoError(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	response := rr.Result()

	defer response.Body.Close()
	actualResponse := rest.Result{}

	err = json.NewDecoder(response.Body).Decode(&actualResponse)
	require.NoError(t, err)

	return actualResponse, response.StatusCode
}
