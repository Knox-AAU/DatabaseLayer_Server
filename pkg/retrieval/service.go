package retrieval

// Service provides nutrient insertion operations
type Service interface {
	Query(string) (Node, error)
}

// Repository provides access to nutrient repository
type Repository interface {
	Query(string) (Node, error)
}

type service struct {
	r Repository
}

// NewService creates an insertion service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Query(query string) (Node, error) {
	return s.r.Query(query)
}
