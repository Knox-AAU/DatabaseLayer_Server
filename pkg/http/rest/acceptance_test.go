package rest_test

import (
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
	"github.com/stretchr/testify/require"
)

func TestAcceptance(t *testing.T) {

	input := "/get?p=x&p=y&s=z&s=j&o=h&o=k"
	expectedQuery := "SELECT ?s ?p ?o WHERE { ?s ?p ?o . FILTER ((contains(str(?s), 'z') || contains(str(?s), 'j')) && (contains(str(?o), 'h') || contains(str(?o), 'k')) && (contains(str(?p), 'x') || contains(str(?p), 'y'))) . }"
	actualResponse, statusCode := doRequest(input, t)

	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedQuery, actualResponse.Query)
}

func doRequest(path string, t *testing.T) (rest.Result, int) {
	appRepository := config.Repository{}
	config.LoadEnv("../../../", &appRepository)
	virtuosoRepository := virtuoso.NewVirtuosoRepository(appRepository.VirtuosoServerURL)
	service := graph.NewService(virtuosoRepository)
	router := rest.NewRouter(service)

	const GET = "GET"

	req, err := http.NewRequest(GET, path, nil)
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

func printDirContents(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
