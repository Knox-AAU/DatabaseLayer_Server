package graph

import (
	"log"
)

// Service provides operations
type Service interface {
	FindAll() (*VirtuosoObject, error)
	Find(string) (*Node, error)
	Delete(string) error
	Update(*Node) error
	Store(*Node) error
}

// Repository sends queries and turns response from virtuoso server into a Node with eventual children
type Repository interface {
	FindAll() (*VirtuosoObject, error)
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

func (s *service) FindAll() (*VirtuosoObject, error) {
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

// Read creates the node datastructure from a virtuoso response
func (v *VirtuosoObject) Read() (*[]*Node, error) {
	existingNodes := []*Node{}
	for _, binding := range v.Results.Bindings {
		var matchingNode *Node

		for _, node := range existingNodes {
			matchingNode = findDescendant(binding.Subject.Value, node)
			if matchingNode != nil {
				break
			}
		}

		if matchingNode == nil {
			existingNodes = append(existingNodes, newNode(binding.Subject))
			continue
		}

		// found node from subject value, add object to its children,
		// avoiding duplicates
		childExists := false
		for _, child := range *matchingNode.Children {
			if child.Value == binding.Object.Value {
				childExists = true
				break
			}
		}

		if childExists {
			continue
		}

		*matchingNode.Children = append(*matchingNode.Children, newNode(binding.Object))
	}

	return &existingNodes, nil
}

// newNode creates node from binding
func newNode(bindingAttribute BindingAttribute) *Node {
	dataType := URI

	switch bindingAttribute.Type {
	case "uri":
		dataType = URI
	case "literal":
		dataType = Literal
	case "bnode":
		dataType = BNode
	case "typed-literal":
		dataType = TypedLiteral
	default:
		log.Fatalf("unknown datatype: %v", bindingAttribute.Type)
	}

	node := &Node{
		Value:    bindingAttribute.Value,
		DataType: DataType(dataType),
		Children: &[]*Node{},
	}
	return node
}

// findDescendant goes through a node's childs recursively, returning nil or a node matching the passed value
func findDescendant(value string, node *Node) *Node {
	if node.Value == value {
		return node
	}

	if node.Children == nil {
		return nil
	}

	result := []*Node{}
	for _, child := range *node.Children {
		node := findDescendant(value, child)
		if node == nil {
			continue
		}

		result = append(result, node)
	}

	if len(result) > 1 {
		log.Fatalf("internal error: node should never have multiple childs with value %s\n", value)
		// return result[0]
	}

	if len(result) == 0 {
		return nil
	}

	return result[0]
}

// getNodeFromURI will recieve bindingattribute and toggle URI and literal, if URI will fetch nodes, if literal returns node without fetch
func getNodeFromURI(switcher BindingAttribute) *Node {
	return nil
}
