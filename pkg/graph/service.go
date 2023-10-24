package graph

type Service interface {
	Execute(string) ([]Triple, error)
}

type Repository interface {
	Execute(string) ([]Triple, error)
}

// service implements Service interface
type service struct {
	r Repository
}

// NewService creates service instance with given dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Execute(query string) ([]Triple, error) {
	return s.r.Execute(query)
}
