package graph

type Service interface {
	ExecuteGET(string) ([]Triple, error)
	ExeutePOST(string) error
	GETBuilder([]string, []string, []string, int, TargetGraph) string
	POSTBuilder([][3]string, TargetGraph) string
}

type Repository interface {
	// Execute executes a SPARQL query and returns the response from the repository
	ExecuteGET(string) ([]Triple, error)
	ExeutePOST(string) error
	// GETBuilder takes three arrays of strings, and a limit and returns a SPARQL query
	GETBuilder([]string, []string, []string, int, TargetGraph) string
	POSTBuilder([][3]string, TargetGraph) string
}

// service implements Service interface
type service struct {
	r Repository
}

// NewService creates service instance with given dependencies
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) GETBuilder(edges, subjects, objects []string, depth int, targetGraph TargetGraph) string {
	return s.r.GETBuilder(edges, subjects, objects, depth, targetGraph)
}

func (s *service) POSTBuilder(triples [][3]string, targetGraph TargetGraph) string {
	return s.r.POSTBuilder(triples, targetGraph)
}

func (s *service) ExecuteGET(query string) ([]Triple, error) {
	return s.r.ExecuteGET(query)
}

func (s *service) ExeutePOST(query string) error {
	return s.r.ExeutePOST(query)
}
