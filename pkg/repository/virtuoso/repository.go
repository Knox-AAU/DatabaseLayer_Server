package virtuoso

import "github.com/Knox-AAU/DatabaseLayer_Server/pkg/retrieval"

type VirtuosoRepository struct {
}

func (r *VirtuosoRepository) Query(query string) (retrieval.Node, error) {
	return retrieval.Node{
		Value: "test",
	}, nil
}
