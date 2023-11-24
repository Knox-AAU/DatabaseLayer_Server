package rest_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/http/rest"
	virtuoso "github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage/virtuoso/http"
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
	GET             method = "GET"
	POST            method = "POST"
	testingGraphURI        = "http://testing"
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

func toPostBody(getTriple []graph.GetTriple) graph.PostBody {
	var body graph.PostBody
	body.Triples = make([][3]string, len(getTriple))
	for i, triple := range getTriple {
		body.Triples[i] = [3]string{
			triple.S.Value,
			triple.P.Value,
			triple.O.Value,
		}
	}

	return body
}

func TestAcceptancePOSTGET(t *testing.T) {
	router := setupApp()
	triple1, err := createRandomTriple()
	require.NoError(t, err, "could not create random triple")

	triple2, err := createRandomTriple()
	require.NoError(t, err, "could not create random triple")

	triples := []graph.GetTriple{*triple1, *triple2}
	body := toPostBody(triples)

	expectedPostQuery := fmt.Sprintf("INSERT DATA { GRAPH <%s> {\n",
		testingGraphURI)

	for _, triple := range triples {
		expectedPostQuery += fmt.Sprintf("<%s> <%s> <%s> .\n", triple.S.Value, triple.P.Value, triple.O.Value)
	}

	expectedPostQuery += "}\n}"

	parameters := fmt.Sprintf("?p=%s&p=%s&s=%s&s=%s&o=%s&o=%s",
		triple1.S.Value, triple2.S.Value, triple1.P.Value, triple2.P.Value, triple1.O.Value, triple2.O.Value)
	expectedGetQuery := fmt.Sprintf("SELECT ?s ?p ?o WHERE { GRAPH <%s> { ?s ?p ?o . FILTER ((contains(str(?s), '%s') || contains(str(?s), '%s')) && (contains(str(?o), '%s') || contains(str(?o), '%s')) && (contains(str(?p), '%s') || contains(str(?p), '%s'))) . }}",
		testingGraphURI, triple1.S.Value, triple2.S.Value, triple1.P.Value, triple2.P.Value, triple1.O.Value, triple2.O.Value)

	gotPOSTKnowledgeBaseResponse, gotPOSTKnowledgebaseStatus := doRequest(router, string(rest.KnowledgeBase), t, method("POST"), &body)

	require.Equal(t, http.StatusOK, gotPOSTKnowledgebaseStatus)
	require.Equal(t, expectedPostQuery, gotPOSTKnowledgeBaseResponse.Query)

	gotPOSTOntologyPostResponse, gotPOSTOntologyStatus := doRequest(router, string(rest.Ontology), t, method("POST"), &body)
	require.Equal(t, http.StatusOK, gotPOSTOntologyStatus)
	require.Equal(t, expectedPostQuery, gotPOSTOntologyPostResponse.Query)

	gotGETKnowledgebaseResponse, statusCode := doRequest(router, string(rest.KnowledgeBase)+parameters, t, method("GET"), nil)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedGetQuery, gotGETKnowledgebaseResponse.Query)

	gotGETOntologyResponse, statusCode := doRequest(router, string(rest.Ontology)+parameters, t, method("GET"), nil)
	require.Equal(t, http.StatusOK, statusCode)
	require.Equal(t, expectedGetQuery, gotGETOntologyResponse.Query)
	// assert that triples have been inserted
	require.Equal(t, triples, gotGETOntologyResponse.Triples)
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
