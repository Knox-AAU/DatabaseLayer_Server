package graph

// Service provides operations
type Service interface {
	FindAll() (*[]Triple, error)
	Find(string) (*[]Triple, error)
	Delete(string) error
	Update(*Triple) error
	Create(*Triple) error
}

// Repository sends queries and turns response from virtuoso server into a Node with eventual children
type Repository interface {
	FindAll() (*[]Triple, error)
	Find(string) (*[]Triple, error)
	Delete(string) error
	Update(*Triple) error
	Create(*Triple) error
}

// service implements Service interface
type service struct {
	r Repository
}

// NewService creates service instance with given dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) FindAll() (*[]Triple, error) {
	return s.r.FindAll()
}

func (s *service) Find(query string) (*[]Triple, error) {
	return s.r.Find(query)
}

func (s *service) Delete(query string) error {
	return s.r.Delete(query)
}

func (s *service) Update(n *Triple) error {
	return s.r.Update(n)
}

func (s *service) Create(n *Triple) error {
	return s.r.Create(n)
}
