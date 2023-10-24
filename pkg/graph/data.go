package graph

type DataType int

const (
	URI = iota
	Literal
)

type Attribute string

const (
	Subject   Attribute = "s"
	Predicate Attribute = "p"
	Object    Attribute = "o"
)

// Triple requires the json tags to match with the queries that are used to retrieve it.
// swagger:model
type Triple struct {
	// S is the subject
	S BindingAttribute `json:"s"`
	// P is the predicate
	P BindingAttribute `json:"p"`
	// O is the object
	O BindingAttribute `json:"o"`
}

// swagger:model
type BindingAttribute struct {
	Type  string
	Value string
}

type VirtuosoResponse struct {
	Results struct {
		Bindings []Triple
	}
}
