package graph

type (
	Attribute                 string
	TargetGraph               string
	OntologyGraphURI          TargetGraph
	KnowledgeBaseGraphURI     TargetGraph
	TestKnowledgeBaseGraphURI TargetGraph
)

const (
	// Subject matches with the json tag of the subject in the get query
	Subject Attribute = "s"
	// Predicate matches with the json tag of the predicate in the get query
	Predicate Attribute = "p"
	// Object matches with the json tag of the object in the get query
	Object Attribute = "o"
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
	// Triples is an array of triples.
	// Each triple's first element is the subject, second is the predicate and third is the object.
	// Only accepts exactly 3 elements per triple.
	// required: true
	Triples [][]string `json:"triples"`
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
