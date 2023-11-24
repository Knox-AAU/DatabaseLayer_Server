package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type virtuosoRepository struct {
	VirtuosoServerURL config.VirtuosoURL
	Username          string
	Password          string
}

// encode adds necessary parameters for virtuoso
func encode(query string) string {
	params := url.Values{}
	params.Add("query", query)
	params.Add("format", "json")
	return params.Encode()
}

func NewVirtuosoRepository(url config.VirtuosoURL, username, password string) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
		Username:          username,
		Password:          password,
	}
}

func (r virtuosoRepository) send(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(r.Username, r.Password)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized request")
	}

	return res, nil
}

func (r virtuosoRepository) ExecuteGET(query string) ([]graph.GetTriple, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", r.VirtuosoServerURL, encode(query)), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	res, err := r.send(req)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	virtuosoRes := graph.VirtuosoResponse{}
	if err := json.Unmarshal(buf.Bytes(), &virtuosoRes); err != nil {
		log.Println(buf.Bytes())
		return nil, fmt.Errorf("error unmarshalling response: %s", err.Error())
	}

	return virtuosoRes.Results.Bindings, nil
}

func (r virtuosoRepository) ExeutePOST(query string) error {
	req, err := http.NewRequest("POST", string(r.VirtuosoServerURL), bytes.NewBuffer([]byte(query)))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/sparql-update")
	if _, err := r.send(req); err != nil {
		return fmt.Errorf("executing query: %w", err)
	}

	return nil
}
