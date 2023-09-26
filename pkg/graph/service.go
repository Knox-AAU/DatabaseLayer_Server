package graph

// Service provides operations
type Service interface {
	FindAll() (*Node, error)
	Find(string) (*Node, error)
	Delete(string) error
	Update(*Node) error
	Store(*Node) error
}

// Repository sends queries and turns response from virtuoso server into a Node with eventual children
type Repository interface {
	FindAll() (*Node, error)
	Find(string) (*Node, error)
	Delete(string) error
	Update(*Node) error
	Store(*Node) error
}

// service implements Service interface
type service struct {
	r Repository
}

// NewService creates service instance with given dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindAll() (*Node, error) {
	return s.r.FindAll()
}

func (s *service) Find(query string) (*Node, error) {
	return s.r.Find(query)
}

func (s *service) Delete(query string) error {
	return s.r.Delete(query)
}

func (s *service) Update(n *Node) error {
	return s.r.Update(n)
}

func (s *service) Store(n *Node) error {
	return s.r.Store(n)
}

func NewGetAllQuery() string {
	return `SELECT ?s ?p ?o
	WHERE {
		?s ?p ?o
	}
	LIMIT 100`
}
