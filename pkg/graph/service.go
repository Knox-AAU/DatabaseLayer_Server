package graph

type Service interface {
	Execute(string) ([]Triple, error)
	GetURI() string
}

type Repository interface {
	Execute(string) ([]Triple, error)
}

// service implements Service interface
type service struct {
	r   Repository
	uri string
}

// NewService creates service instance with given dependencies
func NewService(r Repository, uri string) Service {
	return &service{r: r, uri: uri}
}

func (s *service) Execute(query string) ([]Triple, error) {
	return s.r.Execute(query)
}

func (s *service) GetURI() string {
	return s.uri
}
