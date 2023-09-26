package http

import "github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"

type virtuosoRepository struct {
	VirtuosoServerURL string
}

func NewVirtuosoRepository(url string) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
	}
}

func (r virtuosoRepository) FindAll() (*graph.Node, error) {
	return &graph.Node{}, nil
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
