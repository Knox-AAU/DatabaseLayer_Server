package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/storage"
)

type virtuosoRepository struct {
	VirtuosoServerURL string
}

func NewVirtuosoRepository(url string) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
	}
}

func (r virtuosoRepository) FindAll() (*[]graph.Triple, error) {
	res, err := http.Get(r.VirtuosoServerURL + "?" + fomatQuery(storage.GetAll))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	virtuosoRes := graph.VirtuosoResponse{}

	err = json.Unmarshal(buf.Bytes(), &virtuosoRes)
	if err != nil {
		return nil, err
	}

	return virtuosoResponseToTriples(virtuosoRes), nil
}

func (r virtuosoRepository) Find(serviceQuery string) (*[]graph.Triple, error) {
	return &[]graph.Triple{}, nil
}

func (r virtuosoRepository) Delete(serviceQuery string) error {
	return nil
}

func (r virtuosoRepository) Update(node *graph.Triple) error {
	return nil
}

func (r virtuosoRepository) Create(node *graph.Triple) error {
	return nil
}

// formatQuery adds necessary parameters for virtuoso
func fomatQuery(query string) string {
	params := url.Values{}
	params.Add("query", query)
	params.Add("format", "json")
	return params.Encode()
}

func virtuosoResponseToTriples(response graph.VirtuosoResponse) *[]graph.Triple {
	result := make([]graph.Triple, len(response.Results.Bindings))

	copy(result, response.Results.Bindings)

	return &result
}
