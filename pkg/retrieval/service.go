package retrieval

// Service provides operations
type Service interface {
	Query(string) (Node, error)
}

// Repository provides access to data store
type Repository interface {
	Query(string) (Node, error)
}

// service implements Service interface
type service struct {
	r Repository
}

// NewService creates service instance with given dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Query(query string) (Node, error) {
	return s.r.Query(query)
}
