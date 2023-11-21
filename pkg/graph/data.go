package graph

type (
	DataType                  int
	Attribute                 string
	TargetGraph               string
	OntologyGraphURI          TargetGraph
	KnowledgeBaseGraphURI     TargetGraph
	TestKnowledgeBaseGraphURI TargetGraph
)

const (
	URI DataType = iota
	Literal
	Subject   Attribute = "s"
	Predicate Attribute = "p"
	Object    Attribute = "o"
)

// GetTriple requires the json tags to match with the queries that are used to retrieve it.
// swagger:model
type GetTriple struct {
	// S is the subject
	S BindingAttribute `json:"s"`
	// P is the predicate
	P BindingAttribute `json:"p"`
	// O is the object
	O BindingAttribute `json:"o"`
}

//swagger:model
type PostBody struct {
	// Triples is an array of triples, where each triple's first element is the subject, second is the predicate and third is the object.
	Triples [][3]string `json:"triples"`
}

// swagger:model
type BindingAttribute struct {
	Type  string
	Value string
}

type VirtuosoResponse struct {
	Results struct {
		Bindings []GetTriple
	}
}
