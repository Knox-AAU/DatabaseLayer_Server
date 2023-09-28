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

// VirtuosoToNode takes a virtuoso object and returns a Node
func VirtuosoToNode(v *VirtuosoObject) *[]Node {
	result := []Node{}

	//for _, b := range v.Results.Bindings {

	/*
		bindingExists = false
	*/

	//for _, node := range result {

	/*
		if (binding.subject.value = node.value)
			node.children.append(call helperfunction getNodeFromURI with b.object.value)
			bindingExists = true
			break
	*/

	//}

	/*
		if (bindingExists = false)
			result.append(call helperfunction getNodeFromURI with b.subject.value)
	*/
	//}

	return &result
}

// getNodeFromURI will recieve bindingattribute and toggle URI and literal, if URI will fetch nodes, if literal returns node without fetch
func getNodeFromURI(switcher BindingAttribute) *Node {
	return nil
}
