package graph

type Service interface {
	ExecuteGET(string) ([]GetTriple, error)
	ExeutePOST(string) error
	GETBuilder([]string, []string, []string, int, TargetGraph) string
	POSTBuilder(PostBody, TargetGraph) string
}

type Repository interface {
	ExecuteGET(string) ([]GetTriple, error)
	ExeutePOST(string) error
	// GETBuilder takes three arrays of strings, and a limit and returns a SPARQL query
	GETBuilder([]string, []string, []string, int, TargetGraph) string
	POSTBuilder(PostBody, TargetGraph) string
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

func (s *service) POSTBuilder(triples PostBody, targetGraph TargetGraph) string {
	return s.r.POSTBuilder(triples, targetGraph)
}

func (s *service) ExecuteGET(query string) ([]GetTriple, error) {
	return s.r.ExecuteGET(query)
}

func (s *service) ExeutePOST(query string) error {
	return s.r.ExeutePOST(query)
}
