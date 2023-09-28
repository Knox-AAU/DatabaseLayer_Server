package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type virtuosoRepository struct {
	VirtuosoServerURL string
}

func NewVirtuosoRepository(url string) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
	}
}

func (r virtuosoRepository) FindAll() (*graph.VirtuosoObject, error) {
	query := `/sparql?query=SELECT+?subject+?predicate+?object+WHERE+{+?subject+?predicate+?object+}&format=json`
	response, err := http.Get(r.VirtuosoServerURL + query)
	if err != nil {
		return nil, fmt.Errorf("could not get response from virtuoso server: %v", err)
	}

	result := graph.VirtuosoObject{}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("could not decode response from virtuoso server: %v", err)
	}

	return &result, nil
}

func (r virtuosoRepository) Find(serviceQuery string) (*graph.Node, error) {
	return &graph.Node{}, nil
}

func (r virtuosoRepository) Delete(serviceQuery string) error {
	return nil
}

func (r virtuosoRepository) Update(node *graph.Node) error {
	return nil
}

func (r virtuosoRepository) Store(node *graph.Node) error {
	return nil
}
